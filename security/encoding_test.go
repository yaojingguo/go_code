package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"testing"
)

var enc = base64.URLEncoding.WithPadding(base64.NoPadding)


// Verify https://jwt.io/#debugger-io. "secret base64 encoded" is unchecked.
func TestJwtGen(t *testing.T) {
	header := `{"alg":"HS256","typ":"JWT"}`
	payload := `{"sub":"1234567890","name":"John Doe","iat":1516239022}`
	base64Header := enc.EncodeToString([]byte(header))
	base64Payload := enc.EncodeToString([]byte(payload))

	data, sign := sign(base64Header, base64Payload, "abc")

	jwt := data + "." + sign
	expectedJwt := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.nZ86hUWPdG43W6HVSGFy6DJnDVOZhx8a73LhQ3gIxY8`

	if jwt != expectedJwt {
		t.Errorf("generated jwt %s is not the same as the expected jwt %s", jwt, expectedJwt)
	}
}

func sign(header, payload, key string) (data, sign string) {
	mac := hmac.New(sha256.New, []byte(key))
	data = header + "." + payload
	mac.Write([]byte(data))
	sign = enc.EncodeToString(mac.Sum(nil))
	return
}
// Play with 3.1.  Example JWT in RFC 7919.
func TestJwtExample(t *testing.T) {
	header := "{\"typ\":\"JWT\",\r\n \"alg\":\"HS256\"}"
	for _, r := range header {
		fmt.Printf("%v, ", r)
	}
	fmt.Println()
	fmt.Printf("encoded JOSE header: %s\n", enc.EncodeToString([]byte(header)))
	fmt.Printf("\r\n")

	payload := []byte{123, 34, 105, 115, 115, 34, 58, 34, 106, 111, 101, 34, 44, 13, 10,
		32, 34, 101, 120, 112, 34, 58, 49, 51, 48, 48, 56, 49, 57, 51, 56,
		48, 44, 13, 10, 32, 34, 104, 116, 116, 112, 58, 47, 47, 101, 120, 97,
		109, 112, 108, 101, 46, 99, 111, 109, 47, 105, 115, 95, 114, 111,
		111, 116, 34, 58, 116, 114, 117, 101, 125}
	fmt.Printf("encoded claims: %s\n", enc.EncodeToString(payload))
}

func Test(t *testing.T) {
	for _, r := range `\r` {
		fmt.Println(r)
	}
	fmt.Println()
	for _, r := range `\x0d` {
		fmt.Println(r)
	}
}
