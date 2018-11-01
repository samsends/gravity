package main

import (
	"net/http"

	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
)

// CreateDatabase ...
func CreateDatabase() (*Database, error) {
	sourcePair, err := keypair.Random()
	if err != nil {
		return nil, err
	}
	_, err = http.Get("https://friendbot.stellar.org/?addr=" + sourcePair.Address())
	if err != nil {
		return nil, err
	}
	account, err := horizon.DefaultTestNetClient.LoadAccount(sourcePair.Address())
	if err != nil {
		return nil, err
	}
	return &Database{
		Address: sourcePair.Address(),
		Signer:  sourcePair.Seed(),
		Account: account,
	}, nil
}

// LoadDatabase ...
func LoadDatabase(address string, signer string) (*Database, error) {
	account, err := horizon.DefaultTestNetClient.LoadAccount(address)
	if err != nil {
		return nil, err
	}
	return &Database{
		Address: address,
		Signer:  signer,
		Account: account,
	}, nil
}
