```dep ensure```

```make fmt; make build```

```make test```

```make reboot``` reboots server from fresh database.

```make docker```

`make build; ./bin/addrgen`

`make; make run`, `./bin/iotc [commands] [flags]`

Ubuntu Library: export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:$GOPATH/src/github.com/freitx-project/freitx-network-blockchain/crypto/lib:$GOPATH/src/github.com/freitx-project/freitx-network-blockchain/crypto/lib/blslib


Test Contract

pragma solidity ^0.4.0;

contract SimpleStorage {
   uint storedData;

   function set(uint x) public {
       storedData = x;
   }

   function get() public constant returns (uint) {
       return storedData;
   }
}
