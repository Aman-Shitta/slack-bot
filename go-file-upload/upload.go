package uploader

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func Run() {
	// file uploader
	os.Setenv("CHANNEL_ID", "C08EULSKDNK")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channels := []string{os.Getenv("CHANNEL_ID")}

	files := []string{"go.txt"}

	for i := 0; i < len(files); i++ {
		file := files[i]
		params := slack.UploadFileV2Parameters{
			Channel:  channels[0],
			File:     file,
			Filename: "go.txt",
			FileSize: 20,
		}

		summary, err := api.UploadFileV2(params)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("File Uploaded : ", summary)

	}
}
