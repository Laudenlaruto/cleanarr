package cleanarr

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type sonarStatusResponse struct {
	Version           string    `json:"version"`
	Buildtime         time.Time `json:"buildTime"`
	Isdebug           bool      `json:"isDebug"`
	Isproduction      bool      `json:"isProduction"`
	Isadmin           bool      `json:"isAdmin"`
	Isuserinteractive bool      `json:"isUserInteractive"`
	Startuppath       string    `json:"startupPath"`
	Appdata           string    `json:"appData"`
	Osname            string    `json:"osName"`
	Osversion         string    `json:"osVersion"`
	Ismonoruntime     bool      `json:"isMonoRuntime"`
	Ismono            bool      `json:"isMono"`
	Islinux           bool      `json:"isLinux"`
	Isosx             bool      `json:"isOsx"`
	Iswindows         bool      `json:"isWindows"`
	Branch            string    `json:"branch"`
	Authentication    string    `json:"authentication"`
	Sqliteversion     string    `json:"sqliteVersion"`
	Urlbase           string    `json:"urlBase"`
	Runtimeversion    string    `json:"runtimeVersion"`
	Runtimename       string    `json:"runtimeName"`
}

func sonarStatus() {
	resp, err := http.Get("http://192.168.1.19:8989/api/system/status?apikey=e695509add214a7c8f7f20d829ea8d44")
	if err != nil {
		log.Fatalln(err)
	}

	var status sonarStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sonarr version: %s\n", status.Version)

}
