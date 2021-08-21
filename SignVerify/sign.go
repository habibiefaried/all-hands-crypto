package main

import (
	"crypto"
	"crypto/sha256"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"fmt"
	"encoding/base64"
)

func writeFile(kalimat, sig string){
	d1 := []byte(kalimat)
    err := ioutil.WriteFile("msg.txt", d1, 0644)
    if (err != nil){
    	panic(err)
    }

    d2 := []byte(sig)
    err = ioutil.WriteFile("sig.txt", d2, 0644)
    if (err != nil){
    	panic(err)
    }
}

func main(){
    var kalimat string = "Cyber Academy Luar Biasa"

	priv, err := ioutil.ReadFile("private.key")
	if err != nil {
		panic(err)
	}

	privPem, _ := pem.Decode(priv)
	var privPemBytes []byte

	privPemBytes, err = x509.DecryptPEMBlock(privPem, []byte("1234"))
	if err != nil {
		panic(err)
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(privPemBytes); err != nil {
		fmt.Println(err)
		if parsedKey, err = x509.ParsePKCS8PrivateKey(privPemBytes); err != nil { 
			fmt.Println(err)
		}
	}

	var privateKey *rsa.PrivateKey
	var ok bool
	privateKey, ok = parsedKey.(*rsa.PrivateKey)
	if !ok {
		panic("Something wrong, cannot continue")
	}

	// Signing
	sum := sha256.Sum256([]byte(kalimat))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, sum[:])
	if (err != nil){
		panic(err)
	}
	sEnc := base64.StdEncoding.EncodeToString([]byte(signature))

	writeFile(kalimat, sEnc)
}