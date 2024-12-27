package main

import (
	"math/big"

	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

func main() {
	// Instantiate the circuit
	var circuit ProofCircuit

	// Compile the circuit into R1CS
	r1cs, err := frontend.Compile(r1cs.NewBuilder, &circuit)
	if err != nil {
		panic(err)
	}

	// Generate proving and verifying keys
	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		panic(err)
	}

	// Assign values to public and private variables
	assignment := ProofCircuit{
		X: big.NewInt(3),  // Private input (secret exponent)
		G: big.NewInt(2),  // Public input (generator)
		Y: big.NewInt(8),  // Public input (g^x mod p = 2^3 mod 13 = 8)
		P: big.NewInt(13), // Public input (prime modulus)
	}

	// Generate a proof
	proof, err := groth16.Prove(r1cs, pk, &assignment)
	if err != nil {
		panic(err)
	}

	// Verify the proof
	err = groth16.Verify(proof, vk, &assignment)
	if err != nil {
		panic(err)
	}

	println("Proof verified successfully!")
}
