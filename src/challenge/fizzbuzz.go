package challenge

import (
	"fmt"
	"strconv"
)

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		fmt.Println(internalFizzBuzz(i))
	}
}

func internalFizzBuzz(i int) string {
	if i%3 == 0 && i%5 == 0 {
		return "FizzBuzz"
	} else if i%3 == 0 {
		return "Fizz"
	} else if i%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(i)
	}
}
