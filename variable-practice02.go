package main
import (
    "fmt"
)

var zzz string = "Package level (global) variable\n"

func main() {

    fmt.Printf(zzz)

    var mug string = "filed with tea\n"
    fmt.Printf(mug)

    var a = "Hello"
    var b = 23
    var c = true
    var d = 2.3

    a = "Goodbye"
    //var a = "Goodbye"

    fmt.Printf("The type of a, b, c, d isr: %T, %T, %T, and %T respectively\n", a, b, c, d)


}
