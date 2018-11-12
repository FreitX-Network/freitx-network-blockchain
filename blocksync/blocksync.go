// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package blocksync

import (
	"context"

	"github.com/pkg/errors"

	"github.com/freitx-project/freitx-network-blockchain/actpool"
	"github.com/freitx-project/freitx-network-blockchain/blockchain"
	"github.com/freitx-project/freitx-network-blockchain/config"
	"github.com/freitx-project/freitx-network-blockchain/logger"
	"github.com/freitx-project/freitx-network-blockchain/network"
	"github.com/freitx-project/freitx-network-blockchain/network/node"
	"github.com/freitx-project/freitx-network-blockchain/pkg/lifecycle"
	pb "github.com/freitx-project/freitx-network-blockchain/proto"
)

// BlockSync defines the interface of blocksyncer
type BlockSync interface {
	lifecycle.StartStopper

	P2P() network.Overlay
	ProcessSyncRequest(sender string, sync *pb.BlockSync) error
	ProcessBlock(blk *blockchain.Block) error
	ProcessBlockSync(blk *blockchain.Block) error
}

// blockSyncer implements BlockSync interface
type blockSyncer struct {
	ackBlockCommit bool // acknowledges latest committed block
	ackBlockSync   bool // acknowledges old block from sync request
	ackSyncReq     bool // acknowledges incoming Sync request
	buf            *blockBuffer
	worker         *syncWorker
	bc             blockchain.Blockchain
	p2p            network.Overlay
}

// NewBlockSyncer returns a new block syncer instance
func NewBlockSyncer(
	cfg *config.Config,
	chain blockchain.Blockchain,
	ap actpool.ActPool,
	p2p network.Overlay,
) (BlockSync, error) {
	if cfg == nil || chain == nil || ap == nil || p2p == nil {
		return nil, errors.New("cannot create BlockSync: missing param")
	}
	startHeight, err := findSyncStartHeight(chain)
	if err != nil {
		return nil, err
	}
	buf := &blockBuffer{
		blocks:          make(map[uint64]*blockchain.Block),
		bc:              chain,
		ap:              ap,
		size:            cfg.BlockSync.BufferSize,
		startHeight:     startHeight,
		confirmedHeight: startHeight - 1,
	}
	w := newSyncWorker(cfg, p2p, buf)
	return &blockSyncer{
		ackBlockCommit: cfg.IsDelegate() || cfg.IsFullnode(),
		ackBlockSync:   cfg.IsDelegate() || cfg.IsFullnode(),
		ackSyncReq:     cfg.IsDelegate() || cfg.IsFullnode(),
		bc:             chain,
		buf:            buf,
		p2p:            p2p,
		worker:         w,
	}, nil
}

// P2P returns the network overlay object
func (bs *blockSyncer) P2P() network.Overlay {
	return bs.p2p
}

// Start starts a block syncer
func (bs *blockSyncer) Start(ctx context.Context) error {
	logger.Debug().Msg("Starting block syncer")
	return bs.worker.Start(ctx)
}

// Stop stops a block syncer
func (bs *blockSyncer) Stop(ctx context.Context) error {
	logger.Debug().Msg("Stopping block syncer")
	return bs.worker.Start(ctx)
}

// ProcessBlock processes an incoming latest committed block
func (bs *blockSyncer) ProcessBlock(blk *blockchain.Block) error {
	if !bs.ackBlockCommit {
		// node is not meant to handle latest committed block, simply exit
		return nil
	}

	var needSync bool
	moved, re := bs.buf.Flush(blk)
	switch re {
	case bCheckinLower:
		logger.Debug().Msg("Drop block lower than buffer's accept height.")
	case bCheckinExisting:
		logger.Debug().Msg("Drop block exists in buffer.")
	case bCheckinHigher:
		needSync = true
	case bCheckinValid:
		needSync = !moved
	}

	if needSync {
		bs.worker.SetTargetHeight(blk.Height())
	}
	return nil
}

func (bs *blockSyncer) ProcessBlockSync(blk *blockchain.Block) error {
	if !bs.ackBlockSync {
		// node is not meant to handle sync block, simply exit
		return nil
	}
	bs.buf.Flush(blk)
	return nil
}

// ProcessSyncRequest processes a block sync request
func (bs *blockSyncer) ProcessSyncRequest(sender string, sync *pb.BlockSync) error {
	if !bs.ackSyncReq {
		// node is not meant to handle sync request, simply exit
		return nil
	}

	for i := sync.Start; i <= sync.End; i++ {
		blk, err := bs.bc.GetBlockByHeight(i)
		if err != nil {
			return err
		}
		// TODO: send back multiple blocks in one shot
		if err := bs.p2p.Tell(node.NewTCPNode(sender), &pb.BlockContainer{Block: blk.ConvertToBlockPb()}); err != nil {
			logger.Warn().Err(err).Msg("Failed to response to ProcessSyncRequest.")
		}
	}
	return nil
}
