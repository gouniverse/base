package chacha20poly1305

import (
	"golang.org/x/crypto/chacha20poly1305"
)

// const header = "$CHACHA20_POLY1305_VAULT;1.0"

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
