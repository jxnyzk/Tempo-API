package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

var key = []byte("@Syncro (faster than zrx (ratio lil bro)) ")

type AESCipher struct {
	blockSize int
	key       []byte
}

func NewAESCipher() *AESCipher {
	return &AESCipher{
		blockSize: aes.BlockSize,
		key:       hashKey(key),
	}
}

func hashKey(key []byte) []byte {
	hash := sha256.New()
	hash.Write(key)
	return hash.Sum(nil)
}

func Encrypt(raw string) (string, error) {
	a := NewAESCipher()
	plainText := []byte(raw)
	plainText = pad(a.blockSize, plainText)

	block, err := aes.NewCipher(a.key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, a.blockSize+len(plainText))
	iv := cipherText[:a.blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[a.blockSize:], plainText)

	hexEncoded := hex.EncodeToString(cipherText)
	return hexEncoded, nil
}


func Decrypt(enc string) (string, error) {
	a := NewAESCipher()
	decoded, err := hex.DecodeString(enc)
	if err != nil {
		return "", err
	}

	iv := decoded[:a.blockSize]
	cipherText := decoded[a.blockSize:]

	block, err := aes.NewCipher(a.key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText = unpad(cipherText)

	return string(cipherText), nil
}

func pad(blockSize int, src []byte) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func unpad(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:length-unpadding]
}
