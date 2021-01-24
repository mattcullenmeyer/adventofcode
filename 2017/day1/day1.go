package main

import (
	"fmt"
	"log"
	"strconv"
)

func captcha(txt string) int {
	// sum of all digits that match the next digit
	// the list is circular, so digit after last digit is first digit
	// input is string and not int due to size limitations
	sum := 0
	prev := txt[len(txt)-1:] // last digit
	var current string
	// iterate through each digit, starting with first
	for _, c := range txt {
		current = string(c)
		// compare current digit to previous
		if prev == current {
			i, err := strconv.Atoi(current)
			if err != nil {
				log.Fatal(err)
			}
			sum += i
		}
		prev = current
	}
	return sum
}

func main() {
	fmt.Println(captcha("1123"))
}
