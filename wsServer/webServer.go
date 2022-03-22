// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wsServer

import (
	"fmt"
	"log"
	"net/http"
	"sanguo_game/config"
)

const (
	defaultHost string = "localhost"
	defaultPort int    = 8080
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "wsServer/home.html")
}

func RunWsWebServer() {
	var host = config.File.MustValue("ws_server", "host", defaultHost)
	var port = config.File.MustInt("ws_server", "port", defaultPort)
	addr := fmt.Sprintf("%s:%d", host, port)
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	log.Printf("start ws web server at: %s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
