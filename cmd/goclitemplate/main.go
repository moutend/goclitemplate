package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/moutend/goclitemplate"
)

func main() {
	err := Run(os.Args)

	if err != nil {
		log.New(os.Stderr, "error: ", 0).Fatal(err)
		os.Exit(1)
	}

	os.Exit(0)
}

func Run(args []string) (err error) {
	u, err := user.Current()
	if err != nil {
		return
	}
	fmt.Println(goclitemplate.Hello(u.Name))
	fmt.Printf("This is %s %s. (version-revision)\n", args[0], goclitemplate.Version())
	return
}
