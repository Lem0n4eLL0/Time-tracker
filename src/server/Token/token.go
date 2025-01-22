package token

import (
	"errors"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
	JwtSekretKey    = "key"
)

func GetTokenMaker() Maker {
	return NewJWTMaker(JwtSekretKey)
}
