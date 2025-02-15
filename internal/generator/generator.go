package generator

import (
	"fmt"

	"github.com/jonathanberhe/pwgen/pkg/password"
)

const (
	basicPw        = "basic"
	alphanumericPw = "alphanumeric"
	complexPw      = "complex"
	pinPw          = "pin"
	customPw       = "custom"
)

const maxPwLenght = 256

type Config struct {
	Length      int
	Type        string
	CustomChars string
}

type Generator struct {
	config *Config
}

func New(config *Config) (*Generator, error) {
	if config.Length <= 0 {
		return nil, fmt.Errorf("invalid length: must be positive")
	}
	if config.Length > maxPwLenght {
		return nil, fmt.Errorf("invalid length: must be less then equal %v", maxPwLenght)
	}
	return &Generator{config: config}, nil
}

func (g *Generator) Generate() (string, error) {
	var charset string

	switch g.config.Type {
	case basicPw:
		charset = password.Lowercase + password.Uppercase
	case alphanumericPw:
		charset = password.Lowercase + password.Uppercase + password.Numbers
	case complexPw:
		charset = password.Lowercase + password.Uppercase + password.Numbers + password.Symbols
	case pinPw:
		charset = password.Numbers
	case customPw:
		if g.config.CustomChars == "" {
			return "", fmt.Errorf("custom character set is required for custom type")
		}
		charset = g.config.CustomChars
	default:
		return "", fmt.Errorf("unsupported password type: %s", g.config.Type)
	}

	return password.GenerateFromCharset(g.config.Length, charset)
}
