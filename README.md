# Blog API

A RESTful blog post API built with Go, Gin, GORM, and PostgreSQL.

## Tech Stack

- **[Go](https://golang.org/)** — language
- **[Gin](https://github.com/gin-gonic/gin)** — HTTP framework
- **[GORM](https://gorm.io/)** — ORM
- **[PostgreSQL](https://www.postgresql.org/)** — database
- **[godotenv](https://github.com/joho/godotenv)** — environment variable loading

## Prerequisites

- Go 1.25+
- PostgreSQL running locally

## Setup

1. Clone the repository and install dependencies:

   ```bash
   go mod download
   ```

2. Copy the environment file and fill in your values:

   ```bash
   cp .env.example .env
   ```

   | Variable | Description                        | Example                          |
   |----------|------------------------------------|----------------------------------|
   | `PORT`   | Port the server listens on         | `3000`                           |
   | `DB_URL` | PostgreSQL connection string (DSN) | `host=localhost user=... dbname=gin_blog port=5432 sslmode=disable` |

3. Run database migrations:

   ```bash
   go run ./migrations/
   ```

4. Start the server:

   ```bash
   go run main.go
   ```

   The API will be available at `http://localhost:<PORT>`.

## API Endpoints

All endpoints are prefixed with `/posts`.

| Method   | Endpoint      | Description          |
|----------|---------------|----------------------|
| `POST`   | `/posts`      | Create a new post    |
| `GET`    | `/posts`      | List all posts       |
| `GET`    | `/posts/:id`  | Get a single post    |
| `PUT`    | `/posts/:id`  | Update a post        |
| `DELETE` | `/posts/:id`  | Delete a post        |

### Post fields

```json
{
  "title":  "string",
  "body":   "string",
  "author": "string",
  "likes":  0,
  "draft":  false
}
```

## Project Structure

```
.
├── controllers/        # Gin route handlers
│   └── postController.go
├── inits/              # App initialization (DB, env)
│   ├── db.go
│   └── envLoader.go
├── migrations/         # Standalone migration binary
│   └── migrations.go
├── middlewares/        # Custom Gin middlewares
├── models/             # GORM models
│   └── post.go
└── main.go
```

## Development

For live reload during development, use [CompileDaemon](https://github.com/githubnemo/CompileDaemon):

```bash
CompileDaemon -command="./blog"
```
