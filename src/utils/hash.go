package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

func Sha512Hashing(str string) string {
	sha512Hash := sha512.Sum512([]byte(str))
	return hex.EncodeToString(sha512Hash[:])
}
