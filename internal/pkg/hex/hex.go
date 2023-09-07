package hex

import (
	"fmt"
	"math/big"
	"strings"
)

type Provider interface {
	HexToInt(hex string) *big.Int
	IntToHex(num int64) string
}

type hexProvider struct {
}

func NewHexProvider() Provider {
	return hexProvider{}
}

func (h hexProvider) HexToInt(hex string) *big.Int {
	hex = strings.TrimPrefix(hex, "0x")
	bInt := new(big.Int)
	bInt.SetString(hex, 16)

	return bInt
}

func (h hexProvider) IntToHex(num int64) string {
	return fmt.Sprintf("0x%x", num)
}
