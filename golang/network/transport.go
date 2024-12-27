package network

import (
	"net"
	"time"
)

type netAddr string

type RPC struct {
	From    netAddr
	Payload []byte
}

type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	SendMessage(net.Addr, []byte) error
	Addr() net.Addr
	Start() error
	Stop() error
}

type Peer struct {
	netAddr
	LastSeen time.Time
}
