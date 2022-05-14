package main

import (
	"crypto/rsa"
	"crypto/x509"
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
