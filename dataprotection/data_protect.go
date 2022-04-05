package dataprotection

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/nacl/secretbox"
)

func DataProtection() {

	secretKeyBytes, err := hex.DecodeString("779797979797809808098080807078588585788868685858")
	if err != nil {
		panic(err)
	}

	var secretKey [32]byte
	copy(secretKey[:], secretKeyBytes)

	var nonce [24]byte
	if _, err := rand.Read(nonce[:]); err != nil {
		panic(err)
	}

	// This encrypts "welome to finmont" and appends the result to the nonce.
	encrypted := secretbox.Seal(nonce[:], []byte("welome to finmont"), &nonce, &secretKey)

	// decrypts nonce
	var decryptNonce [24]byte
	copy(decryptNonce[:], encrypted[:24])
	decrypted, ok := secretbox.Open([]byte{}, encrypted[24:], &decryptNonce, &secretKey)
	if !ok {
		panic("decryption error")
	}

	fmt.Println(string(decrypted))
}
