package main
import "fmt"
func shippingCalc(x int, y int, z int) int {
    return x + y + z
}

func main() {
    fmt.Println("enter Height: 42")
    fmt.Println("Enter Width: 13")
    fmt.Println("Total: ", shippingCalc(42, 13, 20))
}
