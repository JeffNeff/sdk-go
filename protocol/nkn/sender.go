/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package nkn

import (
	"context"

	"github.com/JeffNeff/sdk-go/v2/binding"
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

type SenderOptionFunc func(sender *Sender)
