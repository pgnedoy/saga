build:
	GOARCH=amd64 GOOS=linux go build -ldflags="-d -s -w" -o bin/main main.go
	chmod +x bin/main
	zip -j bin/main.zip bin/main

clean:
	rm -r bin



