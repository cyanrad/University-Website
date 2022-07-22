package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const numeric = "1234567890"

// >> initializes the random number generator
func init() {
	rand.Seed(time.Now().UnixNano())
}

// >> generates a random 64 bit integer
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// >> generates random string of size n (contains numbers and letters)
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// >> generates a link like string
func RandomLink() string {
	return "https://www." + RandomString(7) + ".com"
}

// >> generate random string of size n (contains numbers)
func RandomNumberString(n int) string {
	var sb strings.Builder
	k := len(numeric)

	for i := 0; i < n; i++ {
		c := numeric[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// >> generate student id
func RandomID() string {
	return RandomNumberString(9)
}
