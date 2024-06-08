package main

import "log"

// Run - responsible for the instantiation and startup
func Run() error {
	log.Println("Starting up API")
	return nil
}

func main() {
	log.Println("Welcome!")
	if err := Run(); err != nil {
		log.Println(err)
	}
}
