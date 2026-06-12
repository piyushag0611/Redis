package main

import (
	"flag"
	"log"

	"redis-db/config"
	"redis-db/server"
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the redis-db server")
	flag.IntVar(&config.Port, "port", 7379, "port for the dice server")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("Starting up Redis-DB")
	server.RunSyncTCPServer()
}
