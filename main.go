package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type KeyPair struct {
	CodeName   string `json:"code_name"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

func getKeys(ip string, port string) (keys KeyPair, err error) {
	caCert, err := ioutil.ReadFile("server.crt")
	if err != nil {
		return keys, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	url := fmt.Sprintf("https://%s:%s/getkey", ip, port)

	graveClient := &http.Client{
		Timeout: time.Second * 2,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return keys, err
	}

	req.Header.Set("User-Agent", "Gravemind")

	res, err := graveClient.Do(req)
	if err != nil {
		return keys, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(resBody, &keys)
	if err != nil {
		return keys, err
	}

	return keys, nil
}

func writePub(pub string, name string) {
}

func writePriv(priv string, name string) {
}

func usage() {
	help := "USAGE: ./main <ip> <port>"
	fmt.Println(help)
}

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		usage()
	}
}
