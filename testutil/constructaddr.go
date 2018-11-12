// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package testutil

import (
	"github.com/freitx-project/freitx-network-blockchain/iotxaddress"
	"github.com/freitx-project/freitx-network-blockchain/logger"
	"github.com/freitx-project/freitx-network-blockchain/pkg/keypair"
)

// ConstructAddress constructs an iotex address
func ConstructAddress(pubkey, prikey string) *iotxaddress.Address {
	pubk, err := keypair.DecodePublicKey(pubkey)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to construct the address")
	}
	prik, err := keypair.DecodePrivateKey(prikey)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to construct the address")
	}
	addr, err := iotxaddress.GetAddressByPubkey(iotxaddress.IsTestnet, iotxaddress.ChainID, pubk)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to construct the address")
	}
	addr.PrivateKey = prik
	return addr
}
