package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

func HashPassword(password string) string {
	hasher := sha512.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}
