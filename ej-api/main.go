package main

import "github.com/octaviogarcia1337/arq-software/ej-api/apiCall"

func main() {
	cats, _ := apiCall.ApiCall()
	for _, c := range cats {
		println("URL:", c.URL, "\nID:", c.ID)
	}
}
