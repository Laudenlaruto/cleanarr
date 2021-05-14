package sonarr

import (
	"cleanarr/model"
	"cleanarr/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)


func Status(client *http.Client, url string, key string) {
	resp := httpRequestSonarr(client,url,key,"system/status","GET")
	var status model.SonarStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sonarr version: %s\n", status.Version)
}

func UsedSpace(client *http.Client, url string, key string){
	resp := httpRequestSonarr(client,url,key,"series","GET")
	var series model.Series
	if err := json.NewDecoder(resp.Body).Decode(&series); err != nil {
		log.Fatal(err)
	}
	total := 0
	for _,serie := range series {
		total = total + serie.Sizeondisk
	}
	
	GiB := utils.ConvertBytestoGibiBytes(total)
	fmt.Printf("Total space consumed for Series: "+ strconv.Itoa(GiB) + " GB\n")
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
