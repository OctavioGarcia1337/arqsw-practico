<<<<<<< Updated upstream
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func main() {
	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.mercadolibre.com/sites/MLA", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var responseObject Response
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject)
}

//codigo viejo.
/* package main

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
*/

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
=======
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	ID  string `json:"id"`
	Url string `json:"url"`
}

func main() {
	fmt.Println("Calling API...")

	response, err := http.Get("https://api.thecatapi.com/v1/images/search")
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}

/*
func TestApi(t *testing.T) {
	if response != " " {
		t.Error()
	}
}

//codigo viejo.
/* package main

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
*/

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
>>>>>>> Stashed changes
