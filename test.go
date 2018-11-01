package main

import (
	"context"
	"encoding/json"
	"fmt"

	b "github.com/sbsends/go/clients/horizon"
	"github.com/stellar/go/protocols/horizon/operations"
)

func c() {
	client := b.DefaultTestNetClient
	cursor := b.Cursor("now")
	ctx := context.Background()

	err := client.StreamOperations(ctx, "GBOHNXALOLI6FXK7ECKKOMNG5FIZ3F25RZDLWRNHOUNHCELACSBYYIM5", &cursor, func(o interface{}) {
		assertion := o.(map[string]interface{})
		if assertion["type"] == "manage_data" {
			var properType operations.ManageData
			jsonB, err := json.Marshal(o)
			if err != nil {
				return
			}
			json.Unmarshal(jsonB, &properType)
			fmt.Println(properType)
		}
	})
	if err != nil {
		c()
	}
}
