package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
)

/*
1.RSA 私钥存在 PKCS1 和 PKCS8 两种格式,通过 openssl 生成的私钥格式为 PKCS1, 公钥格式为 PKCS8.
2.生成私钥 (PKCS1 格式)
  openssl genrsa -out private.pem 1024
3.生成公钥 (PKCS8 格式)
  openssl rsa -in private.pem -pubout -out public.pem
4:私钥 PKCS1 转 PKCS8
  openssl pkcs8 -topk8 -inform pem -in private.pem -outform PEM -nocrypt
*/

// 数据签名验证
func VerifyRSASign(data, signed, publicKeyPath string) (bool, error) {
	// 反解 base64
	decodeData, err := base64.StdEncoding.DecodeString(signed)
	if err != nil {
		return false, err
	}

	// 读取公钥
	publicKey, err := ReadRSAPublicKey(publicKeyPath)
	if err != nil {
		return false, err
	}

	// 计算 sha1 散列值
	hash := sha256.New()
	hash.Write([]byte(data))
	bytes := hash.Sum(nil)
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, bytes, decodeData)
	return err == nil, err
}

// 数据签名
func RSASign(data, privateKeyPath string) (string, error) {
	// 读取私钥
	privateKey, err := ReadRSAPrivateKey(privateKeyPath)
	if err != nil {
		return "", err
	}

	// 计算 sha 散列值
	hash := sha256.New()
	hash.Write([]byte(data))
	sum := hash.Sum(nil)

	// 数字签名
	signed, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, sum)

	// 结果转成 base64
	return base64.StdEncoding.EncodeToString(signed), nil
}

// 解密(使用私钥解密)
func RSADecrypt(encData, privateKeyPath string) (string, error) {
	// 反解 base64
	decodeData, err := base64.StdEncoding.DecodeString(encData)
	if err != nil {
		return "", err
	}

	// 读取私钥
	privateKey, err := ReadRSAPrivateKey(privateKeyPath)
	if err != nil {
		return "", err
	}

	// 解密
	decryptData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decodeData)
	return string(decryptData), err

}

// 加密(使用公钥加密)
func RSAEncrypt(data, publicKeyPath string) (string, error) {
	// 获取公钥
	publicKey, err := ReadRSAPublicKey(publicKeyPath)
	if err != nil {
		return "", err
	}

	// 加密
	msg, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(data))
	if err != nil {
		return "", err
	}

	// 把加密结果转成 Base64
	return base64.StdEncoding.EncodeToString(msg), nil

}

// 读取私钥
func ReadRSAPrivateKey(name string) (*rsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	// pem 解码
	pemBlock, _ := pem.Decode(data)

	// x509 解码
	var privateKey interface{}
	switch pemBlock.Type {
	case "RSA PRIVATE KEY":
		privateKey, err = x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	case "PRIVATE KEY":
		privateKey, err = x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	}
	if err != nil {
		return nil, err
	}

	return privateKey.(*rsa.PrivateKey), nil
}

// 读取公钥
func ReadRSAPublicKey(name string) (*rsa.PublicKey, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	// 使用 pem 解码
	pemBlock, _ := pem.Decode(data)

	// x509 解码
	var publicKey interface{}
	switch pemBlock.Type {
	case "RSA PUBLIC KEY":
		publicKey, err = x509.ParsePKCS1PublicKey(pemBlock.Bytes)
	case "PUBLIC KEY":
		publicKey, err = x509.ParsePKIXPublicKey(pemBlock.Bytes)
	}

	if err != nil {
		return nil, err
	}
	return publicKey.(*rsa.PublicKey), nil
}

// 读取 PKCS1 格式私钥
func ReadRSAPKCS1PrivateKey(name string) (*rsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	// pem 解码
	pemBlock, _ := pem.Decode(data)
	// x509 解码
	return x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
}

// 读取 PKCS8 格式私钥
func ReadRSAPKCS8PrivateKey(name string) (*rsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	// pem 解码
	pemBlock, _ := pem.Decode(data)
	// x509 解码
	privateKey, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)

	return privateKey.(*rsa.PrivateKey), err
}
