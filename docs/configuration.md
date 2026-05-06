# Configuration

## Environment Variables

| Variable | Required | Description |
| --- | --- | --- |
| `DB_DSN` | Local only | SQL Server connection string. |
| `DB_DSN_FILE` | Docker/production | Path to a secret file containing the SQL Server connection string. Takes precedence over `DB_DSN`. |
| `PORT` | No | HTTP port. Defaults to `8080`. |
| `API_KEYS` | Local only | Plain API keys in `key:scope` format. |
| `API_KEYS_FILE` | Docker/production | Path to a secret file containing bcrypt-hashed API keys in `hash:scope` format. Takes precedence over `API_KEYS`. |

Supported API key scopes:

| Scope | Permissions |
| --- | --- |
| `read` | `GET` endpoints. |
| `crud` | All `GET`, `POST`, `PUT`, and `DELETE` endpoints. |

## Local Development

Copy the example file:

```bash
cp .env.example .env
```

Example `.env`:

```env
DB_DSN=sqlserver://username:password@host:1433?database=dbname
API_KEYS=read-key:read,crud-key:crud
PORT=8080
```

Run:

```bash
go run .
```

## Docker Secrets

Docker Compose uses secret files instead of plain environment variables for database and API credentials.

Create the database DSN secret:

```bash
mkdir -p secrets
printf '%s\n' 'sqlserver://username:password@host:1433?database=dbname' > secrets/db_dsn
```

Create bcrypt-hashed API key secrets:

```bash
go run ./cmd/hash-key --scope read your-read-key > secrets/api_keys
go run ./cmd/hash-key --scope crud your-crud-key >> secrets/api_keys
```

Start:

```bash
docker-compose up --build
```

Docker Compose uses `PORT` from `.env` for both the container listen port and the published host port. For example, `PORT=8888` publishes the API at `http://localhost:8888`.

`docker-compose.yml` mounts:

| Local file | Container path | Environment variable |
| --- | --- | --- |
| `./secrets/db_dsn` | `/run/secrets/db_dsn` | `DB_DSN_FILE` |
| `./secrets/api_keys` | `/run/secrets/api_keys` | `API_KEYS_FILE` |

The `secrets/` directory is ignored by Git and Docker build context.

## API Key File Format

Local `API_KEYS` uses plaintext keys:

```text
read-key:read,crud-key:crud
```

Docker `API_KEYS_FILE` uses bcrypt hashes:

```text
$2a$10$...:read
$2a$10$...:crud
```

Clients always send the original plaintext key:

```http
Authorization: Bearer read-key
```

## Database

The service expects these Anviz tables to already exist:

| Table | Model |
| --- | --- |
| `Userinfo` | users |
| `Dept` | departments |
| `FingerClient` | devices |
| `Checkinout` | attendance records |
| `Status` | check types |

No migrations are executed. Schema ownership stays with the Anviz database/application.
