// Eli Bendersky [https://eli.thegreenplace.net]
// This code is in the public domain.
package rsa

import (
	"math/big"
)

// PublicKey is the public part of an RSA key pair.
type PublicKey struct {
	N *big.Int
	E *big.Int
}

// PrivateKey is the private part of an RSA key pair. According to RFC 2313 we
// could include the prime factors of N and other data here to make decryption
// faster, but N and D are sufficient for decrypting messages.
type PrivateKey struct {
	N *big.Int
	D *big.Int
}
