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
	resp := httpRequestSonarr(client,url,key,"system/status","GET")
	var status model.SonarStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sonarr version: %s\n", status.Version)
}

func getSeries(client *http.Client, url string, key string){
	resp := httpRequestSonarr(client,url,key,"series","GET")
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
