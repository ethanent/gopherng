package gopherng_test

import (
	"crypto/rand"
	"math"
	"math/big"
	"testing"

	"github.com/ethanent/gopherng"
)

func TestFloat64PRNG(t *testing.T) {
	t.Run("Distribution", func(t *testing.T) {
		buckets := 20
		for j := 0; j < 20; j++ {
			seed, err := randomSeed(j * 33)
			if err != nil {
				t.Fatal(err)
			}

			t.Logf("Seed (len=%d):\n", len(seed))

			p := gopherng.NewFloat64PRNG(seed)

			groupRanges := [][2]float64{}

			bucketWidth := 1 / float64(buckets)

			for i := 0; i < buckets; i++ {
				l := bucketWidth * float64(i)
				groupRanges = append(groupRanges, [2]float64{l, l + bucketWidth})
			}

			groups := map[[2]float64]int{}

			for _, gr := range groupRanges {
				groups[gr] = 0
			}

			total := 0

			for i := 0; i < 100000; i++ {
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
				leeway := 0.003
				gap := math.Abs(portionInGroup - bucketWidth)
				//t.Logf("%v: %f (%f)\n", grange, portionInGroup, gap)
				if gap > leeway {
					t.Fatalf("Disproportionate count within this bucket:\n  %v: %f\n", grange, portionInGroup)
				}
			}
			t.Logf("  OK\n")
		}
	})

	t.Run("Consistency", func(t *testing.T) {
		vcount := 100000
		v := make([]float64, vcount)

		for k := 0; k < 10; k++ {
			seed, err := randomSeed(k * 33)
			if err != nil {
				t.Fatal(err)
			}

			t.Logf("Seed (len=%d):\n", len(seed))

			for j := 0; j < 3; j++ {
				p := gopherng.NewFloat64PRNG(seed)

				for i := 0; i < vcount; i++ {
					vn, err := p.Next()
					if err != nil {
						t.Fatal(err)
					}
					if j != 0 && v[i] != vn {
						t.Fatalf("v1[i]=%f != v2[i]=%f\n", v[i], vn)
					}
					v[i] = vn
				}

				t.Logf("- OK %d\n", j)
			}
		}
	})
}

// randomSeed generates a random seed for PRNG.
// set size to -1 for random seed size.
func randomSeed(size int) ([]byte, error) {
	seedSize := size
	if size == -1 {
		seedSizeInt, err := rand.Int(rand.Reader, big.NewInt(512))
		if err != nil {
			return nil, err
		}
		seedSize = int(seedSizeInt.Int64())
	}
	s := make([]byte, seedSize)
	if _, err := rand.Read(s); err != nil {
		return nil, err
	}
	return s, nil
}
