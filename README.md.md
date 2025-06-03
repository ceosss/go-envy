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
- Provides clear error messages if env files fail to load
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
    "os"

    "github.com/ceosss/go-envy/envy"
)

func main() {
    // Load environment variables from .env.dev
    envy.Load("dev")

    // Access env vars as usual
    fmt.Println("Database Host:", os.Getenv("DB_HOST"))
}