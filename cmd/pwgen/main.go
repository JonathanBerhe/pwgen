package main

import (
	"flag"
	"fmt"
	"github.com/jonathanberhe/pwgen/internal/generator"
	"log"
)

func main() {
	// Define command line flags
	length := flag.Int("length", 12, "Password length")
	passType := flag.String("type", "complex", "Password type: basic, alphanumeric, complex, memorable, pin, custom")
	count := flag.Int("count", 1, "Number of passwords to generate")
	customChars := flag.String("chars", "", "Custom character set for password generation")

	flag.Parse()

	// Create generator configuration
	config := &generator.Config{
		Length:      *length,
		Type:        *passType,
		CustomChars: *customChars,
	}

	// Initialize password generator
	gen, err := generator.New(config)
	if err != nil {
		log.Fatal(err)
	}

	// Generate passwords
	for i := 0; i < *count; i++ {
		pass, err := gen.Generate()
		if err != nil {
			log.Printf("Error generating password: %v", err)
			continue
		}
		fmt.Println(pass)
	}
}
