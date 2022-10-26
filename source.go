package gopherng

import (
	"crypto/sha256"
)

type PRNGSource struct {
	salt   []byte
	buffer []byte
	curIdx int
}

func NewPRNGSource(seed []byte) *PRNGSource {
	var p PRNGSource
	p.curIdx = 0
	sh32b := sha256.Sum256(seed)
	sh32b2 := sha256.Sum256(sh32b[:])
	p.salt = sh32b[:]
	p.buffer = sh32b2[:]
	return &p
}

func (p *PRNGSource) nextBuf() {
	sh32bn := sha256.Sum256(append(p.buffer, p.salt...))
	p.buffer = sh32bn[:]
	p.curIdx = 0
}

func (p *PRNGSource) nextByte() byte {
	if p.curIdx == len(p.buffer) {
		p.nextBuf()
	}
	v := p.buffer[p.curIdx]
	p.curIdx++
	return v
}

// Read reads random bytes. It should always read len(b) bytes without error.
func (p *PRNGSource) Read(b []byte) (int, error) {
	c := len(b)
	for i := 0; i < c; i++ {
		b[i] = p.nextByte()
	}
	return c, nil
}
