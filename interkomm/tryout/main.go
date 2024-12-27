package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

const (
	// Define UDP and TCP port numbers
	udpPort = ":9000"
	tcpPort = ":9001"
)

var (
	// Global variable to hold the list of connected peers
	peers      = make(map[string]*Peer)
	peersMutex sync.RWMutex
)

// Peer structure represents a node in the network
type Peer struct {
	Address string
	Conn    net.Conn
}

// Start UDP for Node Discovery
func startUDP() {
	addr, err := net.ResolveUDPAddr("udp", udpPort)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received message from %s: %s\n", remoteAddr, string(buffer[:n]))
		// Respond with a simple message (hello)
		conn.WriteToUDP([]byte("Hello from UDP"), remoteAddr)
		// Store discovered peers
		discoveredPeer := remoteAddr.String()
		peersMutex.Lock()
		peers[discoveredPeer] = &Peer{Address: discoveredPeer}
		peersMutex.Unlock()
	}
}

// Start TCP for P2P Connections
func startTCP() {
	ln, err := net.Listen("tcp", tcpPort)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle new incoming connection in a separate goroutine
		go handleTCPConnection(conn)
	}
}

// Handle the TCP connection for P2P communication
func handleTCPConnection(conn net.Conn) {
	defer conn.Close()
	address := conn.RemoteAddr().String()
	fmt.Printf("New connection established with %s\n", address)

	// Keep reading messages from the peer
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Connection closed by %s\n", address)
			return
		}
		fmt.Printf("Message from %s: %s\n", address, string(buffer[:n]))
	}
}

// Broadcast a message to all connected peers over TCP
func broadcastMessage(message string) {
	peersMutex.RLock()
	defer peersMutex.RUnlock()
	for _, peer := range peers {
		if peer.Conn != nil {
			peer.Conn.Write([]byte(message))
		}
	}
}

// Join network: Discovery over UDP, then TCP communication
func joinNetwork() {
	// Start UDP for node discovery
	go startUDP()
	// Start TCP server to accept incoming P2P connections
	go startTCP()

	// Simulate joining by broadcasting a hello message
	time.Sleep(2 * time.Second)
	broadcastMessage("Hello, this is a new node!")

	select {} // Block forever
}

// CLI: Interact with the user to send messages
func startCLI() {
	for {
		var input string
		fmt.Print("Enter a message to broadcast or type 'exit' to quit: ")
		fmt.Scanln(&input)
		if input == "exit" {
			os.Exit(0)
		}
		broadcastMessage(input)
	}
}

func main() {
	// Start the node's networking layer
	go joinNetwork()

	// Start the CLI for user interaction
	startCLI()
}
