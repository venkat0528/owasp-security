package crypto

import (
	"crypto/rand"
	"math/big"
)

func GetRandomNumber() (*big.Int, error) {
	rand, err := rand.Int(rand.Reader, big.NewInt(1984))
	return rand, err
}
