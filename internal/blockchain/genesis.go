package blockchain

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"

	"github.com/Magicking/ether-swarm-services/models"
)

func NewAllocator(balance string) (string, *models.Allocator, error) {
	privateKey, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		return "", nil, fmt.Errorf("NewAllocator: %v", err)
	}
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	pkey := fmt.Sprintf("%064x", privateKey.D)
	allocator := &models.Allocator{
		Code: "",
		Storage: nil,
		Balance: &balance,
		PrivateKey: &pkey,
	}
	return address.Hex(), allocator, nil
}
