package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// >> token data
type Payload struct {
	ID        uuid.UUID `json:"id"`
	StudentID string    `json:"student_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// >> creates a new token payload
func NewPayload(ID string, duration time.Duration) (*Payload, error) {
	// >> creating new token
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	// >> creating new payload
	payload := &Payload{
		ID:        tokenID,
		StudentID: ID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}

// >> verifies that the token is not expired
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return errors.New("token has expired")
	}
	return nil
}
