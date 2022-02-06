#!/bin/bash

DATADIR=/tmp/geth-node3
KEYSTOREDIR=$DATADIR/keystore
mkdir -pv $KEYSTOREDIR

cp ./UTC--2021-07-29T01-52-44.777814653Z--7b92b2801bbe0b776176d025fdcac49771fdafd9 $KEYSTOREDIR

../build/bin/geth --testnet \
		--nousb \
		--port 30305 \
        --datadir $DATADIR \
        --syncmode 'full' \
		--gcmode 'archive' \
        --unlock '0x7B92b2801BBE0b776176d025FdcaC49771FdAFD9' \
        --password ./password.txt \
        --allow-insecure-unlock \
        --mine \
		--networkid 888 \
        --miner.gaslimit  30000000 \
        --miner.gastarget  30000000 \
        --txpool.pricelimit 0 \
        --miner.gasprice 0 \
		--http \
		--http.addr '0.0.0.0' \
		--http.port 8081 \
	    --http.corsdomain '*' \
        --http.api 'admin,debug,web3,eth,txpool,personal,bpos,miner,net' \
		--ws \
		--ws.addr '0.0.0.0' \
		--ws.port 8081 \
	    --ws.origins '*' \
        --ws.api 'admin,debug,web3,eth,txpool,personal,bpos,miner,net' \
        --nodekeyhex 2ce35cdb85f7e608b2b145a51059e5de75eb0c50b6beb51da5728a0dbe26ff60 \
        --bootnodes "enode://f987588021270aadcd19a79b2f45285150c9e7924322e40c50339c6f636ed37edb5e44b7afc243162e70064af3b61b26b0d1a69d0fab8ac91893230e7de5c562@127.0.0.1:32668"
