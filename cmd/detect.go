package cmd

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopher_spy/detectors/yolov8m/handlers"
	apiv1 "gopher_spy/protos/gen/go/detector/api/v1"
	"log"
	"net"
)

var detectCommand = &cobra.Command{
	Use:   "detect",
	Short: "starts the detection server",

	Run: func(cmd *cobra.Command, args []string) {
		listener, err := net.Listen("tcp", ":5005")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()

		handler, err := handlers.NewGRPC(&handlers.Config{})
		if err != nil {
			log.Fatalf("failed to create grpc handler: %v", err)
		}

		apiv1.RegisterDetectorAPIServer(s, handler)

		reflection.Register(s)

		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	rootCommand.AddCommand(detectCommand)
}
