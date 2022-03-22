package main

import (
	"fmt"
	"log"
	"sanguo_game/config"
)

func main() {
	fmt.Println(config.File)
	host, err := config.File.GetValue("server", "host")
	if err != nil {
		log.Fatal(err)
	}
	port, err := config.File.GetValue("server", "port")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(host, port)
}
