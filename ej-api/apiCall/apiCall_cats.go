package apiCall

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ApiCall() (Cats, error) {
	resp, err := http.Get("https://api.thecatapi.com/v1/images/search")
	if err != nil {
		println(err)
		return Cats{}, err
	}
	bytes, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		println(er)
		return Cats{}, err
	}
	var cats Cats

	json.Unmarshal(bytes, &cats)
	return cats, nil
}

type Cats []struct {
	Breeds     []interface{} `json:"breeds"`
	Categories []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"categories"`
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
