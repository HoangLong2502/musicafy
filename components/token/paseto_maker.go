package token

import (
	"fmt"
	"time"

	usermodels "example.com/musicafy_be/modules/user/models"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetorMake struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func (maker *PasetorMake) CreateToken(data usermodels.User) (string, *Payload, error) {
	payload, err := NewPayload(data, 24*time.Hour)
	if err != nil {
		return "", nil, err
	}
	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, payload, err
}

func (maker *PasetorMake) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrEncodingToken
	}

	return payload, nil
}

func NewPasetoMaker(symmetricKey string) (TokenMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly #{chacha20poly1305.KeySize} characters")
	}

	maker := &PasetorMake{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}
