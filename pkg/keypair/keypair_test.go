// Copyright (c) FreitX Network
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package keypair

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

const (
	publicKey  = "5e24c6c19eb50a6da14d0d2841ee8b7f8e31771f31413466526f7a726f70d8a619421f066d2033c413cbaeb710de4056061c1ed728274cfaaa69a91436ec3fe2135d0e40e7fbae03"
	privateKey = "c0c08b8f4887c62a07d07388575e03d650cdcf1cf8050ba09881ea49ecb9746a4854ee01"
)

func TestKeypair(t *testing.T) {
	require := require.New(t)

	_, err := DecodePublicKey("")
	require.Equal(ErrPublicKey, errors.Cause(err))
	_, err = DecodePrivateKey("")
	require.Equal(ErrPrivateKey, errors.Cause(err))

	pubKey, err := DecodePublicKey(publicKey)
	require.Nil(err)
	priKey, err := DecodePrivateKey(privateKey)
	require.Nil(err)

	require.Equal(publicKey, EncodePublicKey(pubKey))
	require.Equal(privateKey, EncodePrivateKey(priKey))

	_, err = StringToPubKeyBytes("")
	require.Equal(ErrPublicKey, errors.Cause(err))
	_, err = StringToPriKeyBytes("")
	require.Equal(ErrPrivateKey, errors.Cause(err))

	pubKeyBytes, err := StringToPubKeyBytes(publicKey)
	require.Nil(err)
	priKeyBytes, err := StringToPriKeyBytes(privateKey)
	require.Nil(err)

	pubKeyString, err := BytesToPubKeyString(pubKeyBytes)
	require.Nil(err)
	priKeyString, err := BytesToPriKeyString(priKeyBytes)
	require.Nil(err)
	require.Equal(publicKey, pubKeyString)
	require.Equal(privateKey, priKeyString)
}
