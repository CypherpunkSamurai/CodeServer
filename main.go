// codersock
// cypherounksamurai

package main

import (
	"log"
	"net/http"

	"codeserver/api"

	"github.com/gorilla/mux"
)

// Configs
var server_host string = "127.0.0.1:5000"

func main() {

	// GorillaMux Routing
	r := mux.NewRouter()
	r.HandleFunc("/bridge/{cmd}", api.WSHandler)

	// Assing the Handlers
	http.Handle("/", r)

	// Start HTTP Server
	log.Printf("Starting http server....\n")
	err := http.ListenAndServe(server_host, nil)

	// Check for Errors
	if err != nil {
		log.Fatalf("Error starting http. %s\n", err)
	}

	// Works
	log.Printf("Http server is now running on %s....\n", server_host)
}
