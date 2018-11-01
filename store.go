package main

import (
	"fmt"
	"strconv"

	"github.com/stellar/go/clients/horizon"
)

// Put ...
func (db *Database) Put(key string, value interface{}) error {
	mappedData, err := chunkMap(key, value)
	if err != nil {
		return err
	}
	b64txs, err := db.packTransactions(mappedData)
	if err != nil {
		return err
	}
	for _, val := range b64txs {
		_, err = horizon.DefaultTestNetClient.SubmitTransaction(val)
		if err != nil {
			return err
		}
	}
	return nil
}

// Get ...
func (db *Database) Get(key string) (interface{}, error) {
	var value []byte

	for i := 0; i < maxCount; i++ {
		itt := key + "." + strconv.Itoa(i)
		byteData, err := db.pull(itt)
		if err != nil {
			break
		} else if byteData == nil {
			break
		}
		value = append(value, byteData...)
	}
	return value, nil
}

func (db *Database) pull(key string) ([]byte, error) {
	returnBytes, err := db.Account.GetData(key)
	if err != nil {
		return nil, err
	}
	return returnBytes, nil
}

func (db *Database) push(key string, value string) {
	fmt.Println("syncing: " + key)
	db.Account.Data[key] = value
}
