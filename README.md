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

## Available POST endpoints

```
http://165.22.91.151:8181/issuing-jwt-token
```

Еxpect: Json as post: {"username":"", "password":""}


Example Request:
```
POST /issuing-jwt-token HTTP/1.1
Host: 165.22.91.151:8181
Content-Type: application/json
Cache-Control: no-cache

{"username":"user1", "password":"password1"}
```

Example Response:
```
{
    "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzMDA3ODh9.K1MWq28W3cak4niCP9QLnJB-Qr8vzH6GAA7du7_ofTU"
}
```