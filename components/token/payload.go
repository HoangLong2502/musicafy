package token

import (
	"time"

	usermodels "example.com/musicafy_be/modules/user/models"
	"github.com/google/uuid"
)

type Payload struct {
	usermodels.User
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	UserID   int       `json:"user_id"`
	IssuedAt time.Time `json:"issued_at"`
	ExpireAt time.Time `json:"expire_at"`
}

func NewPayload(data usermodels.User, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:       tokenID,
		Username: data.Username,
		UserID:   data.ID,
		IssuedAt: time.Now(),
		ExpireAt: time.Now().Add(duration),
	}

	return payload, nil
}
