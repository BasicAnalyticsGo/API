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

## Deploy

### Create docker network

docker network create my_network

### Build and run

#### MySQL Database

```bash
docker build -t mysql_image -f Dockerfile.mysql .
docker run --name mysql_db -d -p 3306:3306 --network=my_network mysql_image
```

Then :

```bash
docker exec -it mysql_db /bin/bash
mysql -uroot -pmypassword
CREATE DATABASE analytics;
SHOW databases;
```

#### Golang API

```bash
docker build -t go_app_image -f Dockerfile.go .
docker run --name my_go_app -p 8080:8080 --network=my_network -e DB_CONNECTION_STRING="root:mypassword@tcp(mysql_db:3306)/analytics?charset=utf8mb4" go_app_image
```