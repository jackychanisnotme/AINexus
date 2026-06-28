# AINexus Server / Docker Guide

AINexus can run as a headless HTTP proxy service for servers, NAS devices, Docker hosts, and remote development machines. Server mode keeps the same endpoint rotation, failover, transformer, model list, health check, and statistics APIs used by the desktop app, but runs without the Wails GUI.

## Quick Start

```bash
cd cmd/server
docker compose up -d --build
```

After startup, open:

```text
http://localhost:3021/ui/
```

The mapped host port depends on `cmd/server/docker-compose.yml`. The container listens on port `3000` by default.

## Recommended Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `AINEXUS_PORT` | HTTP listen port inside the container or host process | `3000` |
| `AINEXUS_DATA_DIR` | Data directory | `/data` |
| `AINEXUS_DB_PATH` | SQLite database path | `/data/ainexus.db` |
| `AINEXUS_LOG_LEVEL` | Log level, 0=debug, 1=info, 2=warn, 3=error | `1` |
| `AINEXUS_BASIC_AUTH_USERNAME` | Basic Auth username for server mode | `admin` |
| `AINEXUS_BASIC_AUTH_PASSWORD` | Basic Auth password for server mode | `change-me` |

Persist `/data` as a Docker volume or host directory so endpoint configuration, statistics, and runtime state survive container upgrades.

## Server Features

- Web management UI at `/ui/`
- Proxy health check at `/health`
- Model discovery through `/v1/models`
- Usage statistics through `/stats`
- Endpoint management APIs under `/api/`
- Server-Sent Events for live Web UI updates
- Docker health check support

## Web UI

The Web UI provides:

- Endpoint add/edit/delete/enable/disable controls
- Transformer and model configuration
- Current endpoint switching
- Daily, weekly, monthly, and historical statistics
- Endpoint connectivity testing
- Light and dark themes

## Common Operations

```bash
# View service logs
docker compose logs -f

# Restart service
docker compose restart

# Rebuild after local source changes
docker compose up -d --build

# Enter the running container
docker compose exec ainexus sh
```

## Troubleshooting

If the UI cannot be reached, check:

- The container is running: `docker ps`
- The host port mapping in `cmd/server/docker-compose.yml`
- The service health endpoint: `curl http://localhost:3021/health`
- Container logs: `docker compose logs -f`

If requests fail after startup, check:

- Endpoint API URL and API key
- Transformer choice, especially `openai2` for OpenAI Responses / Codex CLI
- Target model name for non-Claude upstreams
- API quota or upstream rate limits

## Security Notes

- Put the server behind HTTPS when exposing it outside localhost or a trusted LAN.
- Use Basic Auth or a reverse proxy access control layer for remote deployments.
- Do not expose API keys, access tokens, or refresh tokens in logs, screenshots, or public issue reports.
- Back up the SQLite database before major upgrades or migrations.
