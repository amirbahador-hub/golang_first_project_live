package main

import (
	"bytes"
	"io"
	"fmt"
	"net/http"
)


func createUserRequest() {
	url := "http://localhost:8083/users/"

	jsonBody := []byte(`{"name": "Benyamin", "password":"abcd658"}`)
	bodyReader := bytes.NewReader(jsonBody)
	response, err := http.Post(url, "application/json", bodyReader)
	defer response.Body.Close()
	json_response, _ := io.ReadAll(response.Body)
	fmt.Println(string(json_response))
	fmt.Println(err)
}

func login() {
	url := "http://localhost:8083/users/login/"

	jsonBody := []byte(`{"name": "Benyamin", "password":"abcd658"}`)
	bodyReader := bytes.NewReader(jsonBody)
	response, err := http.Post(url, "application/json", bodyReader)
	defer response.Body.Close()
	json_response, _ := io.ReadAll(response.Body)
	fmt.Println(string(json_response))
	fmt.Println(err)
}

func verify() {
	url := "http://localhost:8083/users/verify/"

	jsonBody := []byte(`{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQmVueWFtaW4iLCJyb2xlIjoiIiwiZXhwIjoxNzIzMjM5NjcyfQ.jnW8ln1wnbdxgbROMYh3jCKI7tNom8TXiS93p7WeqeY"}`)
	bodyReader := bytes.NewReader(jsonBody)
	response, err := http.Post(url, "application/json", bodyReader)
	defer response.Body.Close()
	json_response, _ := io.ReadAll(response.Body)
	fmt.Println(string(json_response))
	fmt.Println(err)
	fmt.Println((response.StatusCode))
}


func main(){
	verify()
}