package main

import (
    "fmt"
)

var t = true

func main() {
    f := false
    {
        i := 20
        fmt.Println(i)
    }
    fmt.Println(t, f)
}
