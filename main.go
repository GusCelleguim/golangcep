package main

import "github.com/go-resty/resty/v2"



func main() {
   cep := "07851000"
	// https://viacep.com.br/ws/01001000/json/
	client := resty.New()
	resp, err := client.R().Get("https://viacep.com.br/ws/"+cep+"/json/")
	if err != nil {
		panic(err)
	}
	println(resp.String())
}


// go get github.com/go-resty/resty/v2