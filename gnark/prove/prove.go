// prove/prove.go
package prove

import (
	"log"

	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/yourusername/gnark-circuit/circuit"
)

// CompileAndProve compiles the circuit and generates/verifies a proof
func CompileAndProve() {
	// Initialize the circuit
	var myCircuit circuit.MyCircuit

	// Compile the circuit into an R1CS (Rank-1 Constraint System)
	r1cs, err := frontend.Compile(frontend.NewCS, &myCircuit)
	if err != nil {
		log.Fatalf("failed to compile the circuit: %v", err)
	}

	// Generate proving and verifying keys
	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		log.Fatalf("failed to setup the keys: %v", err)
	}

	// Define the witness with actual values
	witness := circuit.MyCircuit{
		X: 3,  // Prover's private input
		Y: 4,  // Prover's private input
		Z: 12, // Public input
	}

	// Generate a proof
	proof, err := groth16.Prove(r1cs, pk, &witness)
	if err != nil {
		log.Fatalf("failed to generate proof: %v", err)
	}

	// Verify the proof
	publicWitness := circuit.MyCircuit{
		Z: 12, // Public input known to the verifier
	}

	err = groth16.Verify(proof, vk, &publicWitness)
	if err != nil {
		log.Fatalf("proof verification failed: %v", err)
	} else {
		log.Println("proof verified successfully")
	}
}
