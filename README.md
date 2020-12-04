# force-go
A Go client for the Salesforce REST API.

## Install

```sh
go get github.com/davidji99/force-go
```

## Usage

### Sample:

```go
package main

import (
    "fmt"

    "github.com/davidji99/force-go/force"
)


func main() {
    client, err := force.New(force.OAuthCred(
        "OAUTH_USER",
        "OAUTH_PASSWORD",
        "OAUTH_CLIENT_ID",
        "OAUTH_CLIENT_SECRET",
    ))

    if err != nil {
        panic(err)
    }

    c, _, describeErr := client.Describe("Case")

    if describeErr != nil {
        panic(describeErr)
    }

    fmt.Println(c.GetName())
}
```

## Author

[davidji99](https://github.com/davidji99)
