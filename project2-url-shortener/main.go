package main

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateCode returns a random n-character code.
func GenerateCode(n int) string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, n)
	for i := range code {
		code[i] = charset[rng.Intn(len(charset))]
	}
	return string(code)
}

func main() {
	code := GenerateCode(6)
	fmt.Println("Generated code:", code)
}