package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random number between min and max
const alphabetNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const numbers = "0123456789"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabetNum)

	for i := 0; i < n; i++ {
		c := alphabetNum[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomNumber(n int) string {
	var sb strings.Builder
	k := len(numbers)

	for i := 0; i < n; i++ {
		c := numbers[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomUser() string {
	return RandomString(8)
}

func RandomPhoneNumber() string {
	return RandomNumber(11)
}

func RandomDateInFuture(n int) time.Time {
	return time.Now().AddDate(0, 0, n)
}

func GenerateUUID() string {
	return RandomString(8)
}
