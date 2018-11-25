// Copyright (c) FreitX Network
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package blockchain

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/freitx-project/freitx-network-blockchain/config"
	"github.com/freitx-project/freitx-network-blockchain/pkg/hash"
)

func TestGenesis(t *testing.T) {
	t.Logf("The TotalSupply is %d", Gen.TotalSupply)

	cfg := config.Default
	genesisBlk := NewGenesisBlock(&cfg)

	t.Log("The Genesis Block has the following header:")
	t.Logf("Version: %d", genesisBlk.Header.version)
	t.Logf("ChainID: %d", genesisBlk.Header.chainID)
	t.Logf("Height: %d", genesisBlk.Header.height)
	t.Logf("Timestamp: %d", genesisBlk.Header.timestamp)
	t.Logf("PrevBlockHash: %x", genesisBlk.Header.prevBlockHash)

	assert := assert.New(t)

	expectedParentHash := hash.Hash32B{}

	assert.Equal(uint32(1), genesisBlk.Header.version)
	assert.Equal(cfg.Chain.ID, genesisBlk.Header.chainID)
	assert.Equal(uint64(0), genesisBlk.Header.height)
	assert.Equal(uint64(1524676419), genesisBlk.Header.timestamp)
	assert.Equal(expectedParentHash, genesisBlk.Header.prevBlockHash)
}
