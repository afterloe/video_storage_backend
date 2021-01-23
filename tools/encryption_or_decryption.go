package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// AES GCM 加密
func Encrypt(key, text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	var nonce = make([]byte, 12)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	cipherText := gcm.Seal(nil, nonce, []byte(text), nil)
	return hex.EncodeToString(cipherText), nil
}

// AES GCM 解密
func Decrypt(key, textHex string) (string, error) {
	cipherText, _ := hex.DecodeString(textHex)
	var nonce = make([]byte, 12)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	return string(plaintext), nil
}

// Sha256Hex 将文本转换为SHA256格式
func Sha256Hex(content string) string {
	s := sha256.New()
	s.Write([]byte(content))
	c := s.Sum(nil)
	data := hex.EncodeToString(c)
	return data
}

