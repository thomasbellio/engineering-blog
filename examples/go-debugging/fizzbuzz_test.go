package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	number := 15

	want := []string{
		"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz",
		"11",
		"Fizz",
		"13",
		"14",
		"FizzBuzz",
	}
	got := FizzBuzz(number)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %v, got %v ", want, got)
	}
}
