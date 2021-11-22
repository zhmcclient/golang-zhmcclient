# Makefile

BUILD_ROOT=${PWD}
PKG_DIR=$(PKG)
FILENAME=$(FILE)
GO_BLD_ARCH=$(shell go env GOHOSTARCH)
GOOS=${shell uname -s | awk '{print tolower($0)}'}


.PHONY: clean
clean:
	rm -rf bin

.PHONY: sample-build
sample-build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=$(GO_BLD_ARCH) go build -a -installsuffix cgo -o $(BUILD_ROOT)/bin/sample $(BUILD_ROOT)/sample.go


.PHONY: unit-test
unit-test:
	bash go-test.sh $(PKG_DIR) $(FILENAME)

.PHONY: sample-build-mac
sample-build-mac: clean
	CGO_ENABLED=0 GOOS=darwin GOARCH=$(GO_BLD_ARCH) go build -a -installsuffix cgo -o $(BUILD_ROOT)/bin/sample-mac $(BUILD_ROOT)/sample.go

.PHONY: sample-build-mac-docker
sample-build-docker:
	docker build $(BUILD_ROOT) -t zhmcclient-sample:dev

