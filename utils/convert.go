package utils

import (
	"math/rand"
	"time"
)

func RandomName() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWYZ"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := make([]byte, 12)

	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}

	return string(result)
}
