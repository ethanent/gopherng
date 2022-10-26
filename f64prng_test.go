package gopherng_test

import (
	"math"
	"testing"

	"github.com/ethanent/gopherng"
)

func TestFloat64PRNG(t *testing.T) {
	t.Run("Distribution", func(t *testing.T) {
		p := gopherng.NewFloat64PRNG([]byte{3, 7, 8, 1, 231, 221})

		groupRanges := [][2]float64{
			{0.0, 0.1},
			{0.1, 0.2},
			{0.2, 0.3},
			{0.3, 0.4},
			{0.4, 0.5},
			{0.5, 0.6},
			{0.6, 0.7},
			{0.7, 0.8},
			{0.8, 0.9},
			{0.9, 1.0},
		}

		groups := map[[2]float64]int{}

		for _, gr := range groupRanges {
			groups[gr] = 0
		}

		total := 0

		for i := 0; i < 1000000; i++ {
			v, err := p.Next()
			if err != nil {
				t.Fatal(err)
			}
			total++
			for grange := range groups {
				if grange[0] < v && grange[1] > v {
					groups[grange]++
					continue
				}
			}
			if v == 1 {
				t.Fatal("v == 1")
			}
			if v < 0 || v > 1 {
				t.Fatal("v < 0 || v > 1")
			}
		}

		for _, grange := range groupRanges {
			c := groups[grange]
			portionInGroup := float64(c) / float64(total)
			t.Logf("%v: %f\n", grange, portionInGroup)
			leeway := 0.005
			if math.Abs(portionInGroup-0.1) > leeway {
				t.Fatal("Disproportionate count within this group.")
			}
		}
	})
}
