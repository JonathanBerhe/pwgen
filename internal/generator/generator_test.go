package generator

import (
	"strings"
	"testing"
	"unicode"

	"github.com/jonathanberhe/pwgen/pkg/password"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: &Config{
				Length: 12,
				Type:   "complex",
			},
			wantErr: false,
		},
		{
			name: "invalid length",
			config: &Config{
				Length: 0,
				Type:   "basic",
			},
			wantErr: true,
		},
		{
			name: "valid custom config",
			config: &Config{
				Length:      8,
				Type:        "custom",
				CustomChars: "ABC123",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerator_Generate(t *testing.T) {
	tests := []struct {
		name        string
		config      *Config
		wantErr     bool
		checkResult func(t *testing.T, result string)
	}{
		{
			name: "basic password",
			config: &Config{
				Length: 10,
				Type:   "basic",
			},
			checkResult: func(t *testing.T, result string) {
				if len(result) != 10 {
					t.Errorf("wrong length: got %d, want 10", len(result))
				}
				for _, char := range result {
					if !unicode.IsLetter(char) {
						t.Errorf("unexpected character in basic password: %c", char)
					}
				}
			},
		},
		{
			name: "alphanumeric password",
			config: &Config{
				Length: 12,
				Type:   "alphanumeric",
			},
			checkResult: func(t *testing.T, result string) {
				if len(result) != 12 {
					t.Errorf("wrong length: got %d, want 12", len(result))
				}
				validChars := password.Lowercase + password.Uppercase + password.Numbers
				for _, char := range result {
					if !strings.ContainsRune(validChars, char) {
						t.Errorf("invalid character in alphanumeric password: %c", char)
					}
				}
			},
		},
		{
			name: "complex password",
			config: &Config{
				Length: 16,
				Type:   "complex",
			},
			checkResult: func(t *testing.T, result string) {
				if len(result) != 16 {
					t.Errorf("wrong length: got %d, want 16", len(result))
				}
				validChars := password.Lowercase + password.Uppercase + password.Numbers + password.Symbols
				for _, char := range result {
					if !strings.ContainsRune(validChars, char) {
						t.Errorf("invalid character in complex password: %c", char)
					}
				}
			},
		},
		{
			name: "pin password",
			config: &Config{
				Length: 4,
				Type:   "pin",
			},
			checkResult: func(t *testing.T, result string) {
				if len(result) != 4 {
					t.Errorf("wrong length: got %d, want 4", len(result))
				}
				for _, char := range result {
					if !unicode.IsNumber(char) {
						t.Errorf("non-numeric character in PIN: %c", char)
					}
				}
			},
		},
		{
			name: "custom password",
			config: &Config{
				Length:      8,
				Type:        "custom",
				CustomChars: "ABC123",
			},
			checkResult: func(t *testing.T, result string) {
				if len(result) != 8 {
					t.Errorf("wrong length: got %d, want 8", len(result))
				}
				for _, char := range result {
					if !strings.ContainsRune("ABC123", char) {
						t.Errorf("invalid character in custom password: %c", char)
					}
				}
			},
		},
		{
			name: "custom password without charset",
			config: &Config{
				Length: 8,
				Type:   "custom",
			},
			wantErr: true,
		},
		{
			name: "invalid password type",
			config: &Config{
				Length: 8,
				Type:   "invalid",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := New(tt.config)
			if err != nil {
				t.Fatalf("failed to create generator: %v", err)
			}

			result, err := g.Generate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && tt.checkResult != nil {
				tt.checkResult(t, result)
			}
		})
	}
}

// TestPasswordUniqueness verifies that generated passwords are not predictable
func TestPasswordUniqueness(t *testing.T) {
	config := &Config{
		Length: 12,
		Type:   "complex",
	}

	g, err := New(config)
	if err != nil {
		t.Fatalf("failed to create generator: %v", err)
	}

	// Generate multiple passwords and ensure they're different
	passwords := make(map[string]bool)
	for i := 0; i < 100; i++ {
		pass, err := g.Generate()
		if err != nil {
			t.Fatalf("failed to generate password: %v", err)
		}

		if passwords[pass] {
			t.Errorf("duplicate password generated: %s", pass)
		}
		passwords[pass] = true
	}
}
