package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandonInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandonString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandonOwner generates a random owner name
func RandonOwner() string {
	return RandonString(6)
}

// RandonMoney generates a random amount of money
func RandonMoney() int64 {
	return RandonInt(0, 1000)
}

// RandonCurrency generates a random currency code
func RandonCurrency() string {
	currencies := []string{"EUR", "USD", "AKZ", "RUP"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
