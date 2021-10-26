# Makefile

BUILD_ROOT=${PWD}
GO_BLD_ARCH=$(shell go env GOHOSTARCH)


.PHONY: clean
clean:
	rm -rf bin

.PHONY: sample-build
sample-build: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=$(GO_BLD_ARCH) go build -a -installsuffix cgo -o $(BUILD_ROOT)/bin/sample $(BUILD_ROOT)/sample.go


.PHONY: sample-build-mac
sample-build-mac: clean
	CGO_ENABLED=0 GOOS=darwin GOARCH=$(GO_BLD_ARCH) go build -a -installsuffix cgo -o $(BUILD_ROOT)/bin/sample $(BUILD_ROOT)/sample.go

.PHONY: sample-build-mac-docker
sample-build-docker:
	docker build $(BUILD_ROOT) -t zhmcclient-sample:dev

