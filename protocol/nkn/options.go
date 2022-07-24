/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package nkn

import (
	"encoding/hex"

	nkn "github.com/nknorg/nkn-sdk-go"
)

// Option is the function signature required to be considered an pubsub.Option.
type Option func(*Protocol) error

const ()

func WithWalletSeed(seed string) Option {
	return func(p *Protocol) error {
		s, err := hex.DecodeString(seed)
		if err != nil {
			return err
		}
		p.SenderAccount = nkn.NewAccount(s)
		return nil
	}
}

func WithConsumerAddress(address string) Option {
	return func(p *Protocol) error {
		p.ConsumerAddress = address
		return nil
	}
}
