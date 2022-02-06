package bpos

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// current block subsidy is 2.345
// go test -v -run ^TestCalcTotalSupply$ github.com/ethereum/go-ethereum/consensus/bpos
func TestCalcTotalSupply(t *testing.T) {
	totalSupply := new(big.Int).Mul(big.NewInt(65000000), big.NewInt(1e18))

	for i := uint64(0); ; i++ {
		blockSubsidy := calcBlockSubsidy(i)
		if blockSubsidy.Cmp(big.NewInt(0)) == 0 {
			break
		}
		totalSupply = new(big.Int).Add(totalSupply, blockSubsidy)
	}
	t.Log("Total supply is: ", totalSupply.String())
	if totalSupply.Cmp(new(big.Int).Mul(big.NewInt(100000000), big.NewInt(1e18))) != 0 {
		t.Errorf("Total supply should be 100000000000000000000000000 ,but get %s\n", totalSupply.String())
	}
}

func TestCalcBlockSubsidy(t *testing.T) {
	type test struct {
		height uint64
		want   *big.Int
	}
	tests := make([]test, 5)
	tests[0] = test{0, blockSubsidy}
	tests[1] = test{1, blockSubsidy}
	tests[2] = test{subsidyEndHeight - 1, blockSubsidy}
	tests[3] = test{subsidyEndHeight, big.NewInt(0)}
	tests[4] = test{subsidyEndHeight + 1, big.NewInt(0)}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.height), func(t *testing.T) {
			if got := calcBlockSubsidy(tt.height); got.Cmp(tt.want) != 0 {
				t.Errorf("hieght= %v, CalcBlockReward() = %v, want %v", tt.height, got, tt.want)
			}
		})
	}
}

func TestCalcSlotOfDevMappingKey(t *testing.T) {
	addr := common.HexToAddress("0x5b38da6a701c568545dcfcb03fcb875f56beddc4")
	slot := calcSlotOfDevMappingKey(addr)
	t.Log(slot.String())
	// want: 0xb314f101a00aa0d8cc6704cc6dd1e9dd7551ec98c9df52079c192c560ba66c4a

}
