package main

import (
	"fmt"
	"io"
	"log"
	"encoding/json"
	"net/http"
)


func commandMap() error {
	fmt.Println("here are the next twenty locaiton areas in pokemon world")
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
	fmt.Println()
	fmt.Println(urlString)

	for _, value := range unmarshalledBody.Results {
		fmt.Println(value.Name)
	}
	return nil
}
