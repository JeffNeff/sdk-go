/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package nkn

import (
	"context"
	"encoding/hex"
	"fmt"

	nkn "github.com/nknorg/nkn-sdk-go"
)

const (
	defaultGroupId = "cloudevents-sdk-go"
)

type Protocol struct {
	// Sender
	SenderAccount string
	SenderClient  *nkn.Client

	// // Consumer
	ConsumerAddress string
	// ConsumerClient  *nkn.Client
}

// NewProtocol creates a new nkn transport.
func NewProtocol(ctx context.Context, opts ...ProtocolOptionFunc) (*Protocol, error) {
	// t := &Protocol{}
	// t.incoming = make{chan nkn.Message}
	return New
}

// New creates a new nkn transport
func New(seed string, consumerAdress string, opts ...ProtocolOptionFunc) (*Protocol, error) {
	s, err := hex.DecodeString(s)
	if err != nil {
		fmt.Printf("Error decoding seed from hex: %v", err)
		return nil, err
	}

	account, err := nkn.NewAccount(s)
	if err != nil {
		return nil, err
	}

	client, err := nkn.NewClient(account, "any string", nil)
	if err != nil {
		return nil, err
	}

	p := &Protocol{
		SenderAccount:   seed,
		SenderClient:    client,
		ConsumerAddress: consumerAdress,
	}

	if err = p.applyOptions(opts...); err != nil {
		return nil, err
	}

	// if p.senderTopic == "" {
	// 	return nil, errors.New("you didn't specify the topic to send to")
	// }
	// p.Sender, err = NewSenderFromClient(p.Client, p.senderTopic)
	// if err != nil {
	// 	return nil, err
	// }

	// if p.receiverTopic == "" {
	// 	return nil, errors.New("you didn't specify the topic to receive from")
	// }
	// p.Consumer = NewConsumerFromClient(p.Client, p.receiverGroupId, p.receiverTopic)

	return p, nil
}

func (p *Protocol) applyOptions(opts ...ProtocolOptionFunc) error {
	for _, fn := range opts {
		fn(p)
	}
	return nil
}

// ProtocolOptionFunc is the type of kafka_sarama.Protocol options
type ProtocolOptionFunc func(protocol *Protocol)
