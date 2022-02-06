#!/bin/bash

export COIN=TOKEN
export ETHEREUM_JSONRPC_VARIANT=geth
export ETHEREUM_JSONRPC_HTTP_URL="http://127.0.0.1:8545"
export ETHEREUM_JSONRPC_WS_URL="ws://127.0.0.1:8545"
export ETHEREUM_JSONRPC_TRACE_URL="http://127.0.0.1:8545"
export BLOCK_TRANSFORMER=clique
export NETWORK="Private Chain"
export MIX_ENV=prod
export PORT=80
export PG_DATADIR=/tmp/geth-private-chain-explorer/postgress/data

make start
