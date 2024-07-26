// codersock
// cypherounksamurai

package main

import (
	"codeserver/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", api.ServeWebContent)
	r.HandleFunc("/bridge/{cmd}", api.CmdBridgeHandler)
	http.Handle("/", r)
	// start
	server := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
	}
	log.Fatalln(server.ListenAndServe())
}
