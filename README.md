# fetch-rewards-receipt-processor-challenge

This Repository consists of the codebase for the take home assesment by Fetch Rewards.

# Assumptions Made
1. One of the rules stated to "round up" to nearest intger, so if the result is "1.05" it is rounded up as "2".
2. A sanity check is done to reject requests that have an invalid purchase date and time. Also, the dollar amount is strictly evaluated to have two digits to the right of decimals. This decision is made from the OpenAPI spec provided.

# Getting Started

## Environment Variables
Make sure the `.env` file is present. It is important for the API to start, as it has the PORT number to which it has to listen.

If `.env` file is not present, copy it from the template `.env.example` by running:
```
cp .env.example .env
```

### Example
`.env`
```
PORT=3001
```
## Installation

## Using Docker and Docker Compose (Recommended)
1. Install [Docker](https://docs.docker.com/engine/install/) and [Docker Compose](https://docs.docker.com/compose/install/).
2. Build Docker image 
From the Project root run:

```
make docker_build
```
or
```
docker-compose build
```
3. Run Docker container
```
make run_docker
```
or
```
docker-compose up -d
```

4. Stopping the container
```
make stop_docker
```
or
```
docker-compose down
```

## Direct Install (Alternative)
## Install Golang

Make sure you have Go 1.19 or higher installed.

https://golang.org/doc/install

## Installing Dependences
From the project root directory run:

```
make dep
```

or

```
go mod download
```

## Building and Running the Project

```
make build_and_run
```
or
```
go build -o fetch-rewards-api main.go
./fetch-rewards-api

```

## Testing
From project root directory run:

```
make test
```

or

```
go test -v
```