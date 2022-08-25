package numbers

import (
	"testing"

	"github.com/shopspring/decimal"

	"github.com/knstl/go-libs/primitives/types"
)

func TestGetTxAmountByChain(t *testing.T) {
	type args struct {
		amount float64
		chain  types.ChainType
	}

	tests := []struct {
		name    string
		args    args
		want    decimal.Decimal
		wantErr bool
	}{
		{"test BSC (less than 1)", args{0.0001, types.BSC}, decimal.NewFromFloat(100000000000000), false},
		{"test BSC (more 1)", args{10.12345, types.BSC}, decimal.NewFromFloat(10123450000000000000), false},
		{"test KNSTL (less than 1 (1))", args{0.0001, types.KNSTL}, decimal.NewFromFloat(100), false},
		{"test KNSTL (less than 1 (2))", args{0.00001, types.KNSTL}, decimal.NewFromFloat(10), false},
		{"test KNSTL (less than 1 (3))", args{0.000001, types.KNSTL}, decimal.NewFromFloat(1), false},
		{"test KNSTL (less than 1 (4))", args{0.0000001, types.KNSTL}, decimal.NewFromFloat(0.1), false},
		{"test KNSTL (more 1)", args{2.05, types.KNSTL}, decimal.NewFromFloat(2050000), false},
		{"test Unsupported", args{0, types.ChainType("Unsupported_chain")}, decimal.NewFromFloat(0), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTxAmountByChain(tt.args.amount, tt.args.chain)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTxAmountByChain() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !got.Equals(tt.want) {
				t.Errorf("GetTxAmountByChain() got = %v, want %v", got, tt.want)
			}
		})
	}
}
