package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func showMsg() {
	fmt.Println("Welcome to mycli")
}

func main() {
	app := &cli.App{
		Name:  "mycli",
		Usage: "我的指令程式",
		Action: func(c *cli.Context) error {
			showMsg()
			for i := 0; i < c.NArg(); i++ {
				fmt.Printf("%d: %s\n", i+1, c.Args().Get(i))
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("=================================================================")
		log.Fatal(err)
	}
}
