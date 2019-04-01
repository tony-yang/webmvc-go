all: build-container test install

build-container:
	docker build -t webmvc-go .

install: build-container
	docker run --rm webmvc-go bash -c "cd dev && go install webmvc"

test: build-container
	docker run --rm webmvc-go bash -c "cd dev/src/webmvc && go test ./... -v -cover"

run-server: build-container
	docker run -d --rm -p 9999:80 webmvc-go bash -c "cd dev/src/webmvc/main && go run main.go"
