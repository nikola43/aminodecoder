package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	//"sync"

	"github.com/binance-chain/go-sdk/client/rpc"
	types "github.com/binance-chain/go-sdk/common/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/codec"
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

		//txBytes, err := base64.StdEncoding.DecodeString(eventData.Tx)
		txBytes, err := base64.StdEncoding.DecodeString("CowGCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMTkyM3B6MDNtaGphenRnY3YzZ2V5MGhqMGFtd3gwMmR5c2thdTUyCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMXh3ZDNrOHh0ZXJkZWZ0M254cWc5MnN6aHB6NnZ4NDNxc3BkcHc2CpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMTh5bGNoZ214eXBodzNjdHNsNzVuNTN1amVxdWttbWFnMm42eDNmCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMWZzOHI2enhtcjVuYzg2ajhjcGNtam1jY2Y4czJjYWZ4enQ1YWxxCpkBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEl4KK3Rjcm8xNG01YTRreHQyZTgydXFxczVndHF6YTI5ZG01d3F6eWEyanc5c2gSL3Rjcm9jbmNsMWZqYTVuc3h6N2dzcXc0emNjdXV5OHI3cGpuam1jN2RzY2RsMnZ6EmwKUQpGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQL6Inwmwwd0nDUwtu9S8U0E+TU86f92eeo/ZUJfq+O1tRIECgIIARieHRIXChEKCGJhc2V0Y3JvEgUxNTAwMBDwkwkaQDfjeHPhtkdkDG3JyqECSN2DTIZeTC3Z2dK82HL1qshIH6dvMvT2JP4NGhmcQW/JK97sZ+FMdxe98GJxQNLfZfk=")
		if err != nil {
			fmt.Println("Failed to decode transaction: %v", err)
		}

		fmt.Println("TxBytes:", txBytes)

		cdc := codec.New()



		var tx sdk.Tx
		fmt.Println("Tx:", tx)

		// decode txBytes into tx
		err = cdc.UnmarshalBinaryLengthPrefixed(txBytes, &tx)
		if err != nil {
			fmt.Println("Failed to unmarshal transaction: %v", err)
		}
		fmt.Printf("Decoded Transaction: %+v\n", tx)

		/*
		err = cdc.UnmarshalBinaryBare(txBytes, &tx)
		if err != nil {
			fmt.Println("Failed to unmarshal transaction: %v", err)
		}
		fmt.Printf("Decoded Transaction: %+v\n", tx)
		*/

	}
}
