package utils

import "crypto/rand"

const alphaNum = "0123456789abcdefghijklmnopqrstuvwxyz"

func GenerateSecret(length uint) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = alphaNum[b%byte(len(alphaNum))]
	}

	return string(bytes), nil
}
