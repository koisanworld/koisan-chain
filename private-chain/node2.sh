#!/bin/bash

DATADIR=/tmp/geth-node2
KEYSTOREDIR=$DATADIR/keystore
mkdir -pv $KEYSTOREDIR

cp ./UTC--2021-07-28T07-10-49.020414527Z--81b0f221c286ebd5e1cffd7bc0d94ee8350da1c3 $KEYSTOREDIR

../build/bin/geth --testnet \
		--nousb \
		--port 30304 \
        --datadir $DATADIR \
        --syncmode 'full' \
		--gcmode 'archive' \
        --unlock '0x81B0F221c286ebd5E1CfFd7BC0d94Ee8350Da1C3' \
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
		--http.port 8080 \
	    --http.corsdomain '*' \
        --http.api 'admin,debug,web3,eth,txpool,personal,bpos,miner,net' \
		--ws \
		--ws.addr '0.0.0.0' \
		--ws.port 8080 \
	    --ws.origins '*' \
        --ws.api 'admin,debug,web3,eth,txpool,personal,bpos,miner,net' \
        --nodekeyhex 93ecaae869e5db671cbfcf13c567dbbb15529b85dc14d7ff0fcbc2587e7a3c2a \
        --bootnodes "enode://f987588021270aadcd19a79b2f45285150c9e7924322e40c50339c6f636ed37edb5e44b7afc243162e70064af3b61b26b0d1a69d0fab8ac91893230e7de5c562@127.0.0.1:32668"
