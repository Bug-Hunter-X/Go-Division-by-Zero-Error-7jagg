```go
package main

import (

    "fmt"

)

func main() {
    var i int
    fmt.Println(i)
    i = 10
    fmt.Println(i)
    if i != 0 {
        fmt.Println(10 / i)
    } else {
        fmt.Println("Error: Cannot divide by zero")
    }
}```