// Copyright (c) FreitX Network
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package dispatcher

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"

	"github.com/freitx-project/freitx-network-blockchain/config"
	"github.com/freitx-project/freitx-network-blockchain/network/node"
	pb "github.com/freitx-project/freitx-network-blockchain/proto"
)

func createDispatcher(t *testing.T, chainID uint32) Dispatcher {
	cfg := &config.Config{
		Consensus:  config.Consensus{Scheme: config.NOOPScheme},
		Dispatcher: config.Dispatcher{EventChanSize: 1024},
	}
	dp, err := NewDispatcher(cfg)
	assert.NoError(t, err)
	dp.AddSubscriber(chainID, &DummySubscriber{})
	return dp
}

func startDispatcher(t *testing.T) (ctx context.Context, d Dispatcher) {
	ctx = context.Background()
	d = createDispatcher(t, config.Default.Chain.ID)
	assert.NotNil(t, d)
	err := d.Start(ctx)
	assert.NoError(t, err)
	return
}

func stopDispatcher(ctx context.Context, d Dispatcher, t *testing.T) {
	err := d.Stop(ctx)
	assert.NoError(t, err)
}

func setTestCase() []proto.Message {
	return []proto.Message{
		&pb.ActionPb{},
		&pb.ProposePb{},
		&pb.EndorsePb{},
		&pb.BlockPb{},
		&pb.BlockSync{},
		&pb.BlockContainer{},
		&pb.BlockContainer{Block: &pb.BlockPb{}},
		&pb.TestPayload{},
	}
}

func TestHandleBroadcast(t *testing.T) {
	msgs := setTestCase()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx, d := startDispatcher(t)
	defer stopDispatcher(ctx, d, t)

	done := make(chan bool, 1000)
	for i := 0; i < 100; i++ {
		for _, msg := range msgs {
			d.HandleBroadcast(config.Default.Chain.ID, msg, done)
		}
	}
}

func TestHandleTell(t *testing.T) {
	msgs := setTestCase()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx, d := startDispatcher(t)
	defer stopDispatcher(ctx, d, t)

	done := make(chan bool, 1000)
	for i := 0; i < 100; i++ {
		for _, msg := range msgs {
			d.HandleTell(config.Default.Chain.ID, node.NewTCPNode("192.168.0.0:10000"), msg, done)
		}
	}
}

type DummySubscriber struct {
}

func (s *DummySubscriber) HandleBlock(*pb.BlockPb) error {
	return nil
}

func (s *DummySubscriber) HandleBlockSync(*pb.BlockPb) error {
	return nil
}

func (s *DummySubscriber) HandleSyncRequest(string, *pb.BlockSync) error {
	return nil
}

func (s *DummySubscriber) HandleAction(*pb.ActionPb) error {
	return nil
}

func (s *DummySubscriber) HandleBlockPropose(*pb.ProposePb) error {
	return nil
}

func (s *DummySubscriber) HandleEndorse(*pb.EndorsePb) error {
	return nil
}
