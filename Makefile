all: build-container test install

build-container:
	docker build -t webmvc-go .

install: build-container
	docker run --rm webmvc-go bash -c "cd dev && go install webmvc"

test: build-container
	@docker run --rm webmvc-go bash -c 'cd dev/src/webmvc && for p in . ./base ./controllers; do echo ""; echo "######## Test $$p ########"; go test $$p -v -covermode=count -coverprofile=webmvc_cover.out; go tool cover -func=webmvc_cover.out; echo "######## END Test $$p ########"; echo ""; done'

run-server: build-container
	docker run -d --rm -p 9999:80 webmvc-go bash -c "cd dev/src/webmvc/main && go run *.go"
