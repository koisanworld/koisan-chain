#!/bin/bash

DATADIR=/tmp/geth-node4
KEYSTOREDIR=$DATADIR/keystore
mkdir -pv $KEYSTOREDIR

cp ./UTC--2021-07-29T01-59-58.348392221Z--cff2fa6784c5b69a61eb2c7d493cba68db5c17e8 $KEYSTOREDIR

../build/bin/geth --testnet \
		--nousb \
		--port 30306 \
        --datadir $DATADIR \
        --syncmode 'full' \
		--gcmode 'archive' \
        --unlock '0xCff2fa6784C5B69a61eB2C7d493Cba68DB5C17e8' \
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
		--http.port 8082 \
	    --http.corsdomain '*' \
        --http.api 'admin,debug,web3,eth,txpool,personal,bpos,miner,net' \
		--ws \
		--ws.addr '0.0.0.0' \
		--ws.port 8082 \
	    --ws.origins '*' \
        --ws.api 'admin,debug,web3,eth,txpool,personal,bpos,miner,net' \
        --nodekeyhex 4065303eb241365748f33062712eb252913492d2d2ae47403ae17dbf88017ba0 \
        --bootnodes "enode://f987588021270aadcd19a79b2f45285150c9e7924322e40c50339c6f636ed37edb5e44b7afc243162e70064af3b61b26b0d1a69d0fab8ac91893230e7de5c562@127.0.0.1:32668"
