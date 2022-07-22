package token

import "time"

// >> An interface for managing tokens
type Maker interface {
	// >> creates a new token for a specific  and duration
	CreateToken(ID string, duration time.Duration) (string, error)

	// >> Checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
