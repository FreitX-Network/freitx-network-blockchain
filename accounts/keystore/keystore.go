// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package keystore

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/freitx-project/freitx-network-blockchain/iotxaddress"
	"github.com/freitx-project/freitx-network-blockchain/pkg/keypair"
)

var (
	// ErrKey indicates the error of key
	ErrKey = errors.New("key error")
	// ErrExist is the error that the key already exists in map
	ErrExist = errors.New("key already exists")
	// ErrNotExist is the error that the key does not exist in map
	ErrNotExist = errors.New("key does not exist")
)

// Key defines the struct to be stored in keystore object
type Key struct {
	PublicKey  string
	PrivateKey string
	RawAddress string
}

// KeyStore defines an interface that supports operations on keystore object
type KeyStore interface {
	Has(string) (bool, error)
	Get(string) (*iotxaddress.Address, error)
	Store(string, *iotxaddress.Address) error
	Remove(string) error
	All() ([]string, error)
}

// plainKeyStore is a filesystem keystore which implements KeyStore interface
type plainKeyStore struct {
	directory string
}

// MemKeyStore is an in-memory keystore which implements KeyStore interface
type memKeyStore struct {
	accounts map[string]*iotxaddress.Address
}

// NewPlainKeyStore returns a new instance of plain keystore
func NewPlainKeyStore(dir string) (KeyStore, error) {
	if _, err := os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			return nil, errors.Wrapf(err, "failed to get the status of directory %s", dir)
		}
		if err := os.Mkdir(dir, 0700); err != nil {
			return nil, errors.Wrapf(err, "failed to make directory %s", dir)
		}
	}
	return &plainKeyStore{directory: dir}, nil
}

// Has returns whether the raw address already exists in keystore filesystem
func (ks *plainKeyStore) Has(rawAddr string) (bool, error) {
	if err := validateAddress(rawAddr); err != nil {
		return false, err
	}
	filePath := filepath.Join(ks.directory, rawAddr)
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, errors.Wrapf(err, "failed to get the status of directory %s", filePath)
	}
	return true, nil
}

// Get returns iotxaddress from keystore filesystem given raw address
func (ks *plainKeyStore) Get(rawAddr string) (*iotxaddress.Address, error) {
	if err := validateAddress(rawAddr); err != nil {
		return nil, err
	}
	filePath := filepath.Join(ks.directory, rawAddr)
	fd, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.Wrapf(ErrNotExist, "raw address = %s", rawAddr)
		}
		return nil, errors.Wrapf(err, "failed to open file %s", filePath)
	}
	defer fd.Close()
	key := &Key{}
	if err := json.NewDecoder(fd).Decode(key); err != nil {
		return nil, errors.Wrap(err, "failed to decode json file to key")
	}
	return keyToAddr(key)
}

// Store stores iotxaddress in keystore filesystem
func (ks *plainKeyStore) Store(rawAddr string, address *iotxaddress.Address) error {
	if err := validateAddress(rawAddr); err != nil {
		return err
	}
	filePath := filepath.Join(ks.directory, rawAddr)

	_, err := os.Stat(filePath)
	if err == nil {
		return errors.Wrapf(ErrExist, "raw address = %s", rawAddr)
	}
	if !os.IsNotExist(err) {
		return errors.Wrapf(err, "failed to get the status of file %s", filePath)
	}

	f, err := os.Create(filePath)
	if err != nil {
		return errors.Wrapf(err, "failed to create file %s", filePath)
	}
	defer f.Close()

	publicKey := keypair.EncodePublicKey(address.PublicKey)
	privateKey := keypair.EncodePrivateKey(address.PrivateKey)
	key := &Key{PublicKey: publicKey, PrivateKey: privateKey, RawAddress: address.RawAddress}
	sKey, err := json.Marshal(key)
	if err != nil {
		return errors.Wrap(err, "failed to marshal key")
	}
	_, err = io.Copy(f, bytes.NewReader(sKey))
	return err
}

