package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/bnb-chain/go-sdk/client/rpc"
	"github.com/bnb-chain/go-sdk/common/types"
	"github.com/gogo/protobuf/codec"
)

func main() {
	type EventData struct {
		Tx string `json:"tx"`
	}

	c := rpc.NewRPCClient("dataseed1.bnbchain.org:80", types.ProdNetwork)

	query := "tm.event = 'Tx'"
	events, err := c.Subscribe(query, 10)
	if err != nil {
		fmt.Println("websocket client", err)
		return
	}

	for event := range events {
		//fmt.Println(event.Data)

		eventDataBytes, err := json.Marshal(event.Data)
		if err != nil {
			fmt.Println("Error converting TMEventData to bytes:", err)
			return
		}

		var eventData EventData
		err = json.Unmarshal(eventDataBytes, &eventData)
		if err != nil {
			fmt.Println("Error unmarshaling event data:", err)
			return
		}

		fmt.Println("Tx:", eventData.Tx)

		txBytes, err := base64.StdEncoding.DecodeString(eventData.Tx)
		if err != nil {
			fmt.Println("Failed to decode transaction: %v", err)
		}
		cdc := codec.New()
		var tx sdk.Tx
		err = cdc.UnmarshalBinaryBare(txBytes, &tx)
		if err != nil {
			fmt.Println("Failed to unmarshal transaction: %v", err)
		}
		fmt.Printf("Decoded Transaction: %+v\n", tx)
	}
}
