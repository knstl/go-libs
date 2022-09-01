package knstl

import (
	"bufio"
	"errors"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	cryptokeyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bip39 "github.com/cosmos/go-bip39"
)

var (
	ErrMnemonicEmpty   = errors.New("mnemonic is empty")
	ErrInvalidMnemonic = errors.New("invalid mnemonic")
)

const (
	Denom = "udarc"

	coinType = uint32(118)
	account  = uint32(0)
	index    = uint32(0)

	keyringDir = "."
	algoStr    = "secp256k1"

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32MainPrefix = "darc"
	// PrefixValidator is the prefix for validator keys
	PrefixValidator = "val"
	// PrefixConsensus is the prefix for consensus keys
	PrefixConsensus = "cons"
	// PrefixPublic is the prefix for public keys
	PrefixPublic = "pub"
	// PrefixOperator is the prefix for operator keys
	PrefixOperator = "oper"

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr = Bech32MainPrefix
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub = Bech32MainPrefix + PrefixPublic
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	Bech32PrefixValAddr = Bech32MainPrefix + PrefixValidator + PrefixOperator
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	Bech32PrefixValPub = Bech32MainPrefix + PrefixValidator + PrefixOperator + PrefixPublic
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	Bech32PrefixConsAddr = Bech32MainPrefix + PrefixValidator + PrefixConsensus
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	Bech32PrefixConsPub = Bech32MainPrefix + PrefixValidator + PrefixConsensus + PrefixPublic
)

func NewKeyring(serviceName string) (cryptokeyring.Keyring, error) {
	kb, err := cryptokeyring.New(serviceName, cryptokeyring.BackendMemory, keyringDir,
		bufio.NewReader(strings.NewReader("")))
	if err != nil {
		return nil, fmt.Errorf("failed to create a new keyring: %w", err)
	}

	return kb, nil
}

func CreateKey(kb cryptokeyring.Keyring, name, mnemonic string) (cryptokeyring.Info, error) {
	// RegisterBech32Prefix()
	cryptokeyringAlgos, _ := kb.SupportedAlgorithms()
	algo, err := cryptokeyring.NewSigningAlgoFromString(algoStr, cryptokeyringAlgos)
	if err != nil {
		return nil, fmt.Errorf("failed to create a supported signature algo: %w", err)
	}

	hdPath := hd.CreateHDPath(coinType, account, index).String()

	var bip39Passphrase string
	if mnemonic == "" {
		return nil, ErrMnemonicEmpty
	}

	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, ErrInvalidMnemonic
	}

	info, err := kb.NewAccount(name, mnemonic, bip39Passphrase, hdPath, algo)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new account: %w", err)
	}

	return info, nil
}

//nolint:gochecknoinits
func init() {
	config := sdk.GetConfig()

	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)

	config.Seal()
}
