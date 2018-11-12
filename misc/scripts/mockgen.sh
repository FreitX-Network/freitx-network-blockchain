#!/bin/bash

rm -rf ./test/mock
mkdir -p ./test/mock

mkdir -p ./test/mock/mock_dispatcher
mockgen -destination=./test/mock/mock_dispatcher/mock_dispatcher.go  \
        -source=./dispatch/dispatcher/dispatcher.go \
        -package=mock_dispatcher \
        Dispatcher

mkdir -p ./test/mock/mock_blockchain
mockgen -destination=./test/mock/mock_blockchain/mock_blockchain.go  \
        -source=./blockchain/blockchain.go \
        -imports =github.com/freitx-project/freitx-network-blockchain/blockchain \
        -package=mock_blockchain \
        Blockchain

mkdir -p ./test/mock/mock_blocksync
mockgen -destination=./test/mock/mock_blocksync/mock_blocksync.go  \
        -source=./blocksync/blocksync.go \
        -package=mock_blocksync \
        BlockSync

mkdir -p ./test/mock/mock_trie
mockgen -destination=./test/mock/mock_trie/mock_trie.go  \
        -source=./trie/trie.go \
        -package=mock_trie \
        Trie

mkdir -p ./test/mock/mock_state
mockgen -destination=./test/mock/mock_state/mock_state.go  \
        -source=./state/factory.go \
        -imports =github.com/freitx-project/freitx-network-blockchain/state \
        -package=mock_state \
        Factory

mkdir -p ./test/mock/mock_consensus
mockgen -destination=./test/mock/mock_consensus/mock_consensus.go  \
        -source=./consensus/consensus.go \
        -imports =github.com/freitx-project/freitx-network-blockchain/consensus \
        -package=mock_consensus \
        Consensus

mkdir -p ./test/mock/mock_network
mockgen -destination=./test/mock/mock_network/mock_overlay.go  \
        -source=./network/overlay.go \
        -package=mock_network \
        Overlay

mkdir -p ./test/mock/mock_lifecycle
mockgen -destination=./test/mock/mock_lifecycle/mock_lifecycle.go \
        github.com/freitx-project/freitx-network-blockchain/pkg/lifecycle StartStopper

mkdir -p ./test/mock/mock_actpool
mockgen -destination=./test/mock/mock_actpool/mock_actpool.go  \
        -source=./actpool/actpool.go \
        -package=mock_actpool \
        ActPool