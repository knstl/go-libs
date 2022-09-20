package address

import (
	"regexp"

	"github.com/knstl/go-libs/primitives/types"
)

//nolint:gochecknoglobals
var regexpMap = map[types.ChainType]*regexp.Regexp{
	types.BSC:    regexp.MustCompile("^0x[0-9a-fA-F]{40}$"),
	types.KNSTL:  regexp.MustCompile("^darc1[0-9a-zA-Z]{38}$"),
	types.SOLANA: regexp.MustCompile("^*$"),
}

func GetRegexpByChain(chain types.ChainType) (*regexp.Regexp, bool) {
	r, ok := regexpMap[chain]

	return r, ok
}
