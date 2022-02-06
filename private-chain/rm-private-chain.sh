#!/bin/bash

docker rm -f blockscout postgres >/dev/null 2>&1
rm -rf /tmp/geth-private-chain-explorer

rm -rf /tmp/geth-node1
rm -rf /tmp/geth-node2
rm -rf /tmp/geth-node3
rm -rf /tmp/geth-node4