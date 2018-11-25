## Minimum requirements

| Components | Version | Description |
|----------|-------------|-------------|
|[Golang](https://golang.org) | >= 1.10.2 | The Go Programming Language |

### Setup Dev Environment
```
mkdir -p ~/go/src/github.com/freitx-project
cd ~/go/src/github.com/freitx-project
git clone git@github.com:freitx-project/freitx-network-blockchain.git
cd freitx-network-blockchain
```

```dep ensure --vendor-only```

```make fmt; make build```

LD_LIBRARY_PATH=$LD_LIBRARY_PATH:$GOPATH/src/github.com/freitx-project/freitx-network-blockchain/crypto/lib:$GOPATH/src/github.com/freitx-project/freitx-network-blockchain/crypto/lib/blslib

### Run Unit Tests
```make test```

### Reboot
```make reboot``` reboots server from fresh database.

### Run
```make run``` continue server.
