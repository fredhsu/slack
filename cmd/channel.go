// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

// channelCmd represents the channel command
var channelCmd = &cobra.Command{
	Use:   "channel",
	Short: "Actions for slack channel",
	Long: `Sending things to a slack channel.  Options include a message, or file.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("channel called")
		cmd.HasFlags()
		if cmd.Flag("file").Value.String() != "" {
			fmt.Printf("file is: %+v\n", cmd.Flag("file").Value)
			fileUpload(args[0], cmd.Flag("file").Value.String())
		}
		fmt.Printf("Args are %+v\n", args)
	},
}

//Upload the file to slack channel
func fileUpload(channel, filename string) {
	token := os.Getenv("SLACK_API_TOKEN")
	api := slack.New(token, slack.OptionDebug(false))

	// Join the channel if not already joined
	_, err := api.JoinChannel(channel)
	if err != nil {
		log.Printf("Error joining channel %s\n", err)
		return
	}
	fup := slack.FileUploadParameters{
		File: filename,
		// Filetype:        "text",
		Title:          "title",
		InitialComment: "comment",
		Channels:       []string{channel},
		// ThreadTimestamp: "time",
	}
	file, err := api.UploadFile(fup)
	if err != nil {
		log.Printf("Error uploading file: %s\n", err)
		return
	}
	log.Printf("Uploaded file %s to %+v", file.Name, file.Channels)

}

func init() {
	sendCmd.AddCommand(channelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	channelCmd.PersistentFlags().String("file", "", "File to upload")
	channelCmd.PersistentFlags().String("message", "", "Message to send")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// channelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