// Remove removes the iotxaddress from keystore filesystem given raw address
func (ks *plainKeyStore) Remove(rawAddr string) error {
	if err := validateAddress(rawAddr); err != nil {
		return err
	}
	filePath := filepath.Join(ks.directory, rawAddr)
	err := os.Remove(filePath)
	if os.IsNotExist(err) {
		return errors.Wrapf(ErrNotExist, "raw address = %s", rawAddr)
	}
	return err
}

// All returns a list of raw addresses currently stored in keystore filesystem
func (ks *plainKeyStore) All() ([]string, error) {
	fd, err := os.Open(ks.directory)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open directory %s", ks.directory)
	}
	names, err := fd.Readdirnames(0)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read directory names")
	}
	rawAddrs := make([]string, 0, len(names))
	for _, rawAddr := range names {
		_, err := iotxaddress.GetPubkeyHash(rawAddr)
		if err != nil {
			return nil, err
		}
		rawAddrs = append(rawAddrs, rawAddr)
	}
	return rawAddrs, nil
}

// NewMemKeyStore creates a new instance of MemKeyStore
func NewMemKeyStore() KeyStore {
	return &memKeyStore{
		accounts: make(map[string]*iotxaddress.Address),
	}
}

// Has returns whether the given raw address already exists in map
func (ks *memKeyStore) Has(rawAddr string) (bool, error) {
	if err := validateAddress(rawAddr); err != nil {
		return false, err
	}
	_, ok := ks.accounts[rawAddr]
	return ok, nil
}

// Get returns iotxaddress stored in map given raw address of the account
func (ks *memKeyStore) Get(rawAddr string) (*iotxaddress.Address, error) {
	if err := validateAddress(rawAddr); err != nil {
		return nil, err
	}
	addr, ok := ks.accounts[rawAddr]
	if !ok {
		return nil, errors.Wrapf(ErrNotExist, "raw address = %s", rawAddr)
	}
	return addr, nil
}

// Store stores iotxaddress in map
func (ks *memKeyStore) Store(rawAddr string, addr *iotxaddress.Address) error {
	if err := validateAddress(rawAddr); err != nil {
		return err
	}
	// check if the key already exists in map
	if _, ok := ks.accounts[rawAddr]; ok {
		return errors.Wrapf(ErrExist, "raw address = %s", rawAddr)
	}

	ks.accounts[rawAddr] = addr
	return nil
}

// Remove removes the entry corresponding to the given raw address from map if exists
func (ks *memKeyStore) Remove(rawAddr string) error {
	if err := validateAddress(rawAddr); err != nil {
		return err
	}
	_, ok := ks.accounts[rawAddr]
	if !ok {
		return errors.Wrapf(ErrNotExist, "raw address = %s", rawAddr)
	}
	delete(ks.accounts, rawAddr)
	return nil
}

// All returns returns a list of raw addresses currently stored in map
func (ks *memKeyStore) All() ([]string, error) {
	rawAddrs := make([]string, 0, len(ks.accounts))
	for rawAddr := range ks.accounts {
		_, err := iotxaddress.GetPubkeyHash(rawAddr)
		if err != nil {
			return nil, err
		}
		rawAddrs = append(rawAddrs, rawAddr)
	}
	return rawAddrs, nil
}

//======================================
// private functions
//======================================
func keyToAddr(key *Key) (*iotxaddress.Address, error) {
	if key == nil {
		return nil, errors.Wrapf(ErrKey, "key must not be nil")
	}
	publicKey, err := keypair.DecodePublicKey(key.PublicKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode public key")
	}
	privateKey, err := keypair.DecodePrivateKey(key.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode private key")
	}
	return &iotxaddress.Address{PublicKey: publicKey, PrivateKey: privateKey, RawAddress: key.RawAddress}, nil
}

func validateAddress(addr string) error {
	// check if the address is valid
	_, err := iotxaddress.GetPubkeyHash(addr)
	if err != nil {
		return errors.Wrapf(err, "address format is invalid %s", addr)
	}
	return nil
}