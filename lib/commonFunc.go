package lib

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"github.com/sirupsen/logrus"
)

func FatalHandler(err error, info string) {
	logrus.Fatal(info, err)
}

func ChangeHouse(house string) string {
	key := "otherBuilding"
	switch house {
	case `服院`:
		key = "costumeBuilding"
	case `艺院`:
		key = "artBuilding"
	case `综合楼`:
		key = "mainBuilding"
	}
	return key
}

func GetHouse(class string) string {
	//综合楼 服装 艺院 其他
	house := class[:3]
	switch house {
	case `服`:
		house = `服院`
	case `艺`:
		house = `艺院`
	case `综`:
		house = `综合楼`
	default:
		house = `其他`
	}
	return house
}

func AesEnc(ciphertext1 string) string {
	ciphertext := []byte(ciphertext1)
	key, err := hex.DecodeString(Config.PwdKey)
	if err != nil {
		panic(err)
	}
	iv, err := hex.DecodeString(Config.PwdIv)
	if err != nil {
		panic(err)
	}
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	plaintext := append(ciphertext, padText...)
	ciphertext = make([]byte, len(plaintext))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)
	return hex.EncodeToString(ciphertext)
}

func AesDec(plantText string) string {
	hexData, _ := hex.DecodeString(plantText)
	key, err := hex.DecodeString(Config.PwdKey)
	if err != nil {
		panic(err)
	}
	iv, err := hex.DecodeString(Config.PwdIv)
	if err != nil {
		panic(err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(hexData, hexData)
	length := len(hexData)
	unPadding := int(hexData[length-1])
	return string(hexData[:(length - unPadding)])
}
