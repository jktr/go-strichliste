# go-strichliste: Go Bindings for the Strichliste API

This library implements a REST client and v2 API bindings for
[hackerspace bootstrap's strichliste](https://github.com/strichliste/strichliste-backend),
which is a pretty neat tally sheet server.

Please note that this library is not stable yet.

Packages are

  * [strichliste](https://godoc.org/github.com/jktr/go-strichliste) — implements the REST client
  * [strichliste/schema](https://godoc.org/github.com/jktr/go-strichliste/schema) — contains the API schemata

All of the current API has been implemented, but test coverage is
currently nonexistant, so the library is probably horribly buggy.

There's a sister project to this library,
[strichliste-cli](https://github.com/jktr/strichliste-cli),
which implements a proper CLI-based application for strichliste
and is also very WIP.

## Example

```go
package main

import (
    "fmt"
    s "github.com/jktr/go-strichliste"
    ss "github.com/jktr/go-strichliste/schema"
    "math"
    "os"
    "strconv"
    "strings"
)

func main() {

    if len(os.Args) <= 2 {
        fmt.Println("usage: example USER AMOUNT [COMMENT...]")
        os.Exit(1)
    }

    amount, err := strconv.ParseFloat(os.Args[2], 64)
    if err != nil || math.Abs(amount) < 0.005 {
        fmt.Println("error: invalid AMOUNT")
        os.Exit(1)
    }

    delta := int(math.Round(amount * 100))
    comment := ""
    if len(os.Args) > 2 {
        comment = strings.Join(os.Args[3:], " ")
    }

    client := s.NewClient(s.WithEndpoint(s.DefaultEndpoint))

    user, _, err := client.User.GetByName(os.Args[1])
    if err != nil {
        // API-specific errors can be disambiguated like this
        if er, ok := err.(*ss.ErrorResponse); ok {
            if er.Class == ss.ErrorUserNotFound {
                fmt.Println("error: no such user")
                os.Exit(1)
            }
        } else {
            fmt.Printf("error: %s", err.Error())
            os.Exit(1)
        }
    }

    tx, _, err := client.Transaction.Context(user.ID).WithComment(comment).Delta(delta)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    fmt.Println("new balance:", float64(tx.Issuer.Balance)/100)
}
```

## License

    Copyright (C) 2019 Konrad Tegtmeier

    This library is free software: you can redistribute it and/or modify
    it under the terms of the Lesser GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This library is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    Lesser GNU General Public License for more details.

    You should have received a copy of the Lesser GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

## v1 API

If you're looking for the v1 API Bindings: There's a set of terribly
legacy bindings baked into the legacy `strichlist-cli` for the v1 API,
which can be found [here](https://git.cs.uni-paderborn.de/jktr/strichliste-cli).

## Acknowledgments

The structure of this library is heavily based on the one of
[Hetzner's hcloud-go](https://github.com/hetznercloud/hcloud-go)
library. Thanks for open sourcing it!
