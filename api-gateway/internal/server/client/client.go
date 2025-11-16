package client

import (
	"api/config"
	"fmt"
	"log"
)

var (
	Cinema  string
	Booking string
)

func CinemaService() {
	srvc, err := config.GetServiceConfig()
	if err != nil {
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

func BookingService() {
	srvc, err := config.GetServiceConfig()
	if err != nil {
		log.Println(err)
	}
	var client = config.Http2Config()
	var url_booking = srvc.Booking
	resp, err := client.Get(url_booking + "/")
	if err != nil {
		log.Fatalf("Failed Connect Booking: %v", err)
	}
	defer resp.Body.Close()

	if resp.ProtoMajor == 2 {
		fmt.Println("Booking service connected....")
	}

	Booking = url_booking
}
