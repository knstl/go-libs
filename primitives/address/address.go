package address

import (
	"encoding/hex"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/crypto/sha3"
)

// Decode decodes a hex string with 0x prefix.
func Remove0x(input string) string {
	if strings.HasPrefix(input, "0x") {
		return input[2:]
	}

	return input
}

// Hex returns an EIP55-compliant hex string representation of the address.
func EIP55Checksum(unchecksummed string) (string, error) {
	v := []byte(Remove0x(strings.ToLower(unchecksummed)))

	_, err := hex.DecodeString(string(v))
	if err != nil {
		return "", fmt.Errorf("failed to decode string: %w", err)
	}

	sha := sha3.NewLegacyKeccak256()
	if _, err = sha.Write(v); err != nil {
		return "", fmt.Errorf("failed to write hash: %w", err)
	}

	hash := sha.Sum(nil)

	result := v
	for i := 0; i < len(result); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte >>= 4
		} else {
			hashByte &= 0xf
		}
		if result[i] > '9' && hashByte > 7 {
			result[i] -= 32
		}
	}

	val := string(result)

	return "0x" + val, nil
}

func GetValAddressAndConsAddress(accAddress string) (string, string, error) {
	acc, err := sdk.AccAddressFromBech32(accAddress)
	if err != nil {
		return "", "", fmt.Errorf("failed to create an AccAddress from a Bech32: %w", err)
	}

	hex := fmt.Sprintf("%+v", acc)

	valAddress, err := sdk.ValAddressFromHex(hex)
	if err != nil {
		return "", "", fmt.Errorf("failed to create a Val from hex: %w", err)
	}

	consAddress, err := sdk.ConsAddressFromHex(hex)
	if err != nil {
		return "", "", fmt.Errorf("failed to create ConsAddr from hex: %w", err)
	}

	return valAddress.String(), consAddress.String(), nil
}
