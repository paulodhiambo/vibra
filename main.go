package main

import (
	"fmt"
	"os"
	"os/user"
	"vibra/repl"
)

func main() {
	osUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Vibra programming language!\n",
		osUser.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
