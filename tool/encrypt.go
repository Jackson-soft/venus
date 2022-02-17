package tool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

func Sha3(txt string) string {
	h := sha3.New256()
	h.Write([]byte(txt))
	return hex.EncodeToString(h.Sum(nil))
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesEncrypt 加密函数
func AesEncrypt(plaintext, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)

	iv := key[:blockSize]
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plaintext))
	blockMode.CryptBlocks(crypted, plaintext)

	return hex.EncodeToString(crypted), nil
}

// AesDecrypt 解密函数
func AesDecrypt(ciphertext string, key []byte) (string, error) {
	decodeData, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()

	iv := key[:blockSize]
	blockMode := cipher.NewCBCDecrypter(block, iv)

	plaintext := make([]byte, len(decodeData))
	blockMode.CryptBlocks(plaintext, decodeData)
	plaintext = PKCS7UnPadding(plaintext)
	return string(plaintext), nil
}
