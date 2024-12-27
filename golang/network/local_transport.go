package network

import "sync"

type LocalTransport struct {
	addr      netAddr
	peers     map[netAddr]*LocalTransport
	lock      sync.RWMutex
	consumeCh chan RPC
}

func NewLocalTransport(addr netAddr) *LocalTransport {
	return &LocalTransport{
		addr:      addr,
		peers:     make(map[netAddr]*LocalTransport),
		consumeCh: make(chan RPC, 1024),
	}
}

func (t *LocalTransport) Connect() {
	return t.consumeCh
}
func (t *LocalTransport) Consume() {

}
func (t *LocalTransport) SendMessaget() {

}
func (t *LocalTransport) Addr() {

}
