package main

import (
	"encoding/json"
	"reflect"
	"strconv"

	build "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

// Database ...
type Database struct {
	Address string // add json later
	Signer  string
	Account horizon.Account
}

// Key is the key being referenced. 56 character max.
type Key string

// Descriptor provides additional functionality to a key.
// Descriptor has two forms.
// 1) 'type'
// 2) 'interger count'
// MUST BE 7 CHARACTERS
type Descriptor string

const (
	chunkSize int = 64
	opSize    int = 20
	maxCount  int = 9999999
)

func chunkMap(key string, inputValue interface{}) (map[string][]byte, error) {
	valueType := reflect.TypeOf(inputValue)
	keyMap := make(map[string][]byte)
	keyMap[key+".type"] = []byte(valueType.String())

	rawBytes, err := json.Marshal(inputValue)
	if err != nil {
		return nil, err
	}
	numBytes := len(rawBytes)
	for i := 0; i < numBytes; i += chunkSize {
		end := i + chunkSize
		if end > numBytes {
			end = numBytes
		}
		keyMap[key+"."+strconv.Itoa(i/chunkSize)] = rawBytes[i:end]
	}
	return keyMap, nil
}

func (db *Database) packTransactions(kvs map[string][]byte) ([]string, error) {
	txs := []string{}
	keys := make([]string, 0, len(kvs))
	for k := range kvs {
		keys = append(keys, k)
	}
	numKeys := len(keys)
	for i := 0; i < numKeys; i += opSize {
		end := i + chunkSize
		if end > numKeys {
			end = numKeys
		}
		tx, err := db.packSingleTransaction(keys[i:end], kvs)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
	}
	return txs, nil
}

func (db *Database) packSingleTransaction(keys []string, kvs map[string][]byte) (string, error) {
	tx := &build.TransactionBuilder{}
	err := tx.Mutate(build.SourceAccount{AddressOrSeed: db.Address})
	if err != nil {
		return "", err
	}
	err = tx.Mutate(build.TestNetwork)
	if err != nil {
		return "", err
	}
	err = tx.Mutate(build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient})
	if err != nil {
		return "", err
	}

	for _, k := range keys {
		err := tx.Mutate(build.SetData(k, kvs[k]))
		if err != nil {
			return "", err
		}
	}

	err = tx.Mutate(build.Defaults{})
	if err != nil {
		return "", err
	}

	txe, err := tx.Sign(db.Signer)
	if err != nil {
		return "", err
	}
	txeB64, err := txe.Base64()
	if err != nil {
		return "", err
	}
	return txeB64, nil
}
