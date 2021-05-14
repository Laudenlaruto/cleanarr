package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


func sonarStatus(client *http.Client, url string, key string) {
	req, err := http.NewRequest("GET", url+"system/status", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("X-Api-Key", key)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	var status sonarr.sonarStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sonarr version: %s\n", status.Version)
}

func getSeries(client *http.Client, url string, key string){
	req, err := http.NewRequest("GET", url+"series", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("X-Api-Key", key)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatalln(err)
    }
//Convert the body to type string
    sb := string(body)
    fmt.Println(sb)
}
