build:
	go build -o bin/app

run: build
	bin/app

lint:
	gofmt -s -w .