package token

import (
	"time"
	s "timeTrackerApp/src/server/Structures"
)

type Maker interface {
	CreateToken(username *s.User, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
