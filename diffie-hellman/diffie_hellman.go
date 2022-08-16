package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

var bigInt big.Int

func PrivateKey(p *big.Int) *big.Int {
	priv, _ := rand.Int(rand.Reader, bigInt.Sub(p, big.NewInt(2)))
	priv = bigInt.Add(priv, big.NewInt(2))
	return priv
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return bigInt.Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	a := PrivateKey(p)
	A := PublicKey(a, p, g)
	return a, A
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return bigInt.Exp(public2, private1, p)
}
