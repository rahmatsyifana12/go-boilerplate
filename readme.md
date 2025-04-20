# 🧱 Go Boilerplate App

A modular, production-ready Golang boilerplate application that supports local and Dockerized development environments. It integrates with PostgreSQL and Redis, includes JWT authentication, and supports database migration and testing tools.

---

## 📦 Tech Stack

- **Golang** `1.24`
- **PostgreSQL**
- **Redis**
- **Docker & Docker Compose**
- **JWT Authentication**
- **Air** (Live reload for development)
- **Migrate** (Database migrations)
- **Mockery** (Test mocks)
- **Supervisor** (Optional for deployment)

---

## ✨ Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/go-boilerplate-app.git
cd go-boilerplate-app
```

### 2. Create and Configure `.env`

Duplicate the `.env` file and fill in your environment-specific values:

```env
PORT=5000
ENVIRONMENT=local

POSTGRES_USERNAME=your_username
POSTGRES_PASSWORD=your_password
POSTGRES_DB_NAME=your_db
POSTGRES_PORT=5432
POSTGRES_HOST=localhost

REDIS_USERNAME=
REDIS_PASSWORD=
REDIS_PORT=6379
REDIS_HOST=localhost
REDIS_DB=0

JWT_ACCESS_SECRET=your_jwt_secret
CORS_ALLOWED_ORIGINS=http://localhost:3000
```

---

## 🐳 Running with Docker

Make sure Docker and Docker Compose are installed. Then run:

```bash
docker-compose up --build
```

> Ensure the Docker network `app-network` exists or remove `external: true` from the `docker-compose.yml` to let Docker create it automatically.

---

## 💻 Running Locally

Make sure PostgreSQL and Redis are running and match your `.env` configuration.

### Build and Run

```bash
make build
make start
```

### Or use live reload in development

```bash
make dev
```

> Requires [Air](https://github.com/cosmtrek/air) to be installed.

---

## 🔧 Available Make Commands

| Command               | Description                               |
|-----------------------|-------------------------------------------|
| `make build`          | Build the binary                          |
| `make run`            | Run using `go run`                        |
| `make start`          | Run the compiled binary                   |
| `make dev`            | Run with live reload via Air             |
| `make compile`        | Cross compile for Linux & Windows         |
| `make deploy`         | Build and restart via Supervisor (Linux) |
| `make migration-up`   | Run database migrations                   |
| `make migration-down` | Rollback all migrations                   |
| `make migration-down-1` | Rollback last migration                |
| `make migration-create` | Create new migration (prompt input)    |
| `make test`           | Run tests with coverage output            |
| `make mocks`          | Generate mocks using script               |

---

## 🧪 Testing

```bash
make test
```

HTML coverage report will be available at `./coverage/coverage.html`.

---

## 📁 Directory Structure

```
.
├── Dockerfile
├── docker-compose.yml
├── .env
├── Makefile
├── migrations/
├── scripts/
│   └── generate_mocks.sh
├── src/
│   ├── main.go
│   └── ...
└── coverage/
```

---

## 🛠 Troubleshooting

- Make sure your `.env` file is complete before running Docker or local commands.
- PostgreSQL and Redis must be reachable with the provided credentials.
- If Air is not installed: `go install github.com/cosmtrek/air@latest`

---

## 🧰 Tools

- [Air](https://github.com/cosmtrek/air) – Hot reload for Go
- [Mockery](https://github.com/vektra/mockery) – Interface mocking
- [Migrate](https://github.com/golang-migrate/migrate) – DB migration
- [Codecov](https://about.codecov.io/) – Test coverage reporting

---

## 📄 License

This project is licensed under the MIT License.

