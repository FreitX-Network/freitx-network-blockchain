// Copyright (c) FreitX Network
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package keystore

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/freitx-project/freitx-network-blockchain/test/testaddress"
)

const (
	pubKey1  = "5e24c6c19eb50a6da14d0d2841ee8b7f8e31771f31413466526f7a726f70d8a619421f066d2033c413cbaeb710de4056061c1ed728274cfaaa69a91436ec3fe2135d0e40e7fbae03"
	priKey1  = "c0c08b8f4887c62a07d07388575e03d650cdcf1cf8050ba09881ea49ecb9746a4854ee01"
	rawAddr1 = "io1qyqsqqqq8uhx9jtdc2xp5wx7nxyq3xf4c3jmxknzj23d2m"
	pubKey2  = "f8261681ee6e3261eb4aa61123b0edc10bd95c9bb366c6b54348cfef3a055f2f3a3d800277cb15a2c13ac1a44ff1c05191c5729aa62955cb0303e80eeeb24885c8df033405fc5201"
	priKey2  = "6bee2200fa46913e8802a594580f26fa42f75d90ae599cab700bfd22bc6d4b52b34e5301"
	rawAddr2 = "io1qyqsqqqqa3nkp636trcg85x2jfq5rhflcut8ge042xzyfu"
)

func TestPlainKeyStore(t *testing.T) {
	require := require.New(t)

	ksDir := filepath.Join(os.TempDir(), "keystore")

	os.RemoveAll(ksDir)
	defer func() {
		require.NoError(os.RemoveAll(ksDir))
	}()

	ks, err := NewPlainKeyStore(ksDir)
	require.NoError(err)
	addr1 := testaddress.ConstructAddress(1, pubKey1, priKey1)
	//Test Store
	err = ks.Store("123", addr1)
	require.Error(err)

	err = ks.Store(rawAddr1, addr1)
	require.NoError(err)

	err = ks.Store(rawAddr1, addr1)
	require.Equal(ErrExist, errors.Cause(err))

	// Test Get
	_, err = ks.Get("123")
	require.Error(err)

	_, err = ks.Get("io1qyqsyqcy6m6hkqkj3f4w4eflm2gzydmvc0mumm7kgax4l3")
	require.Equal(ErrNotExist, errors.Cause(err))

	val, err := ks.Get(addr1.RawAddress)
	require.NoError(err)
	require.Equal(addr1, val)

	// Test Has
	_, err = ks.Has("123")
	require.Error(err)

	exist, err := ks.Has(rawAddr1)
	require.NoError(err)
	require.Equal(true, exist)

	// Test Remove
	err = ks.Remove("123")
	require.Error(err)

	err = ks.Remove("io1qyqsyqcy6m6hkqkj3f4w4eflm2gzydmvc0mumm7kgax4l3")
	require.Equal(ErrNotExist, errors.Cause(err))

	err = ks.Remove(rawAddr1)
	require.NoError(err)
	exist, err = ks.Has(rawAddr1)
	require.NoError(err)
	require.Equal(false, exist)

	// Test All
	addr2 := testaddress.ConstructAddress(1, pubKey2, priKey2)
	err = ks.Store(rawAddr1, addr1)
	require.NoError(err)
	err = ks.Store(rawAddr2, addr2)
	require.NoError(err)
	rawAddrs, err := ks.All()
	require.NoError(err)
	require.Equal(2, len(rawAddrs))
}

func TestMemKeyStore(t *testing.T) {
	require := require.New(t)

	ks := NewMemKeyStore()
	addr1 := testaddress.ConstructAddress(1, pubKey1, priKey1)
	// Test Store
	err := ks.Store("123", addr1)
	require.Error(err)

	err = ks.Store(rawAddr1, addr1)
	require.NoError(err)

	err = ks.Store(rawAddr1, addr1)
	require.Equal(ErrExist, errors.Cause(err))

	// Test Get
	_, err = ks.Get("123")
	require.Error(err)

	_, err = ks.Get("io1qyqsyqcy6m6hkqkj3f4w4eflm2gzydmvc0mumm7kgax4l3")
	require.Equal(ErrNotExist, errors.Cause(err))

	val, err := ks.Get(rawAddr1)
	require.NoError(err)
	require.Equal(addr1, val)

	// Test Has
	_, err = ks.Has("123")
	require.Error(err)

	exist, err := ks.Has(rawAddr1)
	require.NoError(err)
	require.Equal(true, exist)

	// Test Remove
	err = ks.Remove("123")
	require.Error(err)

	err = ks.Remove("io1qyqsyqcy6m6hkqkj3f4w4eflm2gzydmvc0mumm7kgax4l3")
	require.Equal(ErrNotExist, errors.Cause(err))

	err = ks.Remove(rawAddr1)
	require.NoError(err)
	exist, err = ks.Has(rawAddr1)
	require.NoError(err)
	require.Equal(false, exist)

	// Test All
	addr2 := testaddress.ConstructAddress(1, pubKey2, priKey2)
	err = ks.Store(rawAddr1, addr1)
	require.NoError(err)
	err = ks.Store(rawAddr2, addr2)
	require.NoError(err)
	rawAddrs, err := ks.All()
	require.NoError(err)
	require.Equal(2, len(rawAddrs))
}
