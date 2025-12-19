
# E-Commerce Backend API

A production-grade RESTful API built with Go 1.2x, engineered for high performance, strict type safety, and clean architecture. This project demonstrates a deep understanding of the Dependency Injection Principle (DIP) and the Ports & Adapters (Hexagonal) pattern.


# Introduction
This is a streamlined E-commerce backend focusing on the core pillars of reliability: Orders and Products. Rather than using a heavy framework, this API leverages the Go standard library and a curated selection of "low-magic" tools to ensure the codebase is maintainable, testable, and lightning-fast.

Key Technical Highlights:
- Zero Floating-Point Errors: All currency is handled in integers (cents) to ensure 100% financial accuracy.
- Context-Driven Lifecycle: Every request is bounded by context.Context to prevent resource leaks and handle timeouts gracefully.
- Compile-Time Type Safety: Uses SQLC to generate Go code from raw SQL, eliminating runtime mapping errors.


## Demo & API Endpoints

Insert gif or link to demo


## Tech Stack

- Language: Go (Golang)
- Router: Chi (Context-aware & Lightweight)
- Database: PostgreSQL with pgx/v5
- Migrations: Goose
- SQL Generator: SQLC
- Structured Logging: slog (Standard Library)


## Run Locally

Clone the project

```bash
  git clone https://github.com/JeffreyOmoakah/E-commerce-backend-API.git
```

Go to the project directory

```bash
  cd E-commerce-backend-API
```

Install dependencies
This command downloads all the necessary packages (Chi, pgx, SQLC libs) and ensures your go.mod is in sync.
```bash
  go mod tidy
```

Start the server

```bash
  air
```


## Resources 

- sqlc GitHub Repository: https://github.com/sqlc-dev/sqlc on GitHub
- clean architecture: https://medium.com/@omidahn/clean-architecture-in-go-golang-a-comprehensive-guide-f8e422b7bfae

