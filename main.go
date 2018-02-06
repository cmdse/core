package main

import (
	"cldse-cli/argparse"
	"fmt"
)

func main() {
	tokens := argparse.ParseArguments([]string{"-l", "-p", "--only", "argument"})
	for _, token := range tokens {
		fmt.Println(token)
	}
}
