package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var makeFirstString bool = true
var urlString string
//var previous *string
//var next *string

func setUrlString() {
	if makeFirstString {
		urlString = "https://pokeapi.co/api/v2/location-area"
	}
	makeFirstString = false

}

func commandMap() error {
	setUrlString()
	fmt.Println("here are the next twenty locaiton areas in pokemon world")
	fmt.Println("taken from: ", urlString)
	res, err := http.Get(urlString)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		msg := fmt.Sprintf("Response failed with status code: %v and \nbody: %s \n", res.StatusCode, body)
		log.Fatal(msg)
	}
	if err != nil {
		log.Fatal(err)
	}

	var unmarshalledBody config
	err = json.Unmarshal(body, &unmarshalledBody) //unmarshall form json to struct config
	if err != nil {
		fmt.Println(err)
	}

	//update the urlString so that the next map command shows the next twenty locations
	fmt.Println()
	fmt.Println(urlString)
	urlString = unmarshalledBody.Next
//	*next = unmarshalledBody.Next
//	*previous = *unmarshalledBody.Previous
	fmt.Println()
	fmt.Println(urlString)

	for _, value := range unmarshalledBody.Results {
		fmt.Println(value.Name)
	}
//	urlString = *next
//	next = nil
//	*previous = urlString
	return nil
}
