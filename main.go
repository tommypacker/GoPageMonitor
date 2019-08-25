package main

import (
	"log"
	"net/http"
	"time"
)

func monitorSite(website Website) {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	var resp *http.Response
	var err error

	for true {
		log.Printf("Waiting for %s to go live\n", website.Name)
		resp, err = client.Get(website.InitialURL)
		if err != nil {
			log.Println(err.Error())
		} else {
			url := resp.Request.URL.String()
			if url == website.LiveURL {
				break
			}
		}
		time.Sleep(5 * time.Second)
	}

	defer resp.Body.Close()
	log.Printf("%s is live\n", website.Name)

	// TODO: Add alert functionality
}

func main() {
	for _, website := range LoadWebsites() {
		go monitorSite(website)
	}
	select {}
}
