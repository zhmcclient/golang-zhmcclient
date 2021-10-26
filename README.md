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

## Sample Usage
```
import (
	"github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

func main() {
	endpoint := "https://192.168.195.118:9955"
	creds := &zhmcclient.Options{
		Username:   "name",
		Password:   "psw",
		VerifyCert: false,
		Trace:		false,
	}
    client, _ := zhmcclient.NewClient(endpoint, creds)
	if client != nil {
		hmcManager := zhmcclient.NewManagerFromClient(client)
		hmcManager.ListCPCs()
	}
}
```
