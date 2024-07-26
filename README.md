# CodeServer

![GoVersion](https://img.shields.io/github/go-mod/go-version/CypherpunkSamurai/CodeServer)
![Websockets](https://img.shields.io/badge/gorilla-websockets-blue)
![LSP](https://img.shields.io/badge/LSP-Language_Server_Protocol-green)

A Golang project that bridges LSP over websockets for VSCode like Autocompletion for Ace Editor (used by Replit), Monaco Editor (used by VSCode) and Code Mirror.

## 🤔 How Does this work?

This uses the `github.com/gorilla/websockets/` library to bridge `cmd.Exec` over websockets. This allows us to do cool things like run a LSP server and get it's output using websocket.

> ⚠️ **Note:** This project was created in 2020 out of my own need to code on a mobile device with LSP Autocompletion. The target editor for this "Ace Editor" had no public LSP Clients back then to interact with this websocket server, thus this project was abandoned in favour of Neovim.

## 🔨 How to build

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

## 😊 Credits

- Golang Developers
- Gorilla Websocket Library
