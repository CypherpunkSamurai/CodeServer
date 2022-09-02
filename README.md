# CodeServer
A Golang project that bridges LSP over websockets

## How Does this work?

This uses the `github.com/gorilla/websockets/` library to bridge `cmd.Exec` over websockets. This allows us to do cool things like run a LSP server and get it's output using websocket.

## How to build

1. clone the repo
```shell
git clone https://github.com/CypherpunkSamurai/CodeServer.git
```
2. Run the go build
```shell
go build
```
3. Run the binary
```shell
./codeserver
# or for windows
codeserver.exe
```

# Credits

* Golang Developers
* Gorilla Websocket Library

