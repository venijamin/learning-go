package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6700868689030-6710575887156-nBSdKZ0lVT5iPFFOIPtvI3y2")
	os.Setenv("CHANNEL_ID", "C06LLRJMWMU")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"test.txt"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			File:     fileArr[i],
			Channels: channelArr,
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Name %s, URL %s", file.Name, file.URL)

	}
}
