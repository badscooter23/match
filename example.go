package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {

	args := os.Args
	numArgs := len(args)

	if numArgs != 3 {
		fmt.Println("expected: ./match regexp string")
		return
	}

	p := args[1]
	s := args[2]

	fmt.Println("Checking to see if ", p, "matches", s)

	fmt.Println("... compare using regex.MatchString()... ")
	match, e := regexp.MatchString(p, s)

	if e != nil {
		fmt.Println("error", e)
	} else {
		fmt.Println("result: ", match)
	}

	fmt.Println()

	fmt.Println("... compare using filepath.Match()... ")
	match, e = filepath.Match(p, s)
	if e != nil {
		fmt.Println("error", e)
	} else {
		fmt.Println("result: ", match)
	}

}
