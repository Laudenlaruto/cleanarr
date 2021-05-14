package main

import (
	"cleanarr/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	var status model.SonarStatusResponse
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
	var series model.Series
	if err := json.NewDecoder(resp.Body).Decode(&series); err != nil {
		log.Fatal(err)
	}
	total := 0
	for _,serie := range series {
		GiB := convertBytestoGibiBytes(serie.Sizeondisk)
		total = total + GiB
	}

	fmt.Printf("Total space consumed for Series: "+ strconv.Itoa(total) + "GiB\n")
}
