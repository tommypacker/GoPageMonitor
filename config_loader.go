package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Website represents the initial url to load and
// the url that denotes that the site is live
type Website struct {
	Name       string `json:"name"`
	InitialURL string `json:"initial_url"`
	LiveURL    string `json:"live_url"`
}

// Websites is a struct holding a list of Website structs
type Websites struct {
	List []Website `json:"websites"`
}

// LoadWebsites will load the list of websites to monitor
func LoadWebsites() []Website {
	jsonFile, err := os.Open("websites.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var websites Websites

	json.Unmarshal(byteValue, &websites)
	return websites.List
}
