package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

// >> Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// >> A PASETO session token maker
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// >> Creates a new PASETO token maker
func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf(
			"invalid key size: must be exactly %d chars", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

// >> creates new token
func (maker *PasetoMaker) CreateToken(ID string, duration time.Duration) (string, error) {
	payload, err := NewPayload(ID, duration)
	if err != nil {
		return "", err
	}

	return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
}

// >> decrypts and validates the recieved token
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	// >> decryting the token
	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	// >> checking if expired
	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
