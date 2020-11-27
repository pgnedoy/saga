package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/9count/go-services/core/log"
)

const (
	defaultGracefulTimeout = 15 * time.Second
	defaultIdleTimeout     = 60 * time.Second
	defaultReadTimeout     = 15 * time.Second
	defaultWriteTimeout    = 15 * time.Second
)

type ServerOptions struct {
	GracefulTimeout time.Duration
	IdleTimeout     time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration

	Killer chan os.Signal
}

type ServerOption func(*ServerOptions)

func ServerGracefulTimeout(t time.Duration) ServerOption {
	return func(args *ServerOptions) {
		args.GracefulTimeout = t
	}
}

func ServerIdleTimeout(t time.Duration) ServerOption {
	return func(args *ServerOptions) {
		args.IdleTimeout = t
	}
}

func ServerReadTimeout(t time.Duration) ServerOption {
	return func(args *ServerOptions) {
		args.ReadTimeout = t
	}
}

func ServerWriteTimeout(t time.Duration) ServerOption {
	return func(args *ServerOptions) {
		args.WriteTimeout = t
	}
}

func ServerKiller(c chan os.Signal) ServerOption {
	return func(args *ServerOptions) {
		args.Killer = c
	}
}

type Server struct {
	handler http.Handler
	options *ServerOptions
}

func NewServer(h http.Handler, options ...ServerOption) (*Server, error) {
	args := &ServerOptions{
		GracefulTimeout: defaultGracefulTimeout,
		IdleTimeout:     defaultIdleTimeout,
		ReadTimeout:     defaultReadTimeout,
		WriteTimeout:    defaultWriteTimeout,
	}

	for _, option := range options {
		option(args)
	}

	return &Server{
		handler: h,
		options: args,
	}, nil
}

func (s *Server) Run(ctx context.Context, port int) error {
	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", port),

		// Good practice to set timeouts to avoid Slowloris attacks.
		IdleTimeout:  s.options.IdleTimeout,
		ReadTimeout:  s.options.ReadTimeout,
		WriteTimeout: s.options.WriteTimeout,

		Handler: s.handler,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Info(ctx, "server error", log.WithError(err))
		}
	}()

	killer := s.options.Killer

	if killer == nil {
		killer = make(chan os.Signal, 1)
	}

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(killer, os.Interrupt)

	// Block until we receive our signal.
	<-killer

	serverCtx, cancel := context.WithTimeout(ctx, s.options.GracefulTimeout)
	defer cancel()

	fmt.Println(s.options.GracefulTimeout)

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	return srv.Shutdown(serverCtx)
}
