package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/krile136/mineSweeper/store"
)

// 暗号化する
func Encrypt(text string) {
	plain := []byte(text)
	fmt.Println("original text: ", text)
	fmt.Println("plain text: ", plain)

	var keyString string = store.Data.Env.AesKey
	key, _ := hex.DecodeString(keyString)
	fmt.Println("keyString: ", keyString)
	fmt.Println("key: ", key)

	iv, _ := GenerateIV()

	block, _ := aes.NewCipher(key)

	padded := Pkcs7Pad(plain)
	encrypted := make([]byte, len(padded))
	cbcEncrypter := cipher.NewCBCEncrypter(block, iv)
	cbcEncrypter.CryptBlocks(encrypted, padded)
	// return iv, encrypted, nil
	fmt.Println("iv: ", iv)
	fmt.Println("encoded to string iv:", hex.EncodeToString(iv))
	fmt.Println("encrypted: ", encrypted)
	fmt.Println("encoded Encrypted:", base64.StdEncoding.EncodeToString(encrypted))
}

func GenerateIV() ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	return iv, nil
}

func Pkcs7Pad(data []byte) []byte {
	length := aes.BlockSize - (len(data) % aes.BlockSize)
	trailing := bytes.Repeat([]byte{byte(length)}, length)
	return append(data, trailing...)
}

// 復号化する
func Decrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decrypted := make([]byte, len(data))
	cbcDecrypter := cipher.NewCBCDecrypter(block, iv)
	cbcDecrypter.CryptBlocks(decrypted, data)
	return Pkcs7Unpad(decrypted), nil
}

func Pkcs7Unpad(data []byte) []byte {
	dataLength := len(data)
	padLength := int(data[dataLength-1])
	return data[:dataLength-padLength]
}
