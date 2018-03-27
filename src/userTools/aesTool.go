package userTools

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"bytes"
)

/*
*@method AES加密
*@param  origData 需要加密的数据
*@param  key 加密的KEY  必须8的倍数
*@return 得到的是加密后的数据
*/
func AesEncry(origData string,key string) string  {
	if len(origData) == 0 {
		return ""
	}
	 block, err :=  aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}
	origDataByte := []byte(origData)
	blockSize := block.BlockSize()
	origDataByte = ZeroPadding(origDataByte,blockSize)
	keyByte := []byte(key)
	blockModel := cipher.NewCBCEncrypter(block,keyByte[:blockSize])
	crypted := make([]byte,len(origDataByte))
	blockModel.CryptBlocks(crypted,origDataByte)
	if len(crypted) != 0 {
		return base64.StdEncoding.EncodeToString(crypted)
	}
	return ""
}

/*
*@method AES解密
*@param decryString 需要解密的数据
*@param	key 解密需要的AESkey
*@return 得到解密后的字符串
*/
func AesDecry(decryString string, key string) string {
	if len(decryString) == 0 {
		return ""
	}
	block, err :=  aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}
	decryStrByte,_:=base64.StdEncoding.DecodeString(decryString)
	blockSize := block.BlockSize()
	keyByte := []byte(key)
	blockModel := cipher.NewCBCDecrypter(block,keyByte[:blockSize])
	decryByte := make([]byte,len(decryStrByte))
	blockModel.CryptBlocks(decryByte,decryStrByte)
	decryByte = ZeroUnPadding(decryByte,blockSize)
	if len(decryByte) !=0 {
		return string(decryByte)
	}

	return ""
}




//cbc 后面填充0x0
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}


func ZeroUnPadding(origData []byte,blockSize int) []byte {
	length := len(origData)
	var a int = 1
	var unpadding  = 0
	for a = 0; a < blockSize; a++ {
		unpadding = int(origData[length-a-1])
		if unpadding !=0 {
			break;
		}
	}
	return origData[:(length - a)]
}