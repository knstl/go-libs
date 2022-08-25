package numbers

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/shopspring/decimal"

	"github.com/knstl/go-libs/primitives/types"
)

const (
	AmountBscUnit   = "1000000000000000000" // in wei (1 eth)
	AmountKnstlUnit = "1000000"             // in udarc (1 darc)
)

// AmountUnitByChain is a map for getting chain unit of amount.
var AmountUnitByChain = map[types.ChainType]string{ //nolint:gochecknoglobals
	types.BSC:   AmountBscUnit,
	types.KNSTL: AmountKnstlUnit,
}

// GetTxAmountByChain converts float string to blockchain unit amount.
// Example 1: [BSC] amount "0.0001" -> "100000000000000"
// Example 2: [KNSTL] amount "2.05" -> "2050000"
func GetTxAmountByChain(amount float64, chain types.ChainType) (decimal.Decimal, error) {
	amountDecimal := decimal.NewFromFloat(amount)

	amountChainUnit, ok := AmountUnitByChain[chain]
	if !ok {
		return amountDecimal, types.ErrChainNotSupported
	}

	amountChainUnitDecimal, err := decimal.NewFromString(amountChainUnit)
	if err != nil {
		return amountDecimal, fmt.Errorf("failed to get decimal from string (amountChainUnit=%s): %w", amountChainUnit, err)
	}

	return amountDecimal.Mul(amountChainUnitDecimal), nil
}

//nolint:goerr113
func DecimalToSatoshis(dec string) (string, error) {
	out := strings.TrimLeft(dec, " ")
	out = strings.TrimRight(out, " ")
	out = strings.Replace(out, ".", "", 1)

	// trim left 0's but keep last
	if l := len(out); l >= 2 {
		out = strings.TrimLeft(out[:l-1], "0") + out[l-1:l]
	}

	if len(out) == 0 {
		return "", errors.New("Invalid empty input: " + dec)
	}

	for _, c := range out {
		if !unicode.IsNumber(c) {
			return "", errors.New("not a number: " + dec)
		}
	}

	return out, nil
}
