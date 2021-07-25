package main

import (
	"fmt"
	challenge "github.com/lixuemin/fizzbuzz"
	"github.com/myuser/calculator"
)

func main() {
	challenge.FizzBuzz()
}

func main2() {
	for i := 1; i <= 4; i++ {
		defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}
}

func main1() {
	total := calculator.Sum(3, 5)
	println(total)
	println("Version: ", calculator.Version)
	//total := calculator.internalSum(5)
	//println(total)
	//println("Version: ", calculator.logMessage)
}
