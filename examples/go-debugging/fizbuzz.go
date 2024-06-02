package main

import "fmt"

func FizzBuzz(n int) (r []string) {
	for i := 1; i <= n; i++ {

		var val string = fmt.Sprint(i)
		if (i % 3) == 0 {
			val = "Fizz"
		}
		if (i % 5) == 0 {
			val = "Buzz"
		}
		if (i%3) == 0 && (i%5) == 0 {
			val = "FizzBuzz"
		}

		r = append(r, val)
	}

	return r
}

func main() {

}
