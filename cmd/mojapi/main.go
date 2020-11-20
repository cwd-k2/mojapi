package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	api := "https://www.mojamoja.cloud/api/v1/environment/latest"

	for {
		resp1, err := http.Get(api)
		if err != nil {
			panic(err)
		}
		defer resp1.Body.Close()

		byteArray, err := ioutil.ReadAll(resp1.Body)
		println(string(byteArray))

		time.Sleep(time.Second * 10)
	}
}
