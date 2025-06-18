package zip_streamer

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func EncryptIt(value []byte, keyPhrase string) string {
	aesBlock, err := aes.NewCipher([]byte(keyPhrase))
	if err != nil {
	 fmt.Println(err)
	}
   
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
	 fmt.Println(err)
	}
	nonce := make([]byte, gcmInstance.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)
	cipheredText := gcmInstance.Seal(nonce, nonce, value, nil)
	enc := hex.EncodeToString(cipheredText)
	return enc
   }

func DecryptIt(ciphered string, keyPhrase string) string {
	fmt.Println(ciphered)
	decodedCipherText, err := hex.DecodeString(ciphered)
	aesBlock, err := aes.NewCipher([]byte(keyPhrase))
	if err != nil {
		fmt.Println("Error creating AES block: ", err)
	}
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		fmt.Println("Error creating GCM instance: ", err)
	}
	nonceSize := gcmInstance.NonceSize()
	nonce, cipheredText := decodedCipherText[:nonceSize], decodedCipherText[nonceSize:]
	fmt.Println(nonce)
	fmt.Println(cipheredText)
	decryptedText, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		fmt.Println("Error decrypting text: ", err)
	}
	return string(decryptedText)
}