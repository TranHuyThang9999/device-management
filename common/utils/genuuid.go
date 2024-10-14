package utils

import (
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

var (
	mu sync.Mutex
)

func GenerateUniqueKey() int64 {
	mu.Lock()
	defer mu.Unlock()
	var length = 8
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	key := int64(0)
	for i := 0; i < length; i++ {
		key = key*10 + int64(seededRand.Intn(9)) + 1
	}
	return key
}

func GenerateTimestamp() int64 {
	timeNow := time.Now()
	return timeNow.Unix()
}
func ConvertTimestampToDateTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	formattedDateTime := t.Format("2006-01-02 15:04:05")
	return formattedDateTime
}
func GenerateTimestampExpiredAt(expiredAt int) *int {
	timeNow := time.Now()
	expirationTime := timeNow.Add(time.Duration(expiredAt) * time.Minute)
	timestamp := int(expirationTime.Unix())
	return &timestamp
}
func GenerateNameFile() string {
	return uuid.NewString()
}

func GenPassWord() string {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	all := lower + upper + digits

	passLength := 12
	password := make([]byte, passLength)

	for i := range password {
		index := rand.Intn(len(all))
		password[i] = all[index]
	}

	return string(password)
}
