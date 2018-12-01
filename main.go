package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "meaningless"
	app.Usage = "🙄"
	app.UsageText = "🙄"
	app.Version = "v0.1.0"
	app.Action = func(c *cli.Context) error {
		fmt.Println("😀 < Hi. Do you like 🍣?")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
