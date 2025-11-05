package main

import (
    "fmt"
    "path"
)

func main() {
    // integer
    x := 255
    var y int = 44
    fmt.Println(x, y)
    
    // float
    a1 := 25.45
    a2 := 34.24
    fmt.Println(a1 + a2)
    
    // complex
    var b1 complex128 = complex(6, 2)
    var b2 complex64 = complex(9, 2)
    fmt.Println(b1)
    fmt.Println(b2)
    
    // path ishlatish
    result := path.Base("\\Users\\user\\go\\src\\salom")
    fmt.Println(result)
}