# Github Api

Get commits and repository information from Github.

### Prerequisites

You need to have:

```
Docker version 19.03.4
Docker-compose version 1.21.2
Git
```

### Installing

Steps for starting development server:

```
git clone https://github.com/yankoyankov8/github-api.git
cd github-api/
docker-compose up -d
docker-compose exec api bash
go run main.go
```

## Running the tests

```
go test ./...
```
