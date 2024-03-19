package hash

import (
	"EduCRM/config"
	"crypto/sha1"
	"fmt"
)

func GeneratePasswordHash(password string) string {
	salt := config.Config().HashKey
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
