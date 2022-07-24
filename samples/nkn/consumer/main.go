/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"context"
	"log"
	"os"

	cenkn "github.com/cloudevents/sdk-go/protocol/nkn"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	WalletSeed string `envconfig:"WALLET_SEED" required:"true"`

	Sink string `envconfig:"K_SINK" required:"true"`
}

// Example is a basic data struct.
type Example struct {
	Sequence int    `json:"id"`
	Message  string `json:"message"`
}

func main() {
	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		os.Exit(1)
	}

	t, err := cenkn.New(env.Sink, env.WalletSeed)
	if err != nil {
		log.Printf("failed to create nkn transport, %s", err.Error())
		os.Exit(1)
	}
	c, err := cloudevents.NewClient(t, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		log.Printf("failed to create client, %s", err.Error())
		os.Exit(1)
	}

	event := cloudevents.NewEvent()
	event.SetType("com.cloudevents.sample.sent")
	event.SetSource("github.com/cloudevents/sdk-go/samples/nkn/sender/")
	_ = event.SetData("application/json", &Example{
		Sequence: 0,
		Message:  "HELLO",
	})

	if result := c.Send(context.Background(), event); cloudevents.IsUndelivered(result) {
		log.Printf("failed to send: %v", err)
		os.Exit(1)
	} else {
		log.Printf("sent, accepted: %t", cloudevents.IsACK(result))
	}

	os.Exit(0)
}
