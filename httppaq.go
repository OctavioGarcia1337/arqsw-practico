package main

import {
	"net/http"
	"fmt"
	"errors"
	"encoding/json"
	"ioutil" //ioutil
}

type Categories []Category


type Category struct{
	ID string `json:"id"`
	Name string `json:"name"`
}


func main(){

	url := "https://api.mercadolibre.com/sites/MLA/search?q=Motorola"
	cats, err := GetCategories(url)
	if err!=nil {
		//completar
	}
	fmt.Println("Las categorias de MLA son: ", Categories)
}

func GetCategories(siteID string)(Categories, error{
	response := http.Get(siteID)
	
	bytes := ioutil.ReadAll(response.Bytes) //completar

	var catgs Categories
	json.Unmarshal(bytes, &catgs)

	return cats, nil
}




/*
		url := ""
		responseGet := http.Get(url)
		responsePost := http.Post(url, body)

		//data, err

		var myVar myStruct
		err := json.Unmarshal (bytes, &myVar)
		fmt.Println(resp)
*/



/*
func main(){
	url := ""
	responseGet := http.Get(url)
	responsePost := http.Post(url, body)

	//data, err

	var myVar myStruct
	err := json.Unmarshal (bytes, &myVar)
	fmt.Println(resp)
}

 

	{
		"nombre": "Pollo",
		"edad": 21
	}


type Persona struct{
	Nombre string
	Edad int
}
*/