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

Еxpect: Json as post: 
```
{"username":"username1", "password":"password1"}
```

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
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzMDA3ODh9.K1MWq28W3cak4niCP9QLnJB-Qr8vzH6GAA7du7_ofTU"
}
```





```
http://165.22.91.151:8181/validate-jwt-token
```

Еxpect: Json as post: 
```
{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUyNzQ3Mzl9.JqFmC69fqKSHS3TvYJVnoUo3Miba2limvF5UDW50XUM"}
```

Example Request:
```
POST /validate-jwt-token HTTP/1.1
Host: 165.22.91.151:8181
Content-Type: application/json
Cache-Control: no-cache

{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUyNzQ3Mzl9.JqFmC69fqKSHS3TvYJVnoUo3Miba2limvF5UDW50XUM"
}
```

Example Response:
```
{
    "Result": "Token is valid for user: user1!"
}
```




```
http://165.22.91.151:8181/get-functionalities
```

Еxpect: Json as post: 
```
{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUyNzQ3Mzl9.JqFmC69fqKSHS3TvYJVnoUo3Miba2limvF5UDW50XUM"}
```

Example Request:
```
POST /get-functionalities HTTP/1.1
Host: 165.22.91.151:8181
Content-Type: application/json
Cache-Control: no-cache

{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzMDA3ODh9.K1MWq28W3cak4niCP9QLnJB-Qr8vzH6GAA7du7_ofTU"
}
```

Example Response:
```
{
    "functions": {
        "GetIssuingJwtToken": {
            "url": "/issuing-jwt-token",
            "request-type": "POST",
            "accept-post-json": {
                "username": "user1",
                "password": "password1"
            },
            "response": "json"
        },
        "GetValidateJwtToken": {
            "url": "/validate-jwt-token",
            "request-type": "POST",
            "accept-post-json": {
                "token": ""
            },
            "response": "json"
        },
        "GetRepoInformation": {
            "url": "/get-repo-information",
            "request-type": "POST",
            "accept-post-json": {
                "token": "",
                "repo": ""
            },
            "response": "json"
        },
        "GetCommitInformation": {
            "url": "/get-commit-information",
            "request-type": "POST",
            "accept-post-json": {
                "token": "",
                "repo": "",
                "commit": ""
            },
            "response": "json"
        }
    }
}
```





```
http://165.22.91.151:8181/get-commit-information
```

Еxpect: Json as post: 
```
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUyOTA4OTF9.mGtY6cMEh7GnRrSSsPEdTkL2lBo2nhWM6_00IUUbBFs", "repo":"admiral", "commit":"553aa1e036f04c414a84ad234586a35dd5845db1"}
```

Example Request:
```
POST /get-commit-information HTTP/1.1
Host: 165.22.91.151:8181
Content-Type: application/json
Cache-Control: no-cache

{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUyOTA4OTF9.mGtY6cMEh7GnRrSSsPEdTkL2lBo2nhWM6_00IUUbBFs", "repo":"admiral", "commit":"553aa1e036f04c414a84ad234586a35dd5845db1"}
```

Example Response:
```
{
    "commit": {
        "author": {
            "date": "2019-11-01T22:27:56Z",
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]"
        },
        "comment_count": 0,
        "committer": {
            "date": "2019-11-04T13:44:20Z",
            "email": "llazarov@vmware.com",
            "name": "Lazarin Lazarov"
        },
        "message": "Bump jackson.version from 2.9.8 to 2.10.0\n\nUpdates `jackson-dataformat-yaml` from 2.9.8 to 2.10.0\n- [Release notes](https://github.com/FasterXML/jackson-dataformats-text/releases)\n- [Commits](https://github.com/FasterXML/jackson-dataformats-text/compare/jackson-dataformats-text-2.9.8...jackson-dataformats-text-2.10.0)\n\nUpdates `jackson-annotations` from 2.9.8 to 2.10.0\n- [Release notes](https://github.com/FasterXML/jackson/releases)\n- [Commits](https://github.com/FasterXML/jackson/commits)\n\nUpdates `jackson-databind` from 2.9.8 to 2.10.0\n- [Release notes](https://github.com/FasterXML/jackson/releases)\n- [Commits](https://github.com/FasterXML/jackson/commits)\n\nUpdates `jackson-datatype-joda` from 2.9.8 to 2.10.0\n- [Release notes](https://github.com/FasterXML/jackson-datatype-joda/releases)\n- [Commits](https://github.com/FasterXML/jackson-datatype-joda/compare/jackson-datatype-joda-2.9.8...jackson-datatype-joda-2.10.0)\n\nSigned-off-by: dependabot[bot] <support@github.com>\nChange-Id: I55ce2c14876f99a288d5bccab3c8d2dd6d85e7ad\nReviewed-on: https://bellevue-ci-gerrit.eng.vmware.com:443/c/92239\nUpgrade-Verified: jenkins <jenkins@vmware.com>\nBellevue-Verified: jenkins <jenkins@vmware.com>\nPG-Verified: jenkins <jenkins@vmware.com>\nCS-Verified: jenkins <jenkins@vmware.com>\nReviewed-by: Miroslav Shipkovenski <mshipkovenski@vmware.com>",
        "tree": {
            "sha": "fb60b9a4b037d7cf85fbe560fbd4106573e99bc3",
            "url": "https://api.github.com/repos/vmware/admiral/git/trees/fb60b9a4b037d7cf85fbe560fbd4106573e99bc3"
        },
        "url": "https://api.github.com/repos/vmware/admiral/git/commits/553aa1e036f04c414a84ad234586a35dd5845db1",
        "verification": {
            "payload": null,
            "reason": "unsigned",
            "signature": null,
            "verified": false
        }
    }
}
```







```
http://165.22.91.151:8181/get-repo-information
```

Еxpect: Json as post: 
```
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUyOTA4OTF9.mGtY6cMEh7GnRrSSsPEdTkL2lBo2nhWM6_00IUUbBFs", "repo":"admiral"}
```

Example Request:
```
POST /get-repo-information HTTP/1.1
Host: 165.22.91.151:8181
Content-Type: application/json
Cache-Control: no-cache

{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUyOTA4OTF9.mGtY6cMEh7GnRrSSsPEdTkL2lBo2nhWM6_00IUUbBFs", "repo":"admiral"}
```

Example Response:
```
{
    "commit": {
        "author": {
            "date": "2019-11-01T22:27:56Z",
            "email": "49699333+dependabot[bot]@users.noreply.github.com",
            "name": "dependabot[bot]"
        },
        "comment_count": 0,
        "committer": {
            "date": "2019-11-04T13:44:20Z",
            "email": "llazarov@vmware.com",
            "name": "Lazarin Lazarov"
        },
        "message": "Bump jackson.version from 2.9.8 to 2.10.0\n\nUpdates `jackson-dataformat-yaml` from 2.9.8 to 2.10.0\n- [Release notes](https://github.com/FasterXML/jackson-dataformats-text/releases)\n- [Commits](https://github.com/FasterXML/jackson-dataformats-text/compare/jackson-dataformats-text-2.9.8...jackson-dataformats-text-2.10.0)\n\nUpdates `jackson-annotations` from 2.9.8 to 2.10.0\n- [Release notes](https://github.com/FasterXML/jackson/releases)\n- [Commits](https://github.com/FasterXML/jackson/commits)\n\nUpdates `jackson-databind` from 2.9.8 to 2.10.0\n- [Release notes](https://github.com/FasterXML/jackson/releases)\n- [Commits](https://github.com/FasterXML/jackson/commits)\n\nUpdates `jackson-datatype-joda` from 2.9.8 to 2.10.0\n- [Release notes](https://github.com/FasterXML/jackson-datatype-joda/releases)\n- [Commits](https://github.com/FasterXML/jackson-datatype-joda/compare/jackson-datatype-joda-2.9.8...jackson-datatype-joda-2.10.0)\n\nSigned-off-by: dependabot[bot] <support@github.com>\nChange-Id: I55ce2c14876f99a288d5bccab3c8d2dd6d85e7ad\nReviewed-on: https://bellevue-ci-gerrit.eng.vmware.com:443/c/92239\nUpgrade-Verified: jenkins <jenkins@vmware.com>\nBellevue-Verified: jenkins <jenkins@vmware.com>\nPG-Verified: jenkins <jenkins@vmware.com>\nCS-Verified: jenkins <jenkins@vmware.com>\nReviewed-by: Miroslav Shipkovenski <mshipkovenski@vmware.com>",
        "tree": {
            "sha": "fb60b9a4b037d7cf85fbe560fbd4106573e99bc3",
            "url": "https://api.github.com/repos/vmware/admiral/git/trees/fb60b9a4b037d7cf85fbe560fbd4106573e99bc3"
        },
        "url": "https://api.github.com/repos/vmware/admiral/git/commits/553aa1e036f04c414a84ad234586a35dd5845db1",
        "verification": {
            "payload": null,
            "reason": "unsigned",
            "signature": null,
            "verified": false
        }
    }
}
```