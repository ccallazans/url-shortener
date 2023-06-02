package utils

import (
	"math/rand"
	"os"
)

func ReadSqlFile(path string) (string, error) {

	sql_query, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(sql_query), nil
}

func GenerateHash() string {
	var letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	hash := make([]byte, 5)
	for i := range hash {
		hash[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(hash)
}