package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"reflect"
	"testing"
)

// Load your secret key from a safe place and reuse it across multiple
// Seal/Open calls. (Obviously don't use this example key for anything
// real.) If you want to convert a passphrase to a key, use a suitable
// package like bcrypt or scrypt.
// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
var key, _ = hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
var plaintext = []byte("exampleplaintext")
var aad = []byte("aad text")
var nonce []byte

func encrypt(t *testing.T) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	text := aesgcm.Seal(nil, nonce, plaintext, aad)
	fmt.Printf("encrypted cipher text: %x\n", text)
	return text
}

func decrypt(t *testing.T, ciphertext []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	text, err := aesgcm.Open(nil, nonce, ciphertext, aad)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("decrypted plain text: %s\n", text)

	return text
}

func genNonce() []byte {
	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	bytes := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		panic(err.Error())
	}
	fmt.Printf("nonce: %x\n", bytes)
	return bytes
}

func TestEncryptionAndDecryption(t *testing.T) {
	// 32 bytes
	fmt.Printf("key size: %d\n", len(key))

	nonce = genNonce()

	ciphertext := encrypt(t)
	text := decrypt(t, ciphertext)

	if !reflect.DeepEqual(text, plaintext) {
		t.Fail()
	}
}
