# go-envy

> A lightweight Go package to load environment-specific `.env` files like `.env.dev`, `.env.prod`, etc.

---

## Why go-envy?

When working with Go projects, managing different environment variables for development, production, staging, and testing can get messy.  
Instead of juggling a single `.env` file and commenting/uncommenting variables, **go-envy** lets you load environment variables based on the environment you specify — cleanly and effortlessly.

---

## Features

- Load `.env.{env}` files dynamically (`.env.dev`, `.env.prod`, `.env.test`, etc.)
- Simple API with just one function: `Load(env string)`
- Uses [joho/godotenv](https://github.com/joho/godotenv) under the hood
- Robust error handling with descriptive error messages
- Input validation for environment names
- Easily extendable for custom workflows

---

## Installation

```bash
go get github.com/ceosss/go-envy
```
---

## Usage

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/ceosss/go-envy/envy"
)

func main() {
    // Load environment variables from .env.dev
    if err := envy.Load("dev"); err != nil {
        log.Fatalf("Failed to load environment: %v", err)
    }

    // Access env vars as usual
    fmt.Println("Database Host:", os.Getenv("DB_HOST"))
}
```

## Error Handling

The `Load` function returns an error in the following cases:

- Empty environment name
- Environment name contains invalid characters
- Environment file cannot be loaded or doesn't exist

Example error handling:

```go
if err := envy.Load("dev"); err != nil {
    switch {
    case errors.Is(err, envy.ErrInvalidEnvName):
        // Handle invalid environment name
    default:
        // Handle other errors
    }
}