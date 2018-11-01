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

	typeKey := key + ".type"
	dataType, err := db.pull(typeKey)
	if err != nil {
		return nil, err
	}

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

	fmt.Println(string(dataType))
	return value, nil
}

func (db *Database) pull(key string) ([]byte, error) {
	returnBytes, err := db.Account.GetData(key)
	if err != nil {
		return nil, err
	}
	return returnBytes, nil
}
