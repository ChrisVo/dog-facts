package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DogFactResponse struct {
	Facts   []string `json:"facts"`
	Success bool     `json:"success"`
}

func main() {
	resp, err := http.Get("https://dog-api.kinduff.com/api/facts")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var dogFactResponse DogFactResponse
	json.Unmarshal(body, &dogFactResponse)

	fmt.Printf(`
	   / \__
	  (    @\___
	  /         O       
	 /   (_____/
	/_____/   U
		
		
	%v
   `, dogFactResponse.Facts[len(dogFactResponse.Facts)-1])

}