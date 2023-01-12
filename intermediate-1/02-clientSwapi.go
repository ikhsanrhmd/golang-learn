package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type StarWarsPeople struct {
	Name     string   `json:"name"`
	Height   string   `json:"height"`
	EyeColor string   `json:"eye_color"`
	Film     []string `json:"films"`
}

func main() {
	response, _ := http.Get("https://swapi.dev/api/people/1")
	responseData, _ := ioutil.ReadAll(response.Body) //membaca data response
	defer response.Body.Close()

	var LukeSkywalker StarWarsPeople
	json.Unmarshal(responseData, &LukeSkywalker)

	fmt.Println(LukeSkywalker.Name)
	fmt.Println(LukeSkywalker.Height)
	fmt.Println(LukeSkywalker.EyeColor)
	for _, v := range LukeSkywalker.Film {
		fmt.Println(v)

	}
}
