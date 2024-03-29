// Copyright (c) FreitX Network
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package explorer

import (
	"github.com/coopernurse/barrister-go"

	"github.com/freitx-project/freitx-network-blockchain/explorer/idl/explorer"
)

// NewExplorerProxy accepts an URL to the endpoint of the Explorer server
// and returns a proxy that implements the Explorer interface
func NewExplorerProxy(url string) explorer.Explorer {
	trans := &barrister.HttpTransport{Url: url}

	client := barrister.NewRemoteClient(trans, true)

	// calc.NewExplorerProxy() is provided by the idl2go
	// generated calc.go file
	return explorer.NewExplorerProxy(client)
}
