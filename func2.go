package main

import (
    "fmt"
    "errors"
)

func main() {
    var a string = "Hello Go!"
    printme(a)

    result, reminder, err := indivision(10, 2)
    if err != nil {
        fmt.Println(err.Error())
    } else if reminder != 0 {
        fmt.Printf("Division result of the integer is %v with reminder %v\n", result, reminder)
    } else {
        fmt.Printf("Division integer result is %v\n", result)
    }

    fmt.Println("Division: ", result)
}

func printme(printvalue string) {
    fmt.Println(printvalue)
    fmt.Println("Printed Successfully")
}

func indivision(num int, denom int) (int, int, error) {
    if denom == 0 {
        err := errors.New("Cannot divide by zero")
        return 0, 0, err
    }
    res := num / denom
    reminder := num % denom
    return res, reminder, nil
}
