# golang-zhmcclient
zhmcclient - A golang client library for the IBM Z HMC Web Services API

## Generate Fake APIs
```
cd ./pkg/zhmcclient
go generate ./...
```

## Build
```
make sample-build
make sample-build-docker
```

## Unit Test
```
cd ./pkg/zhmcclient
go test
```

## Sample Usage
```
make sample-build
export HMC_ENDPOINT="https://192.168.195.118:9955"
export HMC_USERNAME=${username}
export HMC_PASSWORD=${password}
./bin/sample
```
