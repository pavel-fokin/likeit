package dbutil

import (
	"crypto/rand"
	"math/big"
)

const (
	maxLength = 6
	base62    = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// randomID generates a random base62 string of the maxLength.
func RandomID() (string, error) {
	var result string
	for i := 0; i < maxLength; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(62))
		if err != nil {
			return "", err
		}
		result += string(base62[num.Int64()])
	}
	return result, nil
}
