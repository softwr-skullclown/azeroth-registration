package auth

import (
	"crypto/rand"
	"crypto/sha1"
	"math/big"
	"strings"
)

// salt generates a 32 byte salt used in creating the login hash
func salt() ([]byte, error) {
	salt := make([]byte, 32) // 32 bytes for salt

	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	return salt, nil
}

// calculateSRP6Verifier generates the verifier hash for the salt/hash combo of storing the username and password
// currently targeting AzerothCore database model
func calculateSRP6Verifier(username, password string, salt []byte) []byte {
	// Algorithm constants
	g := big.NewInt(7)
	N, _ := new(big.Int).SetString("894B645E89E1535BBDAD5B8B290650530801B18EBFBF5E8FAB3C82872A3E9BB7", 16)

	// Calculate first hash
	h1 := sha1.Sum([]byte(strings.ToUpper(username + ":" + password)))

	// Calculate second hash based on server core
	var h2 []byte
	hash := sha1.New()
	hash.Write(salt)
	hash.Write(h1[:])
	h2 = hash.Sum(nil)

	// Convert to big.Int (little-endian)
	h2Int := new(big.Int).SetBytes(reverseBytes(h2))

	// Calculate g^h2 mod N
	verifier := new(big.Int).Exp(g, h2Int, N)

	// Convert to byte array (little-endian)
	verifierBytes := reverseBytes(verifier.Bytes())

	// Pad to 32 bytes (maintaining little-endian)
	paddedVerifier := make([]byte, 32)
	copy(paddedVerifier, verifierBytes)

	return paddedVerifier
}

// reverseBytes reverses byte array
func reverseBytes(b []byte) []byte {
	reversed := make([]byte, len(b))
	for i := 0; i < len(b); i++ {
		reversed[i] = b[len(b)-1-i]
	}
	return reversed
}
