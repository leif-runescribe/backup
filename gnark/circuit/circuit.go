// circuit/circuit.go
package circuit

import "github.com/consensys/gnark/frontend"

// MyCircuit defines a simple circuit
type MyCircuit struct {
	X frontend.Variable // Private input
	Y frontend.Variable // Private input
	Z frontend.Variable // Public input
}

// Define declares the circuit constraints
func (circuit *MyCircuit) Define(api frontend.API) error {
	// Constraint: X * Y = Z
	product := api.Mul(circuit.X, circuit.Y) // X * Y
	api.AssertIsEqual(product, circuit.Z)    // Ensure X * Y = Z
	return nil
}
