package main

import "fmt"

func FizzBuzz(number int) (r []string) {
	for i := 1; i <= number; i++ {
		var val string
		if i%3 == 0 {
			val += "Fizz"
		}
		if i%5 == 0 {
			val += "Buzz"
		}
		if val == "" {
			val = fmt.Sprint(i)
		}
		r = append(r, val)
	}

	return r
}

func main() {
	var fizzBuzz []string = FizzBuzz(15)
	fmt.Printf("FizzBuzz: %v", fizzBuzz)
}
