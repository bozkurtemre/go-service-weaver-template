# Service Weaver Template

## Introduction

This is a template project for creating a service weaver application. The project is written in Go and uses the [Fiber](https://gofiber.io)

## Getting Started

### Requirements

- Docker and Docker Compose
- [Go](https://golang.org)
- [Service Weaver CLI](https://serviceweaver.dev)
- [Sqlc](https://sqlc.dev)

### Setup and Running

1. **Clone the Project**: Clone the project from GitHub.

   ```bash
   git clone https://github.com/bozkurtemre/go-service-weaver-template.git 
   ```

2. **Navigate to the Project Directory**: Move to the project directory.

   ```bash
   cd go-service-weaver-template
   ```

3. **Start the Services**: Start the services using Docker Compose.

   ```bash
   make run-docker
   ```

4. **Build the Project**: Build the project using the Service Weaver CLI.

   ```bash
   make build
   ```

5. **Run the Project**: Run the project using the Service Weaver CLI.

   ```bash
   make run
   ```

Now you can access the api at `http://localhost:8080`.
