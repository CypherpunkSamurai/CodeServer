package api

import (
	"codeserver/term"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Configs
const buffsize_const = 1 * 1024

/*
	Websocket Connection Upgrader

	- CheckOrigin disabled to allow localhost connections
*/
var upgrader = websocket.Upgrader{
	ReadBufferSize:  buffsize_const,
	WriteBufferSize: buffsize_const,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Websocket HTTP Request Handler
func WSHandler(w http.ResponseWriter, h *http.Request) {
	// Handles the Websocket http request
	vars := mux.Vars(h)
	if !term.CheckCmd(vars["cmd"]) {
		// return json error
		w.Header().Add("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["error"] = "command not found"
		json_resp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Cannot marshall json")
			return
		}
		w.Write(json_resp)
		return
	}
	// Start a Cmd Bridge
	WSUpgrade(w, h, vars["cmd"])
	return
}

// Upgrades the HTTP Connection to WS
func WSUpgrade(w http.ResponseWriter, h *http.Request, command string) {

	log.Printf("[+] New Connection.... [%s]\n", h.RemoteAddr)

	// Fix Command string
	if !strings.Contains(command, ".exe") && term.IsWindows() {
		command = command + ".exe"
	}
	log.Printf("[+] Command: [%s]\n", command)

	con, err := upgrader.Upgrade(w, h, nil)

	if err != nil {
		log.Printf("Unable to upgrade connection....\n")
		log.Printf("Error: %s", err)
	}

	// Start the BridgeCMD
	BridgeCmd(command, con)
}
