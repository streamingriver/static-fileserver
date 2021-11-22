package main

import (
	"log"
	"net/http"

	"gitlab.com/avarf/getenvs"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(getenvs.GetEnvString("FILESERVER_ROOT", "/dev/shm"))))
	log.Printf("Starting server on %s", getenvs.GetEnvString("FILESERVER_PORT", ":80"))
	log.Fatal(http.ListenAndServe(getenvs.GetEnvString("FILESERVER_PORT", ":80"), nil))
}
