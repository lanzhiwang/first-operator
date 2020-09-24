.PHONY: clean
clean:
	rm -rf build/_output

# Build binary
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/_output/bin/imoocpod-operator ./cmd/manager

.PHONY: test
build-image:
	docker build -t quay.io/opstree/imoocpod-operator:v0.1 -f build/Dockerfile .
