package main

import (
	"fmt"
	"os"

	"github.com/eaglewu/parser/repl"
)

func main() {
	fmt.Printf("This is the new PHP programming language!\n")
	repl.Start(os.Stdin, os.Stdout)
}
