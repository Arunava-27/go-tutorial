package main
import (
    "fmt"
    "unicode/utf8"
    "errors"
)

// main is the entry point of the program. It demonstrates various Go language features such as:
// - Printing a string
// - Declaring and using variables
// - Performing integer division with error handling
// - Concatenating strings
// - Using constants
// - Using shorthand variable declarations
// - Declaring multiple variables in a single statement
// - Type conversion
// - Type inference
// - Calculating the length of a string in bytes and runes (Unicode characters)
func main(){
    printMe("Hello, World!")

    // Variables
    var x int = 5
    var y int = 7
    var sum int = x + y
    fmt.Println(sum)

    // Division
    var num1 int = 5
    var num2 int = 2
    var result, remainder int
    var err error
    result, remainder, err = intDivision(num1, num2)
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Printf("Result: %d, Remainder: %d\n", result, remainder)
    }

    // string
    var a string = "Hello"
    var b string = "World"
    fmt.Println(a + " " + b)

    // Constants
    const c int = 5
    fmt.Println(c)

    // short hand
    d := 5
    fmt.Println(d)

    // multiple variables
    e, f := 5, 7
    fmt.Println(e + f)

    // type conversion
    var g float64 = 5.7
    h := int(g)
    fmt.Println(h)

    // type inference
    i := 5
    j := 5.6
    fmt.Println(float64(i) + j)

    // length of string
    k := "Hello, World!γ"
    fmt.Println(len(k)) // to get the byte length of string (not the actual length)
    fmt.Println(utf8.RuneCountInString(k)) // to get the actual length of string (including unicode characters)
}

func printMe(printValue string){
    fmt.Println(printValue)
}

func intDivision(a int, b int) (int, int, error){
    var err error
    if b == 0 {
        err = errors.New("division by zero is not allowed")
        fmt.Println(err)
        return 0, 0, err
    }
    var result int = a / b
    var remainder int = a % b
    return result, remainder, err
}