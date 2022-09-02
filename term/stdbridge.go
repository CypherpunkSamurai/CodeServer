package term

import (
	"fmt"
	"io"
	"log"
	"os/exec"

	"github.com/gorilla/websocket"
)

// HTTP Handlers

func CmdStdOut(r io.Reader, c *websocket.Conn) {
	// get stdout and send to client8
	// scanner doesn't scan empty lines.
	// so use buffer
	buf := make([]byte, 1024*1024*10)
	for {
		// loop reading the buffer
		sb, err := r.Read(buf)

		if err != nil {
			// stdout error
			log.Printf("Error Reading stdout %s", err)
			break
		}
		s := buf[:sb]
		err = c.WriteMessage(websocket.TextMessage, s)
		if err != nil {
			return
		}
		fmt.Printf("[SEND]---> %s\n", string(s))
	}
}

func CmdStdin(w io.WriteCloser, c *websocket.Conn, p *exec.Cmd) {
	// get socket data and send to cmd
	for {
		_, b, err := c.ReadMessage()
		if err != nil {
			log.Printf("[x] Connection force closed....\n")
			p.Process.Kill()
			return
		}
		log.Printf("<---[RECV] %s\n", string(b))
		w.Write([]byte(string(b) + "\n"))
	}
}
