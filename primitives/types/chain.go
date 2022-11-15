package types

import "errors"

var ErrChainNotSupported = errors.New("chain is not supported")

type ChainType string

const (
	KNSTL    = ChainType("KNSTL")
	BSC      = ChainType("BSC")
	SOLANA   = ChainType("SOL")
	ETHEREUM = ChainType("ETH")
)

//nolint:gochecknoglobals
var ChainTypes = map[ChainType]bool{
	KNSTL:    true,
	BSC:      true,
	SOLANA:   true,
	ETHEREUM: true,
}

func IsChainSupported(chain ChainType) bool {
	_, exists := ChainTypes[chain]

	return exists
}
