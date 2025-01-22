package token

import (
	"time"
	s "timeTrackerApp/src/server/Structures"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/google/uuid"
)

type Payload struct {
	jwt.RegisteredClaims
	ID        uuid.UUID `json:"id"`
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Admin     bool      `json:"admin"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(user *s.User, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		RegisteredClaims: jwt.RegisteredClaims{},
		ID:               tokenID,
		UserID:           user.UserID,
		Username:         user.Name,
		Admin:            user.IsAdmin(),
		IssuedAt:         time.Now(),
		ExpiredAt:        time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
