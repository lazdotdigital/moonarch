# moonarch

A dead simple Go wrapper around the hidden moonarch.app API.

# How-To

First, get the repository: `go get github.com/lazdotdigital/moonarch`. `moonarch` only exposes a single function: `Fetch`.

Here is an example of how it can be used:

```
package main

import (
    "fmt"

    "github.com/lazdotdigital/moonarch"
)

func main() {
    res, err := moonarch.Fetch("0x580de58c1bd593a43dadcf0a739d504621817c05")
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
}
```
