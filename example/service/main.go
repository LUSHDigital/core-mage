package main

import "time"

func main() {
	f := func() {
		time.Sleep(5 * time.Second)
	}
	for {
		f()
	}
}
