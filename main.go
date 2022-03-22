package main

import (
	"fmt"
	"sanguo_game/config"
	"sanguo_game/wsServer"
)

const (
	defaultHost string = "localhost"
	defaultPort int    = 4396
)

func main() {
	host := config.File.MustValue("server", "host", defaultHost)
	port := config.File.MustInt("server", "port", defaultPort)
	fmt.Println(host, port)
	wsServer.RunWsWebServer()
}
