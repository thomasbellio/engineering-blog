/*
*
This is the main package for FizzBuzz, this package is for demonstration purposes only
*
*/
package main

import "fmt"

// FizzBuzz Takes a number and returns string slice of size (number) containing:
// 1. The number represented as a string
// 2. The value 'Fizz' if the number is divisible by 3
// 3. The value 'Buzz' if the number is divisible by 5
// 4. The value 'FizzBuzz' if the number is divisible by 3 & 5
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
	var fizzBuzz = FizzBuzz(15)
	fmt.Printf("FizzBuzz: %v", fizzBuzz)
}
