package main

import (
	"flag"
	"log"

	"redis-from-scratch/config"
	"redis-from-scratch/server"
)

func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the redis server")
	flag.IntVar(&config.Port, "port", 6969, "port for the redis server")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("starting the redis server")
	server.RunSyncTCPServer()
}
