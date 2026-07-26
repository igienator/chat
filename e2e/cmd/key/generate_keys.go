package key

import (
	"crypto/rand"
	"crypto/rsa"

	uuid "github.com/samborkent/uuidv8"
)

func GenerateUUID() uuid.UUID {
	return uuid.New()
}

func GenerateKey() (*rsa.PrivateKey, *rsa.PublicKey) {
	prvkey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	var pubkey *rsa.PublicKey
	pubkey = &prvkey.PublicKey

	return prvkey, pubkey
}
