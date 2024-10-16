package main

import (
	"device_management/common/utils"
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(utils.GenerateUniqueKey())
	}
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
