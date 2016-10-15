package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"bytes"
)

func main() {
	//Login
	var encryptPassword = ExampleNewGCMEncrypter([]byte("exampleplaintext"))
	loginRequest := LoginRequest{"admin", string(encryptPassword[:]), 4}

	loginJson, err := json.Marshal(loginRequest)

	client := http.Client{}
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/session", bytes.NewBuffer(loginJson))
	res, _ := client.Do(req)

	resBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal("Error - %s\n", err)
	}
	fmt.Printf("%s\n", resBody)
	var ciphertext = ExampleNewGCMEncrypter([]byte("exampleplaintext"))
	var authHeader = string("TSM ")
	println(authHeader)

	req, _ = http.NewRequest(http.MethodGet, "http://localhost:8080/session", nil)
	req.Header.Add("Authorization", string(ciphertext[:]))
	res, _ = client.Do(req)

	resBody, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal("Error - %s\n", err)
	}
	fmt.Printf("%s\n", resBody)

}
