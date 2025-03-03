package chacha20poly1305

import (
	"encoding/hex"
	"errors"
	"strings"

	"golang.org/x/crypto/chacha20poly1305"
)

// const header = "$CHACHA20_POLY1305_VAULT;1.0"

func prepareOutput(data []byte, header string, lineLength int) []byte {
	hexData := hex.EncodeToString(data)
	lines := []string{header} // Initialize with the header

	for i := 0; i < len(hexData); i += lineLength {
		end := i + lineLength
		if end > len(hexData) {
			end = len(hexData) // No need for min if we slice this way
		}
		lines = append(lines, hexData[i:end])
	}

	finalString := strings.Join(lines, "\n")
	return []byte(finalString)
}

func prepareInput(input []byte) (header string, data []byte, err error) {
	s := string(input)
	s = strings.TrimSpace(s)

	parts := strings.SplitN(s, "\n", 2)

	if len(parts) == 1 {
		return parts[0], []byte(""), nil
	}

	if len(parts) != 2 {
		return "", nil, errors.New("invalid input format: expected header and data")
	}

	header = parts[0]
	dataPart := parts[1]

	dataPart = strings.ReplaceAll(dataPart, "\n", "")
	dataPart = strings.TrimSpace(dataPart)

	if len(dataPart) == 0 {
		return header, []byte(dataPart), nil
	}

	dataBytes, err := hex.DecodeString(dataPart)
	if err != nil {
		return "", nil, err
	}

	return header, dataBytes, nil
}

// Decrypt decrypts the provided ciphertext using the ChaCha20-Poly1305 algorithm.
//
// Parameters:
//
//	ciphertext: The encrypted data to be decrypted.
//	key: The secret key used for decryption.
//	nonce: The nonce (number used once) used for decryption.
//
// Returns:
//
//	plaintext: The decrypted data.
//	error: Any error that occurred during decryption.
func Decrypt(ciphertext []byte, key []byte, nonce []byte) ([]byte, error) {
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, err
	}
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// Encrypt encrypts the provided plaintext using the ChaCha20-Poly1305 algorithm.
//
// Parameters:
//
//	plaintext: The data to be encrypted.
//	key: The secret key used for encryption.
//	nonce: The nonce (number used once) used for encryption.
//
// Returns:
//
//	ciphertext: The encrypted data.
//	error: Any error that occurred during encryption.
func Encrypt(plaintext []byte, key []byte, nonce []byte) ([]byte, error) {
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, err
	}
	ciphertext := aead.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nil
}
