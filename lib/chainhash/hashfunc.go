package chainhash

import "crypto/sha256"

// HashH calculates hash(b)
func HashH(b []byte) Hash {
	return Hash(sha256.Sum256(b))
}

// HashH calculates hash(hash(b))
func DoubleHashH(b []byte) Hash {
	first := sha256.Sum256(b)
	return Hash(sha256.Sum256(first[:]))
}
