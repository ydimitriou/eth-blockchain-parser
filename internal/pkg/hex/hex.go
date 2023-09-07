package hex

import (
	"fmt"
	"math/big"
	"strings"
)

// Provider interface for hex/int convertion
type Provider interface {
	HexToInt(hex string) *big.Int
	IntToHex(num int64) string
}

type hexProvider struct {
}

// NewHexProvider constructor to return the default hex provider
func NewHexProvider() Provider {
	return hexProvider{}
}

// HexToInt converts string to big Int
func (h hexProvider) HexToInt(hex string) *big.Int {
	hex = strings.TrimPrefix(hex, "0x")
	bInt := new(big.Int)
	bInt.SetString(hex, 16)

	return bInt
}

// IntToHex converts int64 to string
func (h hexProvider) IntToHex(num int64) string {
	return fmt.Sprintf("0x%x", num)
}
