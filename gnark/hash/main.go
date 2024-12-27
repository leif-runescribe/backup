package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Compute Keccak-256 hash of combined parent and uncle hashes
func keccak256(data []byte) [32]byte {
	hash := sha256.New224()
	hash.Write(data)
	var result [32]byte
	copy(result[:], hash.Sum(nil))
	return result
}

func main() {
	parentHash := []byte("parent hash example")
	uncleHash := []byte("uncle hash example")

	combinedHash := append(parentHash, uncleHash...)
	computedHash := keccak256(combinedHash)
	fmt.Println("Computed Hash (Hex):", hex.EncodeToString(computedHash[:]))
}
