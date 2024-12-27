// proof_verify.go

package main

import (
	"os"

	"github.com/PolyhedraZK/ExpanderCompilerCollection"
	"github.com/consensys/gnark/backend/groth16"
)

func main() {
	// Load the circuit and witness
	circuitData, _ := os.ReadFile("circuit.txt")
	witnessData, _ := os.ReadFile("witness.txt")
	circuit := ExpanderCompilerCollection.DeserializeCircuit(circuitData)
	witness := ExpanderCompilerCollection.DeserializeWitness(witnessData)

	// Setup (generate proving and verification keys)
	pk, vk, _ := groth16.Setup(circuit)

	// Generate a proof
	proof, _ := groth16.Prove(circuit, pk, witness)

	// Verify the proof
	err := groth16.Verify(proof, vk, witness)
	if err != nil {
		println("Proof verification failed:", err)
	} else {
		println("Proof verified successfully!")
	}

	// Optionally, you can serialize and save the proof and keys
	os.WriteFile("proof.txt", proof.Serialize(), 0644)
	os.WriteFile("proving_key.txt", pk.Serialize(), 0644)
	os.WriteFile("verification_key.txt", vk.Serialize(), 0644)
}
