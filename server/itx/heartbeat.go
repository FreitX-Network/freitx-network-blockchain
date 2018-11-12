// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package itx

import (
	"encoding/json"
	"time"

	"github.com/zjshen14/go-fsm"

	"github.com/freitx-project/freitx-network-blockchain/consensus"
	"github.com/freitx-project/freitx-network-blockchain/consensus/scheme/rolldpos"
	"github.com/freitx-project/freitx-network-blockchain/dispatch"
	"github.com/freitx-project/freitx-network-blockchain/logger"
	"github.com/freitx-project/freitx-network-blockchain/network"
)

// TODO: HeartbeatHandler opens encapsulation of a few structs to inspect the internal status, we need to find a better
// approach to do so in the future

// HeartbeatHandler is the handler to periodically log the system key metrics
type HeartbeatHandler struct {
	s *Server
}

// NewHeartbeatHandler instantiates a HeartbeatHandler instance
func NewHeartbeatHandler(s *Server) *HeartbeatHandler {
	return &HeartbeatHandler{s: s}
}

// Log executes the logging logic
func (h *HeartbeatHandler) Log() {
	// Network metrics
	p2p, ok := h.s.P2P().(*network.IotxOverlay)
	if !ok {
		logger.Error().Msg("value is not the instance of IotxOverlay")
		return
	}
	numPeers := network.LenSyncMap(p2p.PM.Peers)
	lastOutTime := time.Unix(0, 0)
	p2p.PM.Peers.Range(func(_, value interface{}) bool {
		p, ok := value.(*network.Peer)
		if !ok {
			logger.Error().Msg("value is not the instance of Peer")
			return true
		}
		if p.LastResTime.After(lastOutTime) {
			lastOutTime = p.LastResTime
		}
		return true
	})
	lastInTime := p2p.RPC.LastReqTime()

	// Dispatcher metrics
	dp, ok := h.s.Dispatcher().(*dispatch.IotxDispatcher)
	if !ok {
		logger.Error().Msg("dispatcher is not the instance of IotxDispatcher")
		return
	}
	numDPEvts := len(*dp.EventChan())
	dpEvtsAudit, err := json.Marshal(dp.EventAudit())
	if err != nil {
		logger.Error().Msg("error when serializing the dispatcher event audit map")
		return
	}

	// Consensus metrics
	cs, ok := h.s.consensus.(*consensus.IotxConsensus)
	if !ok {
		logger.Error().Msg("consensus is not the instance of IotxConsensus")
		return
	}
	rolldpos, ok := cs.Scheme().(*rolldpos.RollDPoS)
	numPendingEvts := 0
	var state fsm.State
	if ok {
		numPendingEvts = rolldpos.NumPendingEvts()
		state = rolldpos.CurrentState()
	} else {
		logger.Debug().Msg("scheme is not the instance of RollDPoS")
	}

	// Block metrics
	height, err := h.s.Blockchain().TipHeight()
	if err != nil {
		logger.Error().Err(err).Msg("error one getting the the blockchain height")
		height = 0
	}

	logger.Info().
		Uint("num-peers", numPeers).
		Time("last-out", lastOutTime).
		Time("last-in", lastInTime).
		Int("dispatcher-events", numDPEvts).
		Str("dispatcher-events-audit", string(dpEvtsAudit)).
		Int("rolldpos-events", numPendingEvts).
		Str("fsm-state", string(state)).
		Uint64("height", height).
		Msg("node status")
}