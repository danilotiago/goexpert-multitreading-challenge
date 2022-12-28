package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	cep := "01001-000"
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		req, err := http.Get("https://cdn.apicep.com/file/apicep/" + cep + ".json")
		if err != nil {
			println(err)
		}

		res, err := io.ReadAll(req.Body)
		if err != nil {
			println(err)
		}

		ch1 <- string(res)

	}()

	go func() {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			println(err)
		}

		res, err := io.ReadAll(req.Body)
		if err != nil {
			println(err)
		}

		ch2 <- string(res)

	}()

	select {
	case res := <-ch1:
		fmt.Printf("Result %s from %s", res, "apicep\n")
	case res := <-ch2:
		fmt.Printf("Result %s from %s", res, "viacep\n")
	case <-time.After(time.Second * 1):
		fmt.Printf("timeout\n")
	}
}
