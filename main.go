package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/TimDeve/gisp/eval"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("=> ")
		scanner.Scan()
		text := scanner.Text()
		if text == "quit" || text == "exit" {
			break
		}

		result, err := eval.Eval(text)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}
