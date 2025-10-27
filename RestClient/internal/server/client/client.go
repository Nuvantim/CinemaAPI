package client

import (
	"fmt"
	"log"
	"api/config"
)

var Cinema string

func CinemaService() {
	srvc,err := config.GetServiceConfig()
	if err != nil{
		log.Println(err)
	}
	var client = config.Http2Config()
	var url_cinema = srvc.Cinema
	resp, err := client.Get(url_cinema + "/")
	if err != nil {
		log.Fatalf("Failed Connect Cinema: %v", err)
	}
	defer resp.Body.Close()

	if resp.ProtoMajor == 2 {
		fmt.Println("Cinema service connected....")
	}

	Cinema = url_cinema
}