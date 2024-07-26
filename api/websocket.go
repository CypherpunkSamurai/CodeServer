package api

import (
	"codeserver/pkg"
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
	// Check Origin Enabled for Debugging
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Websocket HTTP Request Handler
func CmdBridgeHandler(w http.ResponseWriter, h *http.Request) {
	// Handles the Websocket http request
	vars := mux.Vars(h)
	// read
	command := vars["cmd"]
	log.Printf("Starting LSP with cmd: %s", command)

	log.Printf("[+] New Connection.... [%s]\n", h.RemoteAddr)

	// Fix Command string
	if !strings.Contains(command, ".exe") && pkg.IsWindows() {
		command = command + ".exe"
	}
	log.Printf("[+] Command: [%s]\n", command)

	con, err := upgrader.Upgrade(w, h, nil)

	if err != nil {
		log.Printf("Unable to upgrade connection....\n")
		log.Printf("Error: %s", err)
	}

	// Start the BridgeCMD
	pkg.CmdBridgeStart(command, con)
}
