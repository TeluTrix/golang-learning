package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

var reqUrl = "https://irefurb.ch/"

func main() {
	resp, err := http.Get(reqUrl)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)
}
