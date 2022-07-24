/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package nkn

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/JeffNeff/sdk-go/v2/binding"
	"github.com/JeffNeff/sdk-go/v2/protocol"
	nkn "github.com/nknorg/nkn-sdk-go"
)

type Sender struct {
	nknClient  *nkn.Client
	nknAccount *nkn.Account
	sinkAddres string
}

func (s *Sender) Send(ctx context.Context, in binding.Message, transformers ...binding.Transformer) error {
	response, err := s.nknClient.Send(nkn.NewStringArray(s.sinkAddres), transformers, nil)
	if err != nil {
		// return a.replier.Error(&event, targetce.ErrorCodeAdapterProcess, err, "error sending message")
	}

	return err
}

// NewSender creates a new Sender which wraps an amqp.Sender in a binding.Sender
func NewSender(s string, options ...SenderOptionFunc) protocol.Sender {
	seed, err := hex.DecodeString(s)
	if err != nil {
		fmt.Printf("Error decoding seed from hex: %v", err)
		return nil
	}

	account, err := nkn.NewAccount(seed)
	if err != nil {
		fmt.Printf("Error creating NKN account from seed: %v", err)
		return nil
	}

	client, err := nkn.NewClient(account, "any string", nil)
	if err != nil {
		fmt.Printf("Error creating NKN client: %v", err)
		return nil
	}

	return client
}

type SenderOptionFunc func(sender *Sender)
