package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ConvertBytestoGibiBytes(x int) int {
	y := x / 1024 / 1024 / 1024
	return y
}

func ConvertBytestoGigaBytes(x int) int {
	y := x / 1000 / 1000 / 1000
	return y
}

func PrintOutJsonBody(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	fmt.Println(sb)
}
