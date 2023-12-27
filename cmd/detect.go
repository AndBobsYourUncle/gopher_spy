package cmd

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopher_spy/detectors/yolov8/handlers"
	"gopher_spy/detectors/yolov8/yolov8_model"
	"gopher_spy/detectors/yolov8/yolov8_processor"
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

		yoloModel, err := yolov8_model.New(&yolov8_model.Config{})
		if err != nil {
			log.Fatalf("failed to create yolo yolov8_model: %v", err)
		}

		yoloProcessor, err := yolov8_processor.New(&yolov8_processor.Config{
			YoloModel: yoloModel,
		})
		if err != nil {
			log.Fatalf("failed to create yolo processor: %v", err)
		}

		handler, err := handlers.NewGRPC(&handlers.Config{
			ModelProcessor: yoloProcessor,
		})
		if err != nil {
			log.Fatalf("failed to create grpc handler: %v", err)
		}

		apiv1.RegisterDetectorAPIServer(s, handler)

		reflection.Register(s)

		log.Println("starting detection server on port :5005")

		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	rootCommand.AddCommand(detectCommand)
}
