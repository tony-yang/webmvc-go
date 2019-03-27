all: build-container install test

build-container:
	docker build -t webmvc-go .

install:
	docker run --rm webmvc-go bash -c "cd dev && go install webmvc"

test: build-container
	docker run --rm webmvc-go bash -c "cd dev/src/webmvc && go test ./... -v"

run-server:
	docker run -d --rm -p 9999:80 webmvc-go bash -c "cd dev/src && go run main.go"
