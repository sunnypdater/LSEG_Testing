package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n int
    fmt.Scan(&n)
	fmt.Print(fizzBuzz(n))

}

func fizzBuzz(n int) []string {
    var answer []string
    for i := 1; i <= n; i++ {
        if i%3 == 0 && i%5 == 0 {
            answer = append(answer, "FizzBuzz")
        } else if i%3 == 0 {
            answer = append(answer, "Fizz")
        } else if i%5 == 0 {
            answer = append(answer, "Buzz")
        } else {
            answer = append(answer, strconv.Itoa(i))
        }
    }
    return answer
}