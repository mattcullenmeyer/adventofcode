package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	passphrases := `
	aa bb cc dd ee
	aa bb cc dd aa
	aa bb cc dd aaa
	`
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(passphrases))
	// loop through all passphrases
	for scanner.Scan() {
		valid := true
		line := scanner.Text()
		// convert passphrase text to slice
		passphrase := strings.Fields(line)
		// check for empty passphrase
		if len(passphrase) == 0 {
			continue
		}
		// loop through passphrase to check for duplicates
		for i := 0; i < len(passphrase); i++ {
			for j := i + 1; j < len(passphrase); j++ {
				if passphrase[i] == passphrase[j] {
					valid = false
					break
				}
			}
		}
		if valid {
			count++
		}
	}
	fmt.Println(count)
}
