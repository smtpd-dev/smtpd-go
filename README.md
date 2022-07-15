# smtpd-go

[![Continuous Integration][1]][2]
[![GoDoc][3]][4]
[![Go Report Card][5]][6]

[1]: https://github.com/smtpd-dev/smtpd-go/actions/workflows/pipeline.yml/badge.svg
[2]: https://github.com/smtpd-dev/smtpd-go/actions/workflows/pipeline.yml
[3]: https://godoc.org/github.com/smtpd-dev/smtpd-go?status.svg
[4]: https://godoc.org/github.com/smtpd-dev/smtpd-go
[5]: https://goreportcard.com/badge/github.com/smtpd-dev/smtpd-go
[6]: https://goreportcard.com/report/github.com/smtpd-dev/smtpd-go

The official [SMTPD][smtpd.dev] Go client library.

## Supported versions

The current version is v0.1

Please use go modules and import via `github.com/smtpd-dev/go-smtpd/v0.1`.

## Installation

Make sure your project is using Go Modules (it will have a `go.mod` file in its
root if it already is):

```sh
go mod init
```

Then, reference stripe-go in a Go program with `import`:

```go
import (
    "github.com/smtpd-dev/smtpd-go/v0.1"
)
```

Run any of the normal `go` commands (`build`/`install`/`test`). The Go
toolchain will resolve and fetch the stripe-go module automatically.

Alternatively, you can also explicitly `go get` the package into a project:

```bash
go get -u github.com/smtpd-dev/smtpd-go/v0.1
```

## Usage

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/joho/godotenv"
    "github.com/smtpd/go-smtpd/v0.1"
    "log"
    "os"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    key := os.Getenv("API_KEY")
    secret := os.Getenv("API_SECRET")
    client := smtpd.New(key, secret)

    p, err := client.GetAllProfiles()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(PrettyJSON(p))
}

func PrettyJSON(input interface{}) string {
    b, err := json.Marshal(input)
    if err != nil {
        log.Fatal(err)
    }

    var prettyJSON bytes.Buffer
    err = json.Indent(&prettyJSON, b, "", "\t")
    if err != nil {
        log.Fatal(err)
    }
    return string(prettyJSON.Bytes())
}
```

## Issues

* None

## Development

Pull requests from the community are welcome. If you submit one, please keep
the following guidelines in mind:

1. Code must be `go fmt` compliant.
2. All types, structs and funcs should be documented.
3. Ensure that `make lint ** make test` succeeds.

## Test

Before running the tests, make sure to grab all of the package's dependencies:

    `go get -t -v`

Run all tests:

    `make test`

For any requests, bug or comments, please [open an issue][issues] or [submit a
pull request][pulls].
