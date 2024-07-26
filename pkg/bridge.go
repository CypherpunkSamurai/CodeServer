package pkg

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/gorilla/websocket"
)

// Bridge Command and Websocket
func CmdBridgeStart(command string, con *websocket.Conn) {

	// check if command exists
	if !CheckCmd(command) {
		con.WriteMessage(websocket.TextMessage, []byte(`{"error": "command not found"}`))
		con.Close()
		return
	}

	// Cmd Bridge
	cmd := exec.Command(command)
	cmd.Env = os.Environ()

	// cmd pipe
	stdin, err := cmd.StdinPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdin: %s", err.Error())
	}
	stdout, err := cmd.StdoutPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdout: %s", err.Error())
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("Error obtaining stderr: %s", err.Error())
	}

	// Flush when not in use
	defer stdout.Close()
	defer stdin.Close()
	defer stderr.Close()

	/*
		We dont use bufio writers and scanners as they remove the newline characters and
		long texts. Which would cause loss of parts of the output. Instead we use char buffers.
		// writer := bufio.NewWriter(stdin)
		// reader_o := bufio.NewReader(stdout)
		// reader_e := bufio.NewReader(stderr)
		// reader := io.MultiReader(stdout,stderr)
	*/
	reader := bufio.NewReader(stdout)

	// Streaming
	go CmdStdOut(reader, con)
	go CmdStdin(stdin, con, cmd)

	// Start the CMD
	err = cmd.Run()

	if err != nil {
		log.Printf("Cannot start process. %s", err)
	}
}

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
