# golang-zhmcclient

zhmcclient - A golang client library for the IBM Z HMC Web Services API

## Generate Fake APIs

```bash
cd ./pkg/zhmcclient
go get github.com/maxbrunsfeld/counterfeiter/v6
go generate ./...
```

## Build

```bash
make sample-build
make sample-build-mac
make sample-build-docker
```

## Unit Test

```bash
# 'FILE' corresponds to the filename w/o .go extentsion
# 'PKG' package to be tested
# If only PKG is provided all files (go modules) under the package will be tested
make unit-test PKG=pkg/zhmcclient FILE=lpar
```

## Sample Usage

```bash
make sample-build
export HMC_ENDPOINT="https://192.168.195.118:6794"
export HMC_USERNAME=${username}
export HMC_PASSWORD=${password}
./bin/sample
```

or

```bash
make sample-build-mac
export HMC_ENDPOINT="https://192.168.195.118:6794"
export HMC_USERNAME=${username}
export HMC_PASSWORD=${password}
./bin/sample-mac
```
