package main

import (
	"crypto/rsa"
	"fmt"

	grsa "github.com/forbearing/goutils/pkg/rsa"
)

func main() {
	var (
		privateKey *rsa.PrivateKey
		publicKey  *rsa.PublicKey
		err        error
	)

	privateKey, err = grsa.ReadRSAPKCS1PrivateKey("./private.pem.pkcs1")
	if err != nil {
		panic(err)
	}
	fmt.Println("private key(pkcs1):", privateKey)

	privateKey, err = grsa.ReadRSAPKCS8PrivateKey("./private.pem.pkcs8")
	if err != nil {
		panic(err)
	}
	fmt.Println("private key(pkcs8):", privateKey)

	privateKey, err = grsa.ReadRSAPrivateKey("./private.pem.pkcs1")
	if err != nil {
		panic(err)
	}
	fmt.Println("private key(pkcs1):", privateKey)

	privateKey, err = grsa.ReadRSAPrivateKey("./private.pem.pkcs8")
	if err != nil {
		panic(err)
	}
	fmt.Println("private key(pkcs8):", privateKey)

	publicKey, err = grsa.ReadRSAPublicKey("./public.pem")
	if err != nil {
		panic(err)
	}
	fmt.Println("public key(pkcs8):", publicKey)
}
