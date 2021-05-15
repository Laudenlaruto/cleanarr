package sonarr

import (
	"cleanarr/model"
	"cleanarr/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Status(client *http.Client, url string, key string) {
	resp := httpRequestSonarr(client, url, key, "system/status", "GET", nil)
	var status model.SonarStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		utils.PrintOutJsonBody(resp)
		log.Fatal(err)
	}
	fmt.Printf("Sonarr version: %s\n", status.Version)
}

func UsedSpace(client *http.Client, url string, key string) {
	resp := httpRequestSonarr(client, url, key, "series", "GET", nil)
	var series model.Series
	if err := json.NewDecoder(resp.Body).Decode(&series); err != nil {
		utils.PrintOutJsonBody(resp)
		log.Fatal(err)
	}
	total := 0
	for _, serie := range series {
		total = total + serie.Sizeondisk
	}

	GiB := utils.ConvertBytestoGigaBytes(total)
	fmt.Printf("Total space consumed for Series: " + strconv.Itoa(GiB) + " GB\n")
}
func ListEpisodeWithFilter(client *http.Client, url string, key string, days int) {
	series := getSeries(client, url, key)
	var episodesFile model.EpisodeFile
	// var episodes model.Episode
	for _, serie := range series {
		params := make(map[string]string)
		params["seriesId"] = strconv.Itoa(serie.ID)
		//Get episode file per serie ID
		resp := httpRequestSonarr(client, url, key, "episodefile", "GET", params)
		if err := json.NewDecoder(resp.Body).Decode(&episodesFile); err != nil {
			utils.PrintOutJsonBody(resp)
			log.Fatal(err)
		}
		for _, file := range episodesFile {
			// Filter on age
			if time.Since(file.Dateadded).Hours() > float64(days*24) {
				path := string("episode/" + strconv.Itoa(file.ID))
				var episode model.Episode
				resp2 := httpRequestSonarr(client, url, key, path, "GET", params)
				if err := json.NewDecoder(resp2.Body).Decode(&episode); err != nil {
					utils.PrintOutJsonBody(resp2)
					log.Fatal(err)
				}
				fmt.Println(serie.Title, episode.Title, "Season", episode.Seasonnumber, "Episode", episode.Episodenumber)
			}
		}
	}
}
func getSeries(client *http.Client, url string, key string) model.Series {
	resp := httpRequestSonarr(client, url, key, "series", "GET", nil)
	var series model.Series
	if err := json.NewDecoder(resp.Body).Decode(&series); err != nil {
		utils.PrintOutJsonBody(resp)
		log.Fatal("Error decoding series id")
	}
	return series
}

// Sonarr HTTP request
func httpRequestSonarr(client *http.Client, url string, key string, path string, methode string, params map[string]string) *http.Response {
	req, err := http.NewRequest(methode, url+path, nil)
	if err != nil {
		log.Fatalln("Error creating request")
	}
	req.Header.Add("X-Api-Key", key)
	if params != nil {
		paramsQuery := req.URL.Query()
		for index, param := range params {
			paramsQuery.Add(index, param)
		}
		req.URL.RawQuery = paramsQuery.Encode()
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("HTTP request error")
	}
	return resp
}
