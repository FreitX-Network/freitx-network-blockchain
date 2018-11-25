#!/bin/bash

file=$GOPATH/src/github.com/freitx-project/freitx-network-blockchain/sampleconfig/stonevan/config_$1.yaml
echo "Starting node with config file:" $file
$GOPATH/src/github.com/freitx-project/freitx-network-blockchain/bin/server -stderrthreshold=WARNING -log_dir=./log -config=$file

exit 0