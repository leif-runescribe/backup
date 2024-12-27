package main

import (
	"github.com/consensys/gnark/frontend"
)

type ProofCircuit struct {
	X frontend.Variable `gnark:"x,private"` // Private input
	G frontend.Variable `gnark:"g,public"`  // Public input
	Y frontend.Variable `gnark:"y,public"`  // Public input
	P frontend.Variable `gnark:"p,public"`  // Public input
}

func (circuit *ProofCircuit) Define(api frontend.API) error {
	// Compute g^x mod p
	result := api.Exp(circuit.G, circuit.X, circuit.P)

	// Ensure result equals Y
	api.AssertIsEqual(result, circuit.Y)
	return nil
}
