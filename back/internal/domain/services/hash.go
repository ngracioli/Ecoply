package services

import (
	"crypto/sha256"
	"ecoply/internal/mlog"
	"fmt"
)

func Hash256String(input string) string {
	hash := sha256.New()
	_, err := hash.Write([]byte(input))
	if err != nil {
		mlog.Log("Failed to generate hash: " + err.Error())
		panic(err)
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}
