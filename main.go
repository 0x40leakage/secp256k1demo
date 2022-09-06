package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func main() {
	secp256k1Curve := secp256k1.S256()
	const privStr = "894ca3c5afb5ae1552c55e14c3b73ca2ac8a092408b0e92d31af1b6538035313"
	priv := &ecdsa.PrivateKey{}
	priv.D, _ = new(big.Int).SetString(privStr, 16)
	priv.PublicKey.Curve = secp256k1Curve
	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(priv.D.Bytes())

	pub := elliptic.MarshalCompressed(secp256k1Curve, priv.PublicKey.X, priv.PublicKey.Y)
	fmt.Printf("%x\n", pub) // 0212b55b9431515c7185355f15b48c5e1a1bbfa31af61429fa2bb8709de722f420
}
