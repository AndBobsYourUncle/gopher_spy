package handlers

import (
	"bytes"
	"context"
	apiv1 "gopher_spy/protos/gen/go/detector/api/v1"
	"image"
	"log"
)

type grpcImpl struct{}

type Config struct{}

func NewGRPC(cfg *Config) (*grpcImpl, error) {
	return &grpcImpl{}, nil
}

func (s *grpcImpl) DetectFrame(
	ctx context.Context,
	req *apiv1.DetectFrameRequest,
) (*apiv1.DetectFrameResponse, error) {
	// create image from request bytes
	img, _, err := image.Decode(bytes.NewReader(req.GetFrame()))
	if err != nil {
		return nil, err
	}

	// detect objects in image
	detections, detErr := detect_objects_on_image(img)
	if detErr != nil {
		log.Printf("error detecting objects: %v", detErr)
	}

	retDetections := make([]*apiv1.Detection, len(detections))

	for _, detection := range detections {
		log.Printf("detected: %+v", detection)

		retDetections = append(retDetections, &apiv1.Detection{
			X1:         float32(detection.X1),
			Y1:         float32(detection.Y1),
			X2:         float32(detection.X2),
			Y2:         float32(detection.Y2),
			Label:      detection.Label,
			Confidence: detection.Probability,
		})
	}

	return &apiv1.DetectFrameResponse{
		Detections: retDetections,
	}, nil
}

// Function receives an image,
// passes it through YOLOv8 neural network
// and returns an array of detected objects
// and their bounding boxes
// Returns Array of bounding boxes in format [[x1,y1,x2,y2,object_type,probability],..]
func detect_objects_on_image(img image.Image) ([]*Detection, error) {
	input, img_width, img_height := prepare_input(img)
	output, err := run_model(input)
	if err != nil {
		return nil, err
	}

	data := process_output(output, img_width, img_height)

	return data, nil
}

// Function used to pass provided input tensor to
// YOLOv8 neural network and return result
// Returns raw output of YOLOv8 network as a single dimension
// array
func run_model(input []float32) ([]float32, error) {

	var err error

	if Yolo8Model.Session == nil {
		Yolo8Model, err = InitYolo8Session(input)
		if err != nil {
			return nil, err
		}
	}

	return runInference(Yolo8Model, input)

}
