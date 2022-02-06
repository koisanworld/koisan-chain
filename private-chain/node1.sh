#!/bin/bash

DATADIR=/tmp/geth-node1
KEYSTOREDIR=$DATADIR/keystore
mkdir -pv $KEYSTOREDIR

cp ./UTC--2021-07-28T07-10-15.748605350Z--e769490012e89a7701a817a789478320e3cb211f $KEYSTOREDIR

../build/bin/geth --testnet \
		--nousb \
        --datadir $DATADIR \
        --syncmode 'full' \
		--gcmode 'archive' \
        --unlock '0xE769490012E89A7701A817a789478320E3Cb211F' \
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
		--http.port 8545 \
	    --http.corsdomain '*' \
        --http.api 'admin,debug,web3,eth,txpool,personal,bpos,miner,net' \
		--ws \
		--ws.addr '0.0.0.0' \
		--ws.port 8545 \
	    --ws.origins '*' \
        --ws.api 'admin,debug,web3,eth,txpool,personal,bpos,miner,net' \
        --nodekeyhex 8e963ce1c89e513eba029734005ec4976dd5f82fa9be0cb359881899c89beed8 

