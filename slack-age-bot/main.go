package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"strconv"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6700868689030-6730362499360-1kVWSY6xUgwzFkMnNjtFtJGM")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A06LVVCEWAG-6707550367314-d84ac2c9ded4f5ffdab939d0a4cc842ff9f9b04da6e59e027fc32aef503349d7")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
				return // Exit the function early if there's an error
			}
			age := 2021 - yob
			// Use fmt.Sprintf to format the string without printing
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func printCommandEvents(events <-chan *slacker.CommandEvent) {
	for event := range events {
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
	}
}
