package main

import (
	"fmt"
	"os"

	"github.com/rishavmehra/solana-vanity/pkg"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "Wallet Generator",
		Usage: "Generate Solana wallets with a specified prefix",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "prefix",
				Value:    "100", // default value
				Usage:    "Target prefix for the wallet",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			searchString := c.String("prefix")
			fmt.Printf("Target Prefix: %s\n", searchString)
			for i := 0; i < 16; i++ {
				go pkg.GenerateWallet(searchString)
			}
			fmt.Scanln()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
