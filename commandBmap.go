package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
)

func commandBmap() error {
	fmt.Println("entered bmap")
	if reflect.TypeOf(urlString) == nil {
		fmt.Println("there are no more preivous location areas")
		fmt.Println("You're looking into non-existant path")
		makeFirstString = true
		setUrlString()
		return nil
	}
	fmt.Println("here are the next twenty locaiton previous areas in pokemon world")
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
	urlString = *unmarshalledBody.Previous
//	*previous = *unmarshalledBody.Previous //using *unmarshalled since the .Previous filed is a pointer to a string in config
//	*next = unmarshalledBody.Next
	fmt.Println()
	fmt.Println(urlString)

	for _, value := range unmarshalledBody.Results {
		fmt.Println(value.Name)
	}
//	urlString = *previous
//	previous = nil
//	*next = urlString
	return nil
}
