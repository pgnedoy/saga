package http

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	t.Run("creates a server with no options", func(t *testing.T) {
		s, err := NewServer(mux.NewRouter())

		assert.Nil(t, err)
		assert.NotNil(t, s)
	})

	t.Run("creates a server with GracefulTimeout", func(t *testing.T) {
		ts := 8888 * time.Second

		s, err := NewServer(mux.NewRouter(), ServerGracefulTimeout(ts))

		assert.Nil(t, err)
		assert.NotNil(t, s)
		assert.NotNil(t, ts, s.options.GracefulTimeout)
	})

	t.Run("creates a server with IdleTimeout", func(t *testing.T) {
		ts := 8888 * time.Second

		s, err := NewServer(mux.NewRouter(), ServerIdleTimeout(ts))

		assert.Nil(t, err)
		assert.NotNil(t, s)
		assert.NotNil(t, ts, s.options.IdleTimeout)
	})

	t.Run("creates a server with ReadTimeout", func(t *testing.T) {
		ts := 8888 * time.Second

		s, err := NewServer(mux.NewRouter(), ServerReadTimeout(ts))

		assert.Nil(t, err)
		assert.NotNil(t, s)
		assert.NotNil(t, ts, s.options.ReadTimeout)
	})

	t.Run("creates a server with WriteTimeout", func(t *testing.T) {
		ts := 8888 * time.Second

		s, err := NewServer(mux.NewRouter(), ServerWriteTimeout(ts))

		assert.Nil(t, err)
		assert.NotNil(t, s)
		assert.NotNil(t, ts, s.options.WriteTimeout)
	})
}

// TODO: figure out why this is failing with unexpected deadline exceeded error
func TestServer_Run(t *testing.T) {
	t.Run("runs and stops gracefully", func(t *testing.T) {
		ts := 500 * time.Millisecond

		kill := make(chan os.Signal, 1)

		s, err := NewServer(mux.NewRouter(), ServerGracefulTimeout(ts), ServerKiller(kill))

		if err != nil {
			t.Error(err)
		}

		go func() {
			time.Sleep(1 * time.Millisecond)
			kill <- os.Interrupt
		}()

		err = s.Run(context.Background(), 23456)

		assert.Nil(t, err)
	})
}
