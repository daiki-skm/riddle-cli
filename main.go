package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "quiz"
	app.Usage = "Give a quiz."
	app.Version = "0.0.1"

	app.Commands = []*cli.Command{
		{
			Name:    "quiz",
			Aliases: []string{"q"},
			Usage:   "Give a quiz",
			Subcommands: []*cli.Command{
				// `go run main.go quiz 1 でactionするコマンド
				{
					Name:  "1",
					Usage: "quiz 1",
					Action: func(c *cli.Context) error {
						fmt.Println("1+1=?")
						return nil
					},
				},
				// `go run main.go quiz 2 でactionするコマンド
				{
					Name:  "2",
					Usage: "quiz 2",
					Action: func(c *cli.Context) error {
						fmt.Println("2+2=?")
						return nil
					},
				},
			},
		},
		{
			Name:    "MyAnswer",
			Aliases: []string{"m"},
			Usage:   "Write a your answer",
			Subcommands: []*cli.Command{
				// `go run main.go MyAnswer 1 --text `val` でactionするコマンド
				{
					Name:  "1",
					Usage: "MyAnswer 1",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "text",
							Usage: "true / false",
						},
					},
					Action: func(c *cli.Context) error {
						//fmt.Printf("%s\n", c.String("text"))
						if c.String("text") == "2" {
							fmt.Println("True!")
						} else {
							fmt.Println("False!")
						}
						return nil
					},
				},
				// `go run main.go MyAnswer 2 --text `val` でactionするコマンド
				{
					Name:  "2",
					Usage: "MyAnswer 2",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "text",
							Usage: "true / false",
						},
					},
					Action: func(c *cli.Context) error {
						//fmt.Printf("%s\n", c.String("text"))
						if c.String("text") == "4" {
							fmt.Println("True!")
						} else {
							fmt.Println("False!")
						}
						return nil
					},
				},
			},
		},
		{
			Name:    "answer",
			Aliases: []string{"a"},
			Usage:   "Give an answer",
			Subcommands: []*cli.Command{
				// `go run main.go answer 1 でactionするコマンド
				{
					Name:  "1",
					Usage: "answer 1",
					Action: func(c *cli.Context) error {
						fmt.Println("2")
						return nil
					},
				},
				// `go run main.go answer 2 でactionするコマンド
				{
					Name:  "2",
					Usage: "answer 2",
					Action: func(c *cli.Context) error {
						fmt.Println("4")
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
