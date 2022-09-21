# Upvote API using Go gRPC MongoDB

An API for Upvote Write in Go With MongoDB and gRPC technologies

## Usage

Edit `.env_example` and rename to `.env` passing the database connection info.
Then run `source .env` before start server to configure environment variables:
```sh
source .env
```

Install `go 1.18` and update dependencies:
```sh
# update dependencies: go 1.18+
go mod tidy
# 
```

**Important: This project needs docker and docker-compose to run local db**

Run `docker-compose up` to turn on service local db

Or just ignore to use a online db sevice.

Run server

```sh
go run server/main.go server/server.go
# **A message like this should appear:**
# 2022/08/07 10:33:32 Starting server, listen to port 50051
```

Run client

```sh
go run client/main.go client/client.go
# **A message like this should appear:**
# Starting Client...
# 2022/08/08 10:34:08 Sending Upvote through gRPC...
# 2022/08/08 10:34:09 Response from Upvote: email:"valid@email.com"  score:10
```

---

## Todo:

- [X] MongoDB connection

- [X] Protobuf definition

- [X] gRPC connection
  - [X] server definition
  - [X] cliente definition

- [ ] Deployment

- [X] Tests

- [X] Documentation
