package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	check := 0
	args := os.Args[1:]
	number, err := strconv.Atoi(args[2])
	number0, err0 := strconv.Atoi(args[0])
	_ = number0
	if len(args) == 3 {
		fmt.Printf("\n")
		fmt.Println("Input:")
		for i := 0; i < 3; i++ {
			fmt.Printf("\"%s\" ", args[i])
		}
		fmt.Printf("\n\n")
		fmt.Println("Output:")
	} else {
		panic("Something went wrong!")

	}
	if args[1] == "+" {
		fmt.Printf("\"%s%s\"\n", args[0], args[2])
		check += 1
	}
	if args[1] == "-" {
		str := args[0]
		substr := args[2]
		if strings.Contains(str, substr) {
			modified := strings.Replace(str, substr, "", -1)
			fmt.Printf("\"%s\"\n", modified)
		} else {
			fmt.Printf("\"%s\"\n", str)
		}
		check += 1
	}
	if args[1] == "*" && err == nil && number > 0 && number < 10 && err0 != nil {
		str := ""
		for i := 0; i < number; i++ {
			str += args[0]
		}
		if len(str) > 40 {
			fmt.Printf("\"%s...\"\n", str[:40])

		} else {
			fmt.Printf("\"%s\"\n", str)
		}
		check += 1
	}
	if args[1] == "/" && err == nil && number > 0 && number < 10 && err0 != nil {
		fmt.Printf("\"%s\"\n", args[0][:(len(args[0])/number)])
		check += 1
	}
	if check == 0 {
		panic("Something went wrong!")
	}
}
