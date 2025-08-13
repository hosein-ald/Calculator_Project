Calculator\_Project

Stateless HTTP API in Go (using Chi) that exposes four POST endpoints for basic arithmetic: **add**, **subtract**, **multiply**, and **divide**.
Project layout follows a simple, production-friendly structure with a `cmd/api` entrypoint and route/handler separation. ([GitHub][1])

---

## Features

* POST JSON API for `add`, `subtract`, `multiply`, `divide`
* Consistent request/response schema
* Sensible error handling (bad JSON, division by zero)
* Extensible folder layout (`cmd/`, `routes/`) for future middleware, auth, logging, etc. ([GitHub][1])

---

## Project structure

```
Calculator_Project/
├─ cmd/
│  └─ api/            # main entrypoint (starts HTTP server)
├─ routes/            # route registrations and handlers
├─ go.mod
└─ go.sum
```

(The public repo currently includes `cmd/api`, `routes`, and module files.) ([GitHub][1])

---

## Requirements

* Go 1.21+ (any recent Go version that supports modules)
* Modules fetched automatically on first build (`go mod tidy` if needed)

---

## Running locally

From the repository root:

```bash
go run ./cmd/api
```

The server starts on:

```
http://localhost:9000
```

---

## API

### Common request body

All endpoints accept the same JSON body:

```json
{
  "a": 2,
  "b": 8
}
```

> Keys must be **lowercase** (`"a"`, `"b"`) and wrapped in quotes to be valid JSON.

### Common response (success)

```json
{
  "result": 10
}
```

### Endpoints

#### POST `/add`

* **Description:** Returns `a + b`.
* **200** with `{ "result": number }`
* **400** if JSON is invalid.

**Example**

```bash
curl -s -X POST http://localhost:9000/add \
  -H "Content-Type: application/json" \
  -d '{"a": 5, "b": 3}'
# => {"result":8}
```

#### POST `/subtract`

* **Description:** Returns `a - b`.
* **200** with `{ "result": number }`
* **400** if JSON is invalid.

**Example**

```bash
curl -s -X POST http://localhost:9000/subtract \
  -H "Content-Type: application/json" \
  -d '{"a": 5, "b": 3}'
# => {"result":2}
```

#### POST `/multiply`

* **Description:** Returns `a * b`.
* **200** with `{ "result": number }`
* **400** if JSON is invalid.

**Example**

```bash
curl -s -X POST http://localhost:9000/multiply \
  -H "Content-Type: application/json" \
  -d '{"a": 5, "b": 3}'
# => {"result":15}
```

#### POST `/divide`

* **Description:** Returns `a / b`.
* **200** with `{ "result": number }`
* **400** if JSON is invalid or `b == 0` (division by zero).

**Example**

```bash
curl -s -X POST http://localhost:9000/divide \
  -H "Content-Type: application/json" \
  -d '{"a": 12, "b": 3}'
# => {"result":4}
```

---

## Implementation notes

* **Routing:** Uses \[go-chi/chi] for lightweight, idiomatic routing and grouping.
* **Handlers:** One handler per operation; each decodes JSON into a `MathRequest { a, b }`, validates input, computes the result, and encodes a `MathResponse`.
* **Errors:** Invalid JSON → `400`. Division by zero → `400` with an explanatory message.
* **Statelessness:** No in-memory state required; every call is independent.

> The repo’s top level shows `cmd/api` for the server and `routes` for handlers/registrations, aligning with common Go “cmd/pkg” patterns. ([GitHub][1])

---

## Extending the service (ideas)

* **Validation:** Stricter schema (e.g., disallow `NaN`, `Inf`).
* **Middleware:** Request logging (to stdout or SQLite), request IDs, rate limiting, CORS.
* **Versioning:** Mount under `/v1` and keep a changelog.
* **Deployment:** Containerize with a minimal Dockerfile and health checks.

