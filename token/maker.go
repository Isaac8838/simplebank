package token

import "time"

// maker is an interface for managin token
type Maker interface {
	// CreateToken creates a new token for specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken checks if the token is vaild or not
	VerifyToken(token string) (*Payload, error)
}
