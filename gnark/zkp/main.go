package main

import (
	"fmt"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

// i really need comments to understand this
// It proves that the prover knows two factors of a public value

type Circuit struct {
	// priavte inputs bby default
	X frontend.Variable
	Y frontend.Variable
	// Define public inputs
	Z frontend.Variable `gnark:",public"`
}

// here declaring the circuit constraints --within Define func
func (circuit *Circuit) Define(api frontend.API) error {
	api.AssertIsEqual(api.Mul(circuit.X, circuit.Y), circuit.Z)
	return nil
}

func main() {
	// instantiate the circuit n compile
	var circuit Circuit
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		panic(err)
	}

	// Setup the ceremony, s etup proving and verification keys
	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		panic(err)
	}

	// Define our witness (both public and private inputs)
	assignment := Circuit{
		X: 3,
		Y: 5,
		Z: 15,
	}

	// Generate the proof
	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	if err != nil {
		panic(err)
	}
	proof, err := groth16.Prove(ccs, pk, witness)
	if err != nil {
		panic(err)
	}

	// Verify the proof
	publicWitness, err := witness.Public()
	if err != nil {
		panic(err)
	}
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		fmt.Println("Verification failed:", err)
	} else {
		fmt.Println("Verification successful!")
	}
}
