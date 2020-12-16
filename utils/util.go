package utils

import (
	"math/rand"
	"regexp"
	"time"
)

// RandString is a function return random string of n
func RandString(n int) string {
	var letter = []byte("qwertyuiopasdfghjklzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM")
	result := make([]byte, n)

	rand.Seed(time.Now().UnixNano())
	for i := range result {
		result[i] = letter[rand.Intn(len(letter))]
	}

	return string(result)
}

// EmailFormatCheck to check email format
func EmailFormatCheck(n string) bool {
	r1 := regexp.MustCompile(`[a-z0-9]+@[a-z0-9]+\.[a-z0-9]+`)
	if !r1.MatchString(n) {
		return false
	}

	return true
}
