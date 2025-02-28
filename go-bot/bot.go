package bot

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/slack-io/slacker"
)

func Run() {
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	bot.AddCommand(&slacker.CommandDefinition{
		Command:     "my yob is <year>",
		Description: "yob calculator",
		Examples:    []string{"my yob is 2020"},
		Handler: func(ctx *slacker.CommandContext) {
			year := ctx.Request().Param("year")
			yob, err := strconv.Atoi(year)

			printCommandEvents(ctx.Event())
			// check for only mention events for bot
			if ctx.Event().Type == "app_mention" {
				if yob > 2025 {
					ctx.Response().ReplyError(fmt.Errorf("doesn't look like a valid birth year"))
					return
				}
				if err != nil {
					ctx.Response().ReplyError(err)
					return
				}
				age := 2025 - yob

				r := fmt.Sprintf("age is %d", age)
				ctx.Response().Reply(r)
			}
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}

}

func printCommandEvents(event *slacker.MessageEvent) {
	// for event := range analyticsChannel {
	fmt.Println("\033[95m Command Events")
	fmt.Println("\033[92m [+] Timestamp : ", event.TimeStamp)
	fmt.Println("\033[92m [+] Data : ", event.Data)
	// fmt.Println("\033[92m [+] Parameters : ", event.Pa)
	fmt.Println("\033[92m [+] Event Channel Name : ", event.Channel.Name)
	fmt.Println("\033[92m [+] Event ChannelID : ", event.ChannelID)
	fmt.Println("\033[92m [+] Event Channel Text : ", event.Text)
	fmt.Println("\033[92m [+] Event Channel Type : ", event.Type)
}
