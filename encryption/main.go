package encryption

import (
	"math/rand"
	"time"
)

func RandomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i:= range b {
		b[i] = chars[seededRand.Intn(len(chars))]
	}
	return string(b)
}