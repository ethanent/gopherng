# gopherng
> Simple PRNG using the SHA256 hash function

[![Go Reference](https://pkg.go.dev/badge/github.com/ethanent/gopherng.svg)](https://pkg.go.dev/github.com/ethanent/gopherng)

## Usage

```go
// initialize f64 PRNG with seed (use a longer one irl)
p := gopherng.NewFloat64PRNG([]byte{1, 2, 3, 4, 5})

// generate values!
v, err := p.Next()
if err != nil {/*handle error if any*/}
fmt.Printf("%f\n", v)

// ... keep using p.Next() to generate additional values
```

## Notes

This package relies upon `rand.Int()` from `crypto/rand` reading a
consistent number of bytes for each Int generated from a `gopherng.PRNGSource`
in order to generate consistent output from `Float64PRNG`s.

Use a large seed if you want to make the seed more difficult to figure when
exposing random values.
