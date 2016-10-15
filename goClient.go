package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	var authHeader = string("TSM ")

	var ciphertext = ExampleNewGCMEncrypter([]byte("exampleplaintext"))



	println(authHeader)

	client := http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/test", nil)
	req.Header.Add("Authorization", "TSM "+ciphertext)
	res, _ := client.Do(req)

	resBody, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal("Error - %s\n", err)
	}
	fmt.Printf("%s\n", resBody)

	req, _ = http.NewRequest(http.MethodGet, "http://localhost:8080/session", nil)
	req.Header.Add("Authorization", "TSM ")
	res, _ = client.Do(req)

	resBody, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal("Error - %s\n", err)
	}
	fmt.Printf("%s\n", resBody)

}
