package main

import (
    "fmt"
)

func main() {
    shadow := "world"
    fmt.Println(shadow)
    {
        shadow := "hello"
        fmt.Println(shadow)
    }
    
    fmt.Println(shadow)

}
