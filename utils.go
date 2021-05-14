package main

import (
	"log"
	"net/http"
)

func convertBytestoGibiBytes(x int) int {
	y := x / 1024 / 1024 / 1024
	
	return y
}

// Sonarr HTTP request
func httpRequestSonarr(client *http.Client, url string, key string, path string, methode string) *http.Response{
	req, err := http.NewRequest(methode, url+path, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("X-Api-Key", key)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}
