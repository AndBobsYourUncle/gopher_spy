package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCommand = &cobra.Command{
	Use: "gopher_spy",
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		log.Fatalf("error in rootCommand.Execute: %v", err)
	}
}
