# Development Guide

## Prerequisites

- Go 1.24+
- Node.js 18+
- Wails CLI v2

```bash
# Install Wails CLI
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# Check environment dependencies
wails doctor
```

## Development Mode

```bash
# Install frontend dependencies
cd cmd/desktop/frontend && npm install

# Start development mode (with hot reload)
cd ..
wails dev
```

## Build for Release

```bash
# Generate desktop frontend dist after frontend changes
cd cmd/desktop/frontend && npm run build

# Build desktop app
cd ..
wails build -platform darwin/universal
wails build -platform windows/amd64

# Build server mode
cd ../../cmd/server && go build -ldflags="-s -w" -o ainexus-server .
```

Desktop build output is in `cmd/desktop/build/bin/`.

## Project Structure

```
AINexus/
├── cmd/
│   ├── desktop/            # Wails desktop app
│   │   ├── app.go          # Desktop app logic
│   │   ├── main.go         # Desktop entry
│   │   └── frontend/       # Vue/native modular frontend
│   ├── server/             # Headless server mode
│   └── license-server/     # Online license service
├── internal/
│   ├── proxy/              # HTTP proxy core
│   ├── transformer/        # API format transformers
│   ├── storage/            # SQLite data storage
│   ├── service/            # Endpoint, stats, backup, and update services
│   ├── config/             # Configuration management
│   ├── onlinelicense/      # Online license cards and device activation
│   ├── webdav/             # WebDAV sync
│   ├── logger/             # Logging system
│   └── tray/               # System tray
└── docs/                   # Configuration, development, FAQ, and model API docs
```
