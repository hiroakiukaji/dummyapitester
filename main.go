package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Square struct {
	Root   int `json:"root"`
	Square int `json:"square"`
}

func main() {
	fmt.Println("API Test")
}

func apitest(i int) string {
	url := "http://127.0.0.1:1323/square/" + strconv.Itoa(i)

	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := client.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}
