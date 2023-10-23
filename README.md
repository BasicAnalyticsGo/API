# API

## Dev

### Download dependencies

```bash
    go get .
```

### Run

```bash
    go run .
```

### Build

```bash
    go build .
```

## Test

To add a session in database :

```bash
    curl -d '{"IP":"127.0.0.1"}' -H "Content-Type: application/json" -X POST http://localhost:8080/sessions
    
    {"data":{"id":"","timestamp":"1698032908","IP":"127.0.0.1"}}
```