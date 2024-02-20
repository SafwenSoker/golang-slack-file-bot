package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)


func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6695910382112-6658623966391-2HHCZ013YlSS4X637RuII9Iq")
	os.Setenv("CHANNEL_ID", "C06KLEPA5T8")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"products.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)

	}
}
