package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Started service... sleeping every 5 seconds")
	go func() {
		for {
			time.Sleep(5 * time.Second)
		}
	}()
	select {}
}
