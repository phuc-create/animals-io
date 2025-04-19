# Welcome to Animals Social App

## Why?
Animals Social App was created by a Dog lover who is passionate about animals. My childhood has created such a phenomenal sense about loving dogs and other kinds. So let's explore with me!

# Stacks
- Golang
- NextJS
- PostgresQL
- Docker
- Tanstack Query
- Bun


# Backend Golang folder structure

```graphql
animals-io/api/
├── cmd/
│   └── your-app/         # Entry point (main.go)
├── internal/
│   ├── api/              # API routes and handlers
│   │   └── v1/           # Versioned API
│   │       ├── handlers/ # HTTP handlers per resource (e.g., user, product)
│   │       └── routes.go # Route registration using chi.Router
│   ├── config/           # App configuration
│   ├── models/           # Structs and database models
│   ├── repository/       # DB interaction (repositories)
│   ├── service/          # Business logic layer
│   └── middleware/       # Custom middleware
├── pkg/                  # Shared utilities (e.g., logging, helpers)
├── go.mod
├── go.sum
└── README.md

```
