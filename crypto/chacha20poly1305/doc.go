// Package chacha20poly1305 provides ChaCha20-Poly1305 encryption and decryption functions.
//
// ChaCha20-Poly1305 is a widely used and respected authenticated encryption algorithm that provides both confidentiality and integrity.
// It is designed to be fast and efficient, making it suitable for high-performance applications.
//
// The package provides two main functions: Encrypt and Decrypt.
// The Encrypt function encrypts the provided plaintext using the ChaCha20-Poly1305 algorithm.
// The Decrypt function decrypts the provided ciphertext using the ChaCha20-Poly1305 algorithm.
//
// Both functions take three parameters: the data to be encrypted or decrypted, the secret key, and the nonce.
// The secret key and nonce must be kept secret to ensure the security of the encryption process.
//
// The package also provides an example usage of the Encrypt and Decrypt functions.
//
// Pros of using ChaCha20-Poly1305:
//   - High security
//   - Fast performance
//   - Stream cipher
//   - Nonce-based
//   - Authenticated encryption
//   - Wide platform support
//   - Open-source implementation
//
// Cons of using ChaCha20-Poly1305:
//   - Key management
//   - Nonce management
//   - Limited key size
//   - Limited nonce size
//   - Side-channel attacks
//   - Quantum computer attacks
//   - Limited cryptographic agility
//
// Example usage:
//
//	package main
//
//	import (
//	    "fmt"
//	    "log"
//	)
//
//	func main() {
//	    // Define the key and nonce
//	    key := []byte("your_secret_key_here")
//	    nonce := []byte("your_nonce_here")
//
//	    // Define the plaintext to be encrypted
//	    plaintext := []byte("Hello, World!")
//
//	    // Encrypt the plaintext
//	    ciphertext, err := Encrypt(plaintext, key, nonce)
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//
//	    // Decrypt the ciphertext
//	    decrypted, err := Decrypt(ciphertext, key, nonce)
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//
//	    // Print the decrypted text
//	    fmt.Println(string(decrypted))
//	}
package chacha20poly1305
