package api

import (
	"bufio"
	"codeserver/term"
	"log"
	"os"
	"os/exec"

	"github.com/gorilla/websocket"
)

func BridgeCmd(command string, con *websocket.Conn) {
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
	go term.CmdStdOut(reader, con)
	go term.CmdStdin(stdin, con, cmd)

	// Start the CMD
	err = cmd.Run()

	if err != nil {
		log.Printf("Cannot start process. %s", err)
	}
}
