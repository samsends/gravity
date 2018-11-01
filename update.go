package main

import (
	"context"
	"encoding/json"
	"fmt"

	b "github.com/sbsends/go/clients/horizon"
	"github.com/stellar/go/protocols/horizon/operations"
)

func (db *Database) c() {

	client := b.DefaultTestNetClient
	cursor := b.Cursor("now")
	ctx := context.Background()

	err := client.StreamOperations(ctx, "GBOHNXALOLI6FXK7ECKKOMNG5FIZ3F25RZDLWRNHOUNHCELACSBYYIM5", &cursor, func(o interface{}) {
		manData, err := returnManageData(o)
		if err == nil {
			mutex.Lock()
			db.push(manData.Name, manData.Value)
			mutex.Unlock()
		}
	})
	if err != nil {
		db.c()
	}
}

func returnManageData(o interface{}) (operations.ManageData, error) {
	assertion := o.(map[string]interface{})
	if assertion["type"] == "manage_data" {
		var properType operations.ManageData
		jsonB, err := json.Marshal(o)
		if err != nil {
			return operations.ManageData{}, err
		}
		json.Unmarshal(jsonB, &properType)
		return properType, nil
	}
	return operations.ManageData{}, fmt.Errorf("unsupported")
}
