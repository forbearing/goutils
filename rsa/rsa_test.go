package rsa

import (
	"crypto/rsa"
	"fmt"
	"testing"
	//grsa "github.com/forbearing/goutils/pkg/rsa"
)

func TestRSA(t *testing.T) {
	var (
		privateKey *rsa.PrivateKey
		publicKey  *rsa.PublicKey
		err        error
	)

	privateKey, err = ReadRSAPKCS1PrivateKey("./private.pem.pkcs1")
	if err != nil {
		panic(err)
	}
	fmt.Println("private key(pkcs1):", privateKey)

	privateKey, err = ReadRSAPKCS8PrivateKey("./private.pem.pkcs8")
	if err != nil {
		panic(err)
	}
	fmt.Println("private key(pkcs8):", privateKey)

	privateKey, err = ReadRSAPrivateKey("./private.pem.pkcs1")
	if err != nil {
		panic(err)
	}
	fmt.Println("private key(pkcs1):", privateKey)

	privateKey, err = ReadRSAPrivateKey("./private.pem.pkcs8")
	if err != nil {
		panic(err)
	}
	fmt.Println("private key(pkcs8):", privateKey)

	publicKey, err = ReadRSAPublicKey("./public.pem")
	if err != nil {
		panic(err)
	}
	fmt.Println("public key(pkcs8):", publicKey)

	fmt.Println()
	data := "hello golang"
	encData, err := RSAEncrypt(data, "./public.pem")
	if err != nil {
		panic(err)
	}
	decData, err := RSADecrypt(encData, "./private.pem")
	if err != nil {
		panic(err)
	}
	signed, err := RSASign(data, "./private.pem")
	if err != nil {
		panic(err)
	}
	ok, err := VerifyRSASign(data, signed, "./public.pem")
	if err != nil {
		panic(err)
	}
	fmt.Println("plain data:", data)
	fmt.Println("encrypt data:", encData)
	fmt.Println("decrypt data:", decData)
	fmt.Println("signed data:", signed)
	fmt.Println("verified:", ok)
}
