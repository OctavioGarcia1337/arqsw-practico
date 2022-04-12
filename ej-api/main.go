package main

import "github.com/MatiasLessio/ej-apis/apiCall"

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
