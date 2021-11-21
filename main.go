package main

import (
	"flag"
	"log"
	"net/http"

	"gitlab.com/avarf/getenvs"
)

var (
	flagBindTo = flag.String("bind-to", "localhost:9005", "bind to ")
	flagPath   = flag.String("path", "/dev/shm", "path to serve")
)

func main() {
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(getenvs.GetEnvString("PATH", "/dev/shm"))))
	log.Printf("Starting server on %s", getenvs.GetEnvString("SERVER_PORT", ":80"))
	log.Fatal(http.ListenAndServe(getenvs.GetEnvString("SERVER_PORT", ":80"), nil))
}
