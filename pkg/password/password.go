package password

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const (
	Lowercase = "abcdefghijklmnopqrstuvwxyz"
	Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers   = "0123456789"
	Symbols   = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

// Generator defines the interface for password generation
type Generator interface {
	Generate() (string, error)
}

// generateFromCharset creates a random string using the provided character set
func GenerateFromCharset(length int, charset string) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be positive")
	}
	if len(charset) == 0 {
		return "", fmt.Errorf("charset cannot be empty")
	}

	var builder strings.Builder
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		builder.WriteByte(charset[randomIndex.Int64()])
	}

	return builder.String(), nil
}
