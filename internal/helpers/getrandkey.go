package helpers

import (
	"math/rand"
	"time"
)

const charset = "0123456789ABCEFGHJKLMNPQRSTUVWXY"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GetRandKey(length int) string {
	return StringWithCharset(length, charset)
}
