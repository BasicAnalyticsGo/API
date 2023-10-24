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
    curl -H "Content-Type: application/json" -X POST http://localhost:8080/sessions
    
    {"data":{"id":"","timestamp":"1698032908","IP":"127.0.0.1"}}
```

To get all sessions from database : 

Go to http://localhost:8080/sessions and you should have something like this :

```bash
   {"data":[{"id":"","timestamp":"1698032908","IP":"127.0.0.1"}]}
```