package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefjklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt генерирует рандомное число между min и max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString генерирует рандомную строку размера n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner генерирует рандомное имя для owner
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney генерирует количество денег
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomMoney генерирует код валюты
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "RUB"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
