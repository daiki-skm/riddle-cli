package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Riddle"
	app.Usage = "Give a Riddle."
	app.Version = "0.0.1"

	var riddleNum string
	var ansNum string

	//app.Flags = []cli.Flag {
	//	&cli.StringFlag{
	//		Name: "riddle",
	//		Usage: "Give a riddle.",
	//		Destination: &riddleNum,
	//	},
	//	&cli.StringFlag{
	//		Name: "answer",
	//		Usage: "Give an answer.",
	//		Destination: &ansNum,
	//	},
	//}

	//app.Action = func(c *cli.Context) error {
	//	return nil
	//}

	app.Commands = []*cli.Command{
		{
			Name: "riddle",
			Aliases: []string{"r"},
			Usage: "riddle --num 1 (Choose a number from 1 to 9.)",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "num",
					Usage: "--num 1 (Choose a number from 1 to 9.)",
					Destination: &riddleNum,
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Printf("Problem: ")
				switch riddleNum {
				case "1":
					fmt.Println("Why was ６ afraid of 7?")
				case "2":
					fmt.Println("If you say it you break it. What is it?")
				case "3":
					fmt.Println("What letter of the alphabet has the most water?")
				case "4":
					fmt.Println("What flowers grow under your nose?")
				case "5":
					fmt.Println("What can you catch but not throw?")
				case "6":
					fmt.Println("What is a snake’s favorite subject?")
				case "7":
					fmt.Println("What state is round at both ends, and high in the middle?")
				case "8":
					fmt.Println("What starts with T, ends with T, and is full of T?")
				case "9":
					fmt.Println("I’m tall when I’m young and I’m short when I’m old. What am I?")
				default:
					fmt.Println("Please choose from 1 to 9.")
				}
				return nil
			},
		},
		{
			Name: "answer",
			Aliases: []string{"a"},
			Usage: "answer --num 1 (Choose a number from 1 to 9.)",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "num",
					Usage: "--num 1 (Choose a number from 1 to 9.)",
					Destination: &ansNum,
				},
			},
			Action: func(c *cli.Context) error {
				switch ansNum {
				case "1":
					fmt.Println("Answer: Because ７８９！")
					fmt.Println("Reason: 7 8(ate) 9")
				case "2":
					fmt.Println("Answer: Silence")
					fmt.Println("Reason: break silence")
				case "3":
					fmt.Println("Answer: C")
					fmt.Println("Reason: C()Sea")
				case "4":
					fmt.Println("Answer: Fire")
					fmt.Println("Reason: Fire grows as we eat (burn) things, but it goes out when water is poured on it.")
				case "5":
					fmt.Println("Answer: A cold")
					fmt.Println("Reason: We use 'Catch a cold', but we don't use 'Throw a cold'.")
				case "6":
					fmt.Println("Answer: HISStory!")
					fmt.Println("Reason: The hissing sound made by a snake is called 'hiss' in English. The word 'his' in 'history' and the sound of the snake 'hiss' make 'HISStory'.")
				case "7":
					fmt.Println("Answer: Ohio")
					fmt.Println("Reason: 'O''hi'(high)'o'")
				case "8":
					fmt.Println("Answer: Teapot")
					fmt.Println("Reason: 'Teapot' is spelled starting with a 'T' and ending with a 'T'.'T' and 'Tea'")
				case "9":
					fmt.Println("Answer: Candle")
					fmt.Println("Reason: The more you use a candle, the shorter it becomes.")
				default:
					fmt.Println("Please choose from 1 to 9.")
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
