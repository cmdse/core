package main

import (
	"fmt"
	"github.com/cmdse/core/argparse"
)

func main() {
	tokens := argparse.ParseArguments([]string{"-l", "-p", "--only", "argument"}, nil)
	for _, token := range tokens {
		fmt.Println(token)
	}
}
