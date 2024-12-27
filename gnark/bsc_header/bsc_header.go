package main

import (
	"log"

	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"golang.org/x/crypto/sha3"
)

// Implement a hash function using SHA3-256
func keccak256(data []byte) [32]byte {
	hash := sha3.New256()
	hash.Write(data)
	var result [32]byte
	copy(result[:], hash.Sum(nil))
	return result
}

type BlockHeaderCircuit struct {
	// Public Variables
	ParentHash       frontend.Variable
	UncleHash        frontend.Variable
	StateRoot        frontend.Variable
	TransactionsRoot frontend.Variable
	ReceiptsRoot     frontend.Variable
	GasLimit         frontend.Variable
	GasUsed          frontend.Variable
	Number           frontend.Variable
	Timestamp        frontend.Variable

	// Private Variables
	ComputedHash frontend.Variable
}

func (circuit *BlockHeaderCircuit) Define(api frontend.API) error {
	// Use a hash function to compute a hash value from the parent hash and uncle hash
	// This is a simplified example
	parentHashBytes := api.ToBytes(circuit.ParentHash)
	uncleHashBytes := api.ToBytes(circuit.UncleHash)
	combinedHash := append(parentHashBytes, uncleHashBytes...)
	computedHash := keccak256(combinedHash)

	// Here, we would need to implement the actual Keccak-256 in the circuit,
	// but for simplicity, assume we have a way to do so.

	// Example constraint: check if the computed hash matches the state root
	api.AssertIsEqual(circuit.ComputedHash, circuit.StateRoot)

	// Range checks
	api.AssertIsGreaterOrEqual(circuit.GasUsed, 0)
	api.AssertIsLessOrEqual(circuit.GasUsed, circuit.GasLimit)

	return nil
}

func main() {
	// Initialize the circuit
	circuit := &BlockHeaderCircuit{}

	// Compile the circuit
	r1cs, err := frontend.Compile(r1cs.NewBuilder(), circuit)
	if err != nil {
		log.Fatalf("Failed to compile circuit: %v", err)
	}

	// Generate the proving key and verification key
	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		log.Fatalf("Failed to setup: %v", err)
	}

	// Create a proof
	// Provide actual values in a real scenario
	witness := BlockHeaderCircuit{
		ParentHash:       123, // Example values
		UncleHash:        456,
		StateRoot:        579,
		TransactionsRoot: 101112,
		ReceiptsRoot:     131415,
		GasLimit:         10000,
		GasUsed:          5000,
		Number:           100,
		Timestamp:        1609459200,
	}

	proof, err := groth16.Prove(pk, witness)
	if err != nil {
		log.Fatalf("Failed to prove: %v", err)
	}

	// Verify the proof
	isValid := groth16.Verify(vk, proof, witness)
	if !isValid {
		log.Fatal("Proof verification failed")
	}

	log.Println("Proof verification succeeded")
}
