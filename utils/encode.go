package utils

import (
	"strings"
)

const (
	alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	base     = len(alphabet)
)

func EncodeBase62(num int) string {
	if num == 0 {
		return string(alphabet[0])
	}

	var encoded strings.Builder
	for num > 0 {
		remainder := num % base
		encoded.WriteByte(alphabet[remainder])
		num = num / base
	}

	return reverseString(encoded.String())
}

func reverseString(str string) string {
	var reversed strings.Builder
	strLen := len(str)

	for i := strLen - 1; i >= 0; i-- {
		reversed.WriteByte(str[i])
	}

	return reversed.String()
}
