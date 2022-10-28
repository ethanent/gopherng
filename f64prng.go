package gopherng

import (
	"crypto/rand"
	"math/big"
)

type Float64PRNG struct {
	p *PRNGSource
}

func NewFloat64PRNG(seed []byte) *Float64PRNG {
	f := &Float64PRNG{}
	p := NewPRNGSource(seed)
	f.p = p
	return f
}

func (f *Float64PRNG) Next() (float64, error) {
	// Based on
	// https://cs.opensource.google/go/go/+/refs/tags/go1.19.2:src/math/rand/rand.go;l=179.
	// See license for Go 1.19.2 in NOTICE.
	n, err := rand.Int(f.p, big.NewInt(1<<53))
	if err != nil {
		return 0, err
	}
	return float64(n.Int64()) / (1 << 53), nil
}
