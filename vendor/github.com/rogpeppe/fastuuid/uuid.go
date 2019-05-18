// Package fastuuid provides fast UUID generation of 192 bit
// universally unique identifiers.
//
// It also provides simple support for 128-bit RFC-4122-like UUID strings.
//
// Note that the generated UUIDs are not unguessable - each
// UUID generated from a Generator is adjacent to the
// previously generated UUID.
//
// It ignores RFC 4122.
package fastuuid

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"sync/atomic"
)

// Generator represents a UUID generator that
// generates UUIDs in sequence from a random starting
// point.
type Generator struct {
	// The constant seed. The first 8 bytes of this are
	// copied into counter and then ignored thereafter.
	seed    [24]byte
	counter uint64
}

// NewGenerator returns a new Generator.
// It can fail if the crypto/rand read fails.
func NewGenerator() (*Generator, error) {
	var g Generator
	_, err := rand.Read(g.seed[:])
	if err != nil {
		return nil, errors.New("cannot generate random seed: " + err.Error())
	}
	g.counter = binary.LittleEndian.Uint64(g.seed[:8])
	return &g, nil
}

// MustNewGenerator is like NewGenerator
// but panics on failure.
func MustNewGenerator() *Generator {
	g, err := NewGenerator()
	if err != nil {
		panic(err)
	}
	return g
}

// Next returns the next UUID from the generator.
// Only the first 8 bytes can differ from the previous
// UUID, so taking a slice of the first 16 bytes
// is sufficient to provide a somewhat less secure 128 bit UUID.
//
// It is OK to call this method concurrently.
func (g *Generator) Next() [24]byte {
	x := atomic.AddUint64(&g.counter, 1)
	uuid := g.seed
	binary.LittleEndian.PutUint64(uuid[:8], x)
	return uuid
}

// Hex128 is a convenience method that returns Hex128(g.Next()).
func (g *Generator) Hex128() string {
	return Hex128(g.Next())
}

// Hex128 returns an RFC4122-like representation of the
// first 128 bits of the given UUID. For example:
//
//	f81d4fae-7dec-11d0-a765-00a0c91e6bf6.
//
// It does not bother to set the version or variant bits
// of the UUID, as they only serve to reduce the randomness.
//
// If you want unpredictable UUIDs, you might want to consider
// hashing the uuid (using SHA256, for example) before passing it
// to Hex128.
func Hex128(uuid [24]byte) string {
	b := make([]byte, 36)
	hex.Encode(b[0:8], uuid[0:4])
	b[8] = '-'
	hex.Encode(b[9:13], uuid[4:6])
	b[13] = '-'
	hex.Encode(b[14:18], uuid[6:8])
	b[18] = '-'
	hex.Encode(b[19:23], uuid[8:10])
	b[23] = '-'
	hex.Encode(b[24:], uuid[10:16])
	return string(b)
}

// ValidHex128 reports whether id is a valid UUID as returned by Hex128
// and various other UUID packages, such as github.com/satori/go.uuid's
// NewV4 function.
//
// Note that it does not allow upper case hex.
func ValidHex128(id string) bool {
	if len(id) != 36 {
		return false
	}
	if id[8] != '-' || id[13] != '-' || id[18] != '-' || id[23] != '-' {
		return false
	}
	return isValidHex(id[0:8]) &&
		isValidHex(id[9:13]) &&
		isValidHex(id[14:18]) &&
		isValidHex(id[19:23]) &&
		isValidHex(id[24:])
}

func isValidHex(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !('0' <= c && c <= '9' || 'a' <= c && c <= 'f') {
			return false
		}
	}
	return true
}
