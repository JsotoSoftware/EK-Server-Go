# EK-Server-Go

This is an authoritative server for online card game play, built with Go. The server handles game state management and real-time player interactions using WebSocket connections.

## 🚀 Features

- Real-time WebSocket communication
- Game state management
- Multi-player support
- Secure connection handling

## 📋 Prerequisites

- Go 1.21 or higher
- Git

## 🛠️ Installation

1. Clone the repository
```bash
git clone https://github.com/JsotoSoftware/EK-Server-Go.git
cd EK-Server-Go
```

2. Install dependencies
```bash
go mod download
```

3. Set up environment variables
```bash
cp .envtemplate .env
```

## ⚙️ Configuration

The following environment variables are required:

| Variable | Description | Default |
|----------|-------------|---------|
| PORT | Server port number | 8080 |

## 🚀 Running the Server

Development mode:
```bash
go run cmd/main.go
```

Build and run:
```bash
go build -o server cmd/main.go
./server
```

## 🏗️ Project Structure

```
EK-Server-Go/
├── cmd/            # Application entry points
├── internal/       # Private application code
│   └── network/    # Network handling and WebSocket code
├── .env           # Environment variables (not in git)
├── .envtemplate   # Environment template
└── go.mod         # Go module definition
```

## 📫 Contact

Josue Soto - jsoto.consultant@gmail.com

Project Link: [https://github.com/JsotoSoftware/EK-Server-Go](https://github.com/JsotoSoftware/EK-Server-Go)
