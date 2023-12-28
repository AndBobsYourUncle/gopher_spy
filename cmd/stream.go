package cmd

import (
	"github.com/spf13/cobra"
	"gopher_spy/streamer"
	"log"
)

var streamCommand = &cobra.Command{
	Use:   "stream",
	Short: "starts the streaming server",

	Run: func(cmd *cobra.Command, args []string) {
		stream, err := streamer.New(&streamer.Config{})
		if err != nil {
			log.Fatalf("failed to create streamer: %v", err)
		}

		stream.Stream()
	},
}

func init() {
	rootCommand.AddCommand(streamCommand)
}
