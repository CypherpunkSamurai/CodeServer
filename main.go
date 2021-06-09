// codersock
// cypherounksamurai

package main

import (
        "bufio"
        "fmt"
        "io"
        "log"
        "net/http"
        "os"
        "os/exec"

        "github.com/gorilla/websocket"
)


// config
var server_host string = "127.0.0.1:5000"


var upgrader = websocket.Upgrader{
    ReadBufferSize:   1024,
    WriteBufferSize:  1024,
    CheckOrigin: func(r *http.Request) bool {return true}}

// var upgrader = websocket.Upgrader{
//     ReadBufferSize:  10 * 1024,
//     WriteBufferSize: 10 * 1024,
//     CheckOrigin: func(r *http.Request) bool {return true}}


func handle_stdout(r io.Reader, c *websocket.Conn) {
    // get stdout and send to client8
    // scanner doesn't scan empty lines.
    // so use buffer
    buf := make([]byte, 1024*1024*10)
    for {
        // loop reading the buffer
        sb,err := r.Read(buf)

        if err != nil {
            // stdout error
            log.Printf("Error Reading stdout. %s")
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

func handle_stdin(w io.WriteCloser, c *websocket.Conn, p *exec.Cmd) {
    // get socket data and send to cmd
    for {
        _, b,err := c.ReadMessage()
        if err != nil {
            log.Printf("[x] Connection force closed....\n")
            p.Process.Kill()
            return
        }
        fmt.Printf("<---[RECV] %s\n", string(b))
        w.Write([]byte(string(b) + "\n"))
    }
}

func ws(w http.ResponseWriter, h *http.Request) {

    fmt.Printf("[+] New Connection.... [%s]\n", h.RemoteAddr)

    con, err := upgrader.Upgrade(w,h,nil)

    if err != nil {
        log.Printf("Unable to upgrade connection....\n")
    }


        cmd := exec.Command("sh", "-c", "pyls", "-v")
        cmd.Env = os.Environ()

        // std
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

        defer stdout.Close()
        defer stdin.Close()
        defer stderr.Close()

        // writer := bufio.NewWriter(stdin)
        // reader_o := bufio.NewReader(stdout)
        // reader_e := bufio.NewReader(stderr)

        // reader := io.MultiReader(stdout,stderr)
        reader := bufio.NewReader(stdout)


        go handle_stdout(reader, con)
        go handle_stdin(stdin,con, cmd)


        err = cmd.Run()

        if err != nil {
            log.Printf("Cannot start process. %s", err)
        }

}



func main() {

    http.HandleFunc("/", ws)

    log.Printf("Starting http server....\n")
    err := http.ListenAndServe(server_host, nil)

    if err != nil {
        log.Fatalf("Error starting http. %s\n", err)
    }
}
