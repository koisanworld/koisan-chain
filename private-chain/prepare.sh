#!/bin/bash

sed -i '11s/0x2d2d68a59880eaacf24b96731956d287482dac4b/0xE769490012E89A7701A817a789478320E3Cb211F/i' ../params/system_contract.go
sed -i '17s/0941A01ab7B3A39Ed6f55d6a4907778a3f15E5c9/E769490012E89A7701A817a789478320E3Cb211F/i' ../params/system_contract.go

cd ..
make geth
