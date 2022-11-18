package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	cdn := make(chan string)
	viaCep := make(chan string)

	go func() {
		url := "https://cdn.apicep.com/file/apicep/13289-462.json"
		getCep(url, cdn)
	}()

	go func() {
		url := "https://viacep.com.br/ws/13289-462/json/"
		getCep(url, viaCep)
	}()

	select {
	case cdnCep := <-cdn:
		println("cdnCep:", cdnCep)
	case viaCepCep := <-viaCep:
		println("viaCepCep:", viaCepCep)
	case <-time.After(time.Second):
		println("timeout")
	}
}

func getCep(url string, ch chan<- string) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	ch <- string(body)
}
