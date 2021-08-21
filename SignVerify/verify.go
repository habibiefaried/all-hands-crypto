package main

import (
	"io/ioutil"
	"fmt"
	"crypto"
	"crypto/sha256"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"	
	"encoding/base64"
)

func readFile(loc string) string {
	b, err := ioutil.ReadFile(loc) 
    if err != nil {
        panic(err)
    }
    return string(b)
}

func main(){
	kalimat := readFile("msg.txt")
	sign := readFile("sig.txt")

	pub, err := ioutil.ReadFile("public.key")
	if err != nil {
		panic(err)
	}

	pubPem, _ := pem.Decode(pub)
	if pubPem == nil {
		panic("RSA not in PEM format")
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKIXPublicKey(pubPem.Bytes); err != nil {
		panic(err)
	}

	pubKey, ok := parsedKey.(*rsa.PublicKey) 
	if !ok {
		panic("cannot parse key..")
	}

	// Verify
	sum := sha256.Sum256([]byte(kalimat))
	decodedSig, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		panic(err)
	}
	
	hasilError := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, sum[:], decodedSig)
	if (hasilError != nil){
		fmt.Println("Pesan tidak terverifikasi")
	} else {
		fmt.Println("Pesan terverifikasi")
	}
}