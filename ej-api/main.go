package main

import "github.com/OctavioGarcia1337/arq-software/ej-api/apiCall"

func main() {
	cats, _ := apiCall.ApiCall()
	for _, c := range cats {
		println("URL:", c.URL)
	}
}

/*
Integrantes:
Angelone, Garcia, Llabres y Lessio
*/
