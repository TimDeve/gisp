package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/TimDeve/gisp/eval"
)

func main() {
	if len(os.Args) < 2 {
		repl()
	} else {
		runFile()
	}
}

func repl() {
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

func runFile() {
	path := os.Args[1]

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := eval.Eval(string(bytes))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(result)
	}
}
