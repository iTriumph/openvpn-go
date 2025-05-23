# OpenVPN Client (Go Module)
A lightweight, extensible Go module that wraps the OpenVPN CLI — designed for API and CLI-based VPN management. It provides secure connection handling, automatic cleanup, and contextual lifecycle control.

---

## 🚀 Features
- ✅ Start/Stop/Reconnect OpenVPN securely
- 🔐 Internal handling of username & password (never exposed)
- 🧹 Temporary file cleanup after use
- 📡 Live log & status streaming
- ⛔ Custom error types (`ErrTimeout`, `ErrAlreadyRunning`, etc.)
- ⚙️ Built-in tests and automation via `Makefile`

---

## 📦 Installation
```bash
go get github.com/smantel-ch/openvpn-go
```


## ✨ Example Usage
```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/smantel-ch/openvpn-go"
)

func main() {
	config := []byte("...your .ovpn content...")

	client, err := openvpn.NewVPNClient()
	if err != nil {
		log.Fatal("init error:", err)
	}

    client.SetConfig(config)
    client.SetCredentials("myuser", "mypass")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatal("connection failed:", err)
	}

	fmt.Println("VPN Status:", client.Status())
	client.Disconnect()
}
```
> ⚠️ **Note**: `SetConfig()` expects the full `.ovpn` config file as **bytes** (not a path). <br />Use `os.ReadFile("my.ovpn")` to load it.


## 🐛 Debug Logging
To view debug output from the OpenVPN client, you must provide a logger that implements the following interface:

```go
type Logger interface {
    Debugf(format string, args ...any)
}
```
> ℹ️ **Info**: Only `Debugf()` logs are emitted by this package. <br />All functional errors and connection statuses are returned via Go errors and status channels.

## 🧪 Testing
Run the test suite:
```bash
make test
```

Run tests with coverage:
```bash
make test-cover
```


## 🔧 Dev Commands
```bash
make # runs fmt, lint, test
make fmt # gofmt formatting
make lint # golangci-lint
make build # builds CLI (cmd/demo)
make ci # full local pipeline check
```


## 🖥️ Demo CLI
```bash
go run ./demo/main.go -user myuser -pass mypass -config my.ovpn
```
