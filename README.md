# Password Generator CLI 🔐

A secure and flexible command-line tool for generating various types of passwords. This tool supports multiple password types, custom character sets, and configurable password lengths while maintaining cryptographic security.
Features

- Multiple password types (basic, alphanumeric, complex, PIN, custom)
- Configurable password length
- Secure random generation using crypto/rand
- Cross-platform support (Windows, Linux, macOS)
- Custom character set support
- Multiple password generation in a single command

## Usage

The password generator supports various command-line flags to customize password generation:

```bash
pwgen [flags]

Flags:
  --length      Password length (default: 12, max 256)
  --type        Password type: basic, alphanumeric, complex, memorable, pin, custom (default: complex)
  --count       Number of passwords to generate (default: 1)
  --chars       Custom character set for password generation (required for custom type)
````

## Examples

Generate a default complex password (12 characters):

```bash
pwgen
# Output: K8#mP9$vL2@n
```

Generate a longer complex password:

```bash
pwgen --length 16
# Output: Rx9#mK$vL2@nP5&j
```

Generate multiple passwords:

```bash
pwgen --count 3
# Output:
# K8#mP9$vL2@n
# Rx9#mK$vL2@n
# P5&jN7#kM4$w
````

Generate a PIN:

```bash
pwgen --type pin --length 6
# Output: 847591
```

Generate a basic password (letters only):

```bash
pwgen --type basic --length 10
# Output: KmPvLnRxMk
```

Generate an alphanumeric password:

```bash
pwgen --type alphanumeric --length 8
# Output: Kx9mP2vL
```

Generate a password with a custom character set:

```bash
pwgen --type custom --chars "ABC123!@" --length 8
# Output: A1B2C3@!
```
