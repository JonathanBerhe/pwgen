package password

import (
	"strings"
	"testing"
	"unicode"
)

// TestGenerateFromCharset verifies that generated passwords meet our requirements
func TestGenerateFromCharset(t *testing.T) {
	// table-driven tests
	tests := []struct {
		name        string
		length      int
		charset     string
		wantErr     bool
		checkResult func(t *testing.T, result string) // Custom validation function
	}{
		{
			name:    "valid basic password",
			length:  10,
			charset: Lowercase + Uppercase,
			checkResult: func(t *testing.T, result string) {
				if len(result) != 10 {
					t.Errorf("expected length 10, got %d", len(result))
				}
				// Check if only allowed characters are used
				for _, char := range result {
					if !unicode.IsLetter(char) {
						t.Errorf("unexpected character in result: %c", char)
					}
				}
			},
		},
		{
			name:    "numbers only",
			length:  6,
			charset: Numbers,
			checkResult: func(t *testing.T, result string) {
				if len(result) != 6 {
					t.Errorf("expected length 6, got %d", len(result))
				}
				for _, char := range result {
					if !unicode.IsNumber(char) {
						t.Errorf("expected only numbers, got: %c", char)
					}
				}
			},
		},
		{
			name:    "zero length",
			length:  0,
			charset: Lowercase,
			wantErr: true,
		},
		{
			name:    "empty charset",
			length:  8,
			charset: "",
			wantErr: true,
		},
		{
			name:    "special characters",
			length:  12,
			charset: Symbols,
			checkResult: func(t *testing.T, result string) {
				if len(result) != 12 {
					t.Errorf("expected length 12, got %d", len(result))
				}
				// Verify all characters are from our symbol set
				for _, char := range result {
					if !strings.ContainsRune(Symbols, char) {
						t.Errorf("unexpected character in result: %c", char)
					}
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GenerateFromCharset(tt.length, tt.charset)

			// Check error cases
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			// Run custom validation if provided
			if tt.checkResult != nil {
				tt.checkResult(t, result)
			}
		})
	}
}
