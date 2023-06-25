package main

import (
	"encoding/base64"
	"fmt"

	"github.com/hashicorp/vault/shamir"
)

const NumParts = 5
const Threshold = 3

// getShamirParts returns a slice of byte slices, each of which is a part of the
// secret. The number of parts is specified by the parts argument, and the
// threshold specifies the number of parts required to reconstruct the secret.
func getShamirParts(secret string, parts int, threshold int) ([][]byte, error) {
	fmt.Printf("Secret: %s\n", secret)
	return shamir.Split([]byte(secret), parts, threshold)
}

// getBase64ShamirParts returns a slice of base64 encoded strings, each of which
// is a part of the secret. The number of parts is specified by the parts
// argument, and the threshold specifies the number of parts required to
// reconstruct the secret.
func getBase64ShamirParts(secret string, parts int, threshold int) ([]string, error) {
	shamirParts, err := getShamirParts(secret, parts, threshold)
	if err != nil {
		return nil, err
	}

	encodedParts := make([]string, len(shamirParts))
	for i := range shamirParts {
		encodedParts[i] = base64.StdEncoding.EncodeToString(shamirParts[i])
	}
	return encodedParts, nil
}

// printShamirParts prints the parts of the secret to stdout.
func printShamirParts(parts []string) {
	for i := range parts {
		fmt.Printf("Part %d: %s\n", i, parts[i])
	}
}

// getBase64ShamirPartsFromUser prompts the user to enter the base64 encoded
// parts of the secret. The user is prompted for the number of parts specified
// by the Threshold constant. The parts are returned as a slice of byte slices.
func getBase64ShamirPartsFromUser() ([][]byte, error) {
	var parts [][]byte
	for i := 0; i < Threshold; i++ {
		// This should probably turn off echoing to STDOUT
		fmt.Printf("[%d] Enter any unique part of the original %d parts: ", i+1, NumParts)
		var part string
		fmt.Scanln(&part)
		base64Part, err := base64.StdEncoding.DecodeString(part)
		if err != nil {
			return nil, err
		}
		parts = append(parts, base64Part)
	}
	return parts, nil
}

// decodeShamirSecret returns the secret from the parts provided.
func decodeShamirSecret(parts [][]byte) (string, error) {
	secret, err := shamir.Combine(parts)
	if err != nil {
		return "", err
	}
	return string(secret), nil
}

func main() {
	// Generate the parts of the secret
	secret := "p455w0rdhunt3r2"
	parts, err := getBase64ShamirParts(secret, NumParts, Threshold)
	if err != nil {
		panic(err)
	}
	// Print the parts of the secret
	printShamirParts(parts)

	// Get the parts of the secret from the user, simulating a user entering the
	// parts of the secret at a later time.
	combinedParts, err := getBase64ShamirPartsFromUser()
	if err != nil {
		panic(err)
	}

	// Decode the secret from the parts provided
	secret, err = decodeShamirSecret(combinedParts)
	if err != nil {
		panic(err)
	}

	// Print the secret
	fmt.Printf("Retreived secret: %s\n", secret)
}
