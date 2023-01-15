# blog-app-service

### Installation

1.  Clone the repo and `cd` into the directory
2.  Make sure that `docker, PostgreSQL, go1.19` are installed.

### Running locally
- start:
  1. Run dependency services (database) with `docker compose up’.
  2. Migrate the database with `migrate -database "postgres://user:password@localhost:5432/blog-app?sslmode=disable" -path internal/database/migrations up`.
  3. Run `go run .` in the root folder.
  
- test: run `go test` for all the unit tests to run

### Frontend
https://github.com/edanimark/web-blog-app
