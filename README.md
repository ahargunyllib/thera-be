# Thera BE

![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)

## Description

This is a backend application for Thera BE project. It is designed following the best practices and recommendations to ensure a clean, maintainable, and scalable codebase. It is also built with security in mind to protect against common security threats.

The application is built on [Go v1.24.3](https://tip.golang.org/doc/go1.24) and [PostgreSQL](https://www.postgresql.org/). It uses [Fiber](https://docs.gofiber.io/) as the HTTP framework and [pgx](https://github.com/jackc/pgx) as the driver and [sqlx](github.com/jmoiron/sqlx) as the query builder.

## Getting started

### Run in local environment

> **Note:** This project uses [Task](https://taskfile.dev/) as a task runner. You can also run the commands manually if you don't want to use Task. Just copy the command from the `Taskfile.yml` file.

1. Ensure you have [Go](https://go.dev/dl/) 1.24 or higher installed on your machine:

   ```bash
   go version && task --version # windows
   go version && go-task --version # unix
   ```

2. Create a copy of the `.env.example` file and rename it to `.env`:

   ```bash
   cp ./config/.env.example ./config/.env
   ```

   Update configuration values as needed.

3. Install all dependencies:

   ```bash
   task
   ```

4. Initialize `air` configuration:

	 ```bash
	 air init

	 # Update the `air.toml` file with the following configuration:
	 # [build]
	 # cmd = "go build -o ./tmp/main ./cmd/app" # for unix
	 # cmd = "go build -o ./tmp/main.exe ./cmd/app" # for windows
	 ```

5. Run migrations:

	 ```bash
	 task migrate:up
	 ```

6. Seed the database:

	 ```bash
	 task db:seed
	 ```

7. Run the project in development mode:

   ```bash
   task dev
   ```

### Run in Docker environment
1. Ensure you have [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/) installed on your machine:

   ```bash
   docker --version && docker-compose --version
   ```

2. Create a copy of the `.env.example` file and rename it to `.env`:

   ```bash
    cp ./config/.env.example ./config/.env
    ```

    Update configuration values as needed. Ensure the `DB_HOST` is set to `db`.

3. Build and run the project:

   ```bash
   go-task service:run
   ```

4. Run migrations:

   ```bash
   go-task service:db:migrate
   ```

   Ensure the `DB_HOST` is set to `localhost`.

5. Seed the database:

   ```bash
    go-task db:seed
    ```

    Ensure the `DB_HOST` is set to `localhost`.

> **Note:** If you encounter host resolution issues, try to update the `DB_HOST` value to `localhost` or `db` in the `.env` file.

## Documentation

For database schema documentation, see [here](https://dbdocs.io/ahargunyllib/thera-be), powered by [dbdocs.io](https://dbdocs.io/).

For API documentation, see [here](https://nhppsttnad.apidog.io), powered by [Apidog](https://apidog.com/).

## Architecture

The project is structured following the Clean Architecture, Layered Architecture, Domain-Driven Design principles, Hexagonal Architecture, and SOLID principles.

## Features

- **Migration**: database schema migration using [golang-migrate](https://github.com/golang-migrate/migrate)
- **Validation**: request data validation utilizing [Package validator](https://github.com/go-playground/validator)
- **Logging**: implemented with [zerolog](https://github.com/rs/zerolog)
- **Testing**: unit and integration tests powered by [Testify](https://github.com/stretchr/testify) with formatted output using [gotestsum](https://github.com/gotestyourself/gotestsum)
- **Error handling**: centralized error management system
- **Email functionality**: implemented using [Gomail](https://github.com/go-gomail/gomail)
- **Environment variables**: managed with [Viper](https://github.com/spf13/viper)
- **Security**: HTTP headers secured by [Fiber-Helmet](https://docs.gofiber.io/api/middleware/helmet)
- **CORS**: Cross-Origin Resource-Sharing enabled through [Fiber-CORS](https://docs.gofiber.io/api/middleware/cors)
- **Compression**: gzip compression provided by [Fiber-Compress](https://docs.gofiber.io/api/middleware/compress)
- **Linting**: code quality ensured with [golangci-lint](https://golangci-lint.run)
- **Docker support**
