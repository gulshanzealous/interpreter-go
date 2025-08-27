package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/gulshanzealous/interpreter-go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s, welcome to basix programming language", user.Username)
	fmt.Printf("feel free to type commands\n")

	repl.Start(os.Stdin, os.Stdout)

}
