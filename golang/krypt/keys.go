package crypto

import "crypto/ed25519"

const (
	privKeyLen = 64
	pubKeyLen  = 32
)

type privKey struct {
	key ed25519.PrivateKey
}

func (p *privKey) Bytes() []byte {
	return p.key
}

func (p *privKey) Sign(msg []byte) []byte {
	return ed25519.Sign(p.key, msg)
}


func (p *privKey) Public() *publicKey {
	x := make([]byte, pubKeyLen)
	b := make([]byte, pubKeyLen)
	copy(b, p.key[32:])
	return &publicKey{
		key: b
	}
}

type publicKey struct {
	key ed25519.PublicKey
}