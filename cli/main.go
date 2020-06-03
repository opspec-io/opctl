package main

import (
	"fmt"
	"os"

	"github.com/opctl/opctl/cli/internal/clicolorer"
	corePkg "github.com/opctl/opctl/cli/internal/core"
)

func main() {
	cliColorer := clicolorer.New()
	defer func() {
		if panicArg := recover(); panicArg != nil {
			fmt.Println(
				cliColorer.Error(
					fmt.Sprint(panicArg),
				),
			)
			os.Exit(1)
		}
	}()

	newCli(
		cliColorer,
		corePkg.New,
	).
		Run(os.Args)

}
