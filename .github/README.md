# Go Service Weaver Template 🚀

This is a template project for creating a Service Weaver application, using sqlc and Fiber.

## Packages

- [Service Weaver](https://serviceweaver.dev/)
- [SQLC](https://docs.sqlc.dev/en/stable/overview/install.html)
- [Fiber](https://gofiber.io/)

## Databases

- [PostgreSQL](https://www.postgresql.org/)

### Architecture

This template was inspired by Clean Architecture and offers templates that follow its core concepts. The templates provided facilitate the creation of modular, maintainable, and testable codebases by promoting a clear separation of concerns and the independence of business logic from external dependencies. By using this templates, developers can jumpstart their projects with a well-organized structure that aligns with Clean Architecture, enabling them to focus on implementing the domain-specific logic while adhering to best practices. This template empowers developers to build robust and scalable applications, leveraging the benefits of Clean Architecture for easier understanding, maintenance, and evolution over time.

### Repository Structure

```bash
.
├── bin
│   └── app
├── cmd
│   └── app
│       └── main.go
├── deployments
│   └── compose.yml
├── docs
├── internal
│   ├── example
│   │   ├── db
│   │   │   ├── query.sql
│   │   │   ├── schema.sql
│   │   │   └── sqlc.yaml
│   │   ├── store
│   │   │   ├── db.go
│   │   │   ├── models.go
│   │   │   └── query.sql.go
│   │   ├── entity.go
│   │   ├── service.go
│   │   └── weaver_gen.go
│   └── frontend
│       ├── frontend.go
│       ├── router.go
│       └── weaver_gen.go
├── scripts
├── go.mod
├── go.sum
├── Makefile
└── weaver.toml
```

## Getting Started

### Installing dependencies

```sh
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
```

### Starting the services

```sh
make docker-run
```

### Building the project

```sh
make build
```

### Running the project

```sh
make run
```

## Commands

### Makefile

run: Run the application.
run-multi: Run the application services on multiple processes.
run-docker: Run the application services using Docker.
build: Build the application.

## Configuration

The configuration file is located at `weaver.toml`. You can modify the configuration settings as needed.
