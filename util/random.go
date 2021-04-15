package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	number   = "0123456789"
	email    = "%s@email.com"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	return Random(n, alphabet)
}

// RandomNumber generates a random number of length n
func RandomNumber(n int) string {
	return Random(n, number)
}

// Random generates a random of patter
func Random(n int, pattern string) string {
	var sb strings.Builder
	k := len(pattern)

	for i := 0; i < n; i++ {
		c := pattern[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomName generates 12 character a random name
func RandomName() string {
	return RandomString(12)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf(email, RandomString(6))
}

// RandomPhoneNumber generates a random phone number
func RandomPhoneNumber() string {
	return RandomNumber(12)
}

// RandomCode generates a random code
func RandomCode() string {
	return fmt.Sprintf("%s_%d", RandomNumber(4), time.Now().UnixNano())
}
