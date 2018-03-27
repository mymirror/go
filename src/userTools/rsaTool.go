package userTools

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"
	"io/ioutil"
	"encoding/base64"
	"crypto/sha1"
	"crypto"
)

//参数 bits 私钥的位数 一般是 512 1024 2048
func GenRsaKey(bits int) error  {
	privateKey, err := rsa.GenerateKey(rand.Reader,bits)

	if err != nil {
		 return err
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)

	block := &pem.Block{
		Type: "私钥",
		Bytes:derStream,
	}

	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}

	err  = pem.Encode(file,block)

	if err != nil{
		return err
	}

	//生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "公钥",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil

}

//获取公钥
func GetPublickRsaKey() (error,string)  {
	publickKeyByte , err := ioutil.ReadFile("public.pem")

	if err != nil {
		return err,""
	}

	publickKeyString := string(publickKeyByte)
	return nil,publickKeyString;

}

func GetPrivateRsaKey() (error, string)  {
	privateKeyByte , err := ioutil.ReadFile("private.pem")

	if err != nil {
		return err,""
	}

	privateKeyString := string(privateKeyByte)
	return nil,privateKeyString
}

/*
@method  数据进行RSA加密
@param  origString 需要进行加密的原数据 字符串类型
@param  publickKey 加密所需的RSA公钥 字符串类型
@return 返回加密后的数据 字符串类型
*/
func RsaPublicEncry(origString string ,publicKey string) string  {
	if len(origString) == 0 || len(publicKey) == 0 {
		return ""
	}

	data := []byte(publicKey);
	block, _ := pem.Decode(data)
	if block == nil {
		return ""
	}

	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		return ""
	}

	pub := publicInterface.(*rsa.PublicKey)
	encryByte, err := rsa.EncryptPKCS1v15(rand.Reader,pub,[]byte(origString))

	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(encryByte)


}

/*
@method 私钥解密
@param  decryString 需要解密的数据
@param  privateKey  私钥
@return 得到解密后的原数据
*/
func RsaPrivateDecry(decryString string,privateKey string) string  {
	if len(decryString) == 0 || len(privateKey) == 0 {
		return ""
	}
	decrtByte, err :=  base64.StdEncoding.DecodeString(decryString)
	if err != nil {
		return ""
	}
	data := []byte(privateKey)
	block,_ := pem.Decode(data)
	if block == nil {
		return ""
	}

	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		return ""
	}

	decryByte, err := rsa.DecryptPKCS1v15(rand.Reader,private,decrtByte)
	if err!= nil {
		return ""
	}

	return string(decryByte)

}


/*
*@method 私钥签名
*@param  orignData 签名之前的数据
*@paam   privateKey 签名的私钥
*@return bool 是否签名成功 1 成功 0 失败
*@return string 签名之后的数据 base64字符串
*/
func SignData(orignData string, privateKey string) (bool,string)  {
	if len(orignData) == 0 || len(privateKey) == 0 {
		return false,""
	}
	data := []byte(privateKey)
	block,_ := pem.Decode(data)
	if block == nil {
		return false,""
	}

	pri,err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return false,""
	}

	strData := []byte(orignData)

	h := sha1.New()
	h.Write(strData)
	hashed := h.Sum(nil)

	signture, err := rsa.SignPKCS1v15(rand.Reader,pri,crypto.SHA1,hashed)
	if err != nil {
		return false,""
	}
	return true,base64.StdEncoding.EncodeToString(signture);

}

/*
*method 公钥验证签名
*@param  signedData 签名之后的数据
*@param  publicKey 公钥
*@return bool 是否验证成功 1 success 0 fail
*/
func VerifySign(signedData string,origData string,publicKey string) bool {
	if len(signedData) == 0 || len(origData) == 0 || len(publicKey) == 0 {
		return false
	}

	data := []byte(publicKey)
	block,_ := pem.Decode(data)
	if block == nil {
		return false
	}

	pubInterface,err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false
	}

	pub:= pubInterface.(*rsa.PublicKey)

	signData,err := base64.StdEncoding.DecodeString(signedData)


	if err != nil {
		return false
	}


	hashed := sha1.Sum([]byte(origData))

	err = rsa.VerifyPKCS1v15(pub,crypto.SHA1,hashed[:],signData)

	if err!= nil {
		return false
	}
	return true
}
