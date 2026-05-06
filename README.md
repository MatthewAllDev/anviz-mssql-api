# Anviz MSSQL API

HTTP JSON API for reading and managing data in an existing Anviz Microsoft SQL Server database.

The service is intentionally thin: it does not own the database schema, does not run migrations, and acts as an authenticated access layer over the Anviz tables.

## Features

- JSON API over Anviz MSSQL tables.
- API-key authentication with `read` and `crud` scopes.
- Local `.env` configuration for development.
- Docker Compose deployment with database and API key secrets.
- Pagination on list endpoints.
- Partial updates: `PUT` only updates fields present in the request body.
- Sensitive model fields are excluded from JSON responses.

## Requirements

- Go 1.25 or newer.
- Existing Anviz Microsoft SQL Server database.
- Docker and Docker Compose for containerized deployment.

## Quick Start

Local development:

```bash
cp .env.example .env
go run .
```

Example `.env`:

```env
DB_DSN=sqlserver://username:password@host:1433?database=dbname
API_KEYS=read-key:read,crud-key:crud
PORT=8080
```

Docker:

```bash
cp .env.example .env
mkdir -p secrets
printf '%s\n' 'sqlserver://username:password@host:1433?database=dbname' > secrets/db_dsn
go run ./cmd/hash-key --scope read your-read-key > secrets/api_keys
go run ./cmd/hash-key --scope crud your-crud-key >> secrets/api_keys
docker-compose up --build
```

Docker Compose uses `PORT` from `.env` for both the container listen port and the published host port. For example, `PORT=8888` publishes the API at `http://localhost:8888`.

Smoke test:

```bash
curl -H "Authorization: Bearer read-key" \
  "http://localhost:8080/api/v1/users?limit=10"
```

## Documentation

- [Configuration](docs/configuration.md)
- [API Reference](docs/api.md)

## Development

```bash
go test ./...
go mod tidy
go run .
go run ./cmd/hash-key --scope read example-key
```

## Project Structure

```text
.
├── api/           HTTP server composition
├── api/auth/      API key loading and authorization middleware
├── api/handlers/  Resource HTTP handlers
├── api/httpx/     Shared HTTP request/response helpers
├── cmd/hash-key/  API key hashing utility
├── db/            Database connection and queries
├── docs/          User-facing documentation
├── models/        GORM models for existing Anviz tables
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── main.go
```

## Security Notes

- Do not commit `.env`, `secrets/`, or credential files.
- Docker deployments should use `DB_DSN_FILE` and `API_KEYS_FILE`, not `DB_DSN` and `API_KEYS`.
- API keys are not stored in the database.
- Sensitive fields such as user passwords, documents, contact details, images, card numbers, and device communication passwords are excluded from JSON responses.
- Internal database errors are logged server-side and returned to clients as generic errors.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE).
