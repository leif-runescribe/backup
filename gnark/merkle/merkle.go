package main

import (
	"fmt"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/std/hash/mimc"
)

// MerkleTreeCircuit defines our circuit struct
type MerkleTreeCircuit struct {
	// Public inputs
	Root frontend.Variable `gnark:",public"`

	// Private inputs (witness)
	Leaf   frontend.Variable
	Path   []frontend.Variable
	Helper []frontend.Variable
}

// Define specifies our circuit's constraints
func (circuit *MerkleTreeCircuit) Define(api frontend.API) error {
	mimc, _ := mimc.NewMiMC(api)

	// Start with the leaf
	current := circuit.Leaf

	// Traverse the Merkle path
	for i := 0; i < len(circuit.Path); i++ {
		left := api.Select(circuit.Helper[i], current, circuit.Path[i])
		right := api.Select(circuit.Helper[i], circuit.Path[i], current)

		mimc.Reset()
		mimc.Write(left, right)
		current = mimc.Sum()
	}

	// Assert that our computed root matches the public input
	api.AssertIsEqual(current, circuit.Root)

	return nil
}

func main() {
	// Instantiate the circuit
	var circuit MerkleTreeCircuit
	circuit.Path = make([]frontend.Variable, 4)
	circuit.Helper = make([]frontend.Variable, 4)

	// Compile the circuit
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		panic(err)
	}

	// Setup the ceremony
	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		panic(err)
	}

	// Define our witness (both public and private inputs)
	assignment := MerkleTreeCircuit{
		Root:   "16255948225469683072633460112291607001648355486865220912474029393968916097063",
		Leaf:   "1234567890",
		Path:   []frontend.Variable{"2345678901", "3456789012", "4567890123", "5678901234"},
		Helper: []frontend.Variable{1, 0, 1, 1},
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
