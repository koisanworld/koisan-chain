package systemcontract

import (
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/require"
)

func TestJsonUnmarshalABI(t *testing.T) {
	for _, abiStr := range []string{
		validatorsInteractiveABI,
		punishInteractiveABI,
		proposalInteractiveABI,
		addrListInteractiveABI,
		incentiveInteractiveABI} {
		_, err := abi.JSON(strings.NewReader(abiStr))
		require.NoError(t, err, abiStr)
	}
}
