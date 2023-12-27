package handlers

import (
	"bytes"
	"context"
	"errors"
	"gopher_spy/detectors/yolov8/yolov8_processor"
	apiv1 "gopher_spy/protos/gen/go/detector/api/v1"
	"image"
)

type grpcImpl struct {
	modelProcessor yolov8_processor.Processor
}

type Config struct {
	ModelProcessor yolov8_processor.Processor
}

func NewGRPC(cfg *Config) (*grpcImpl, error) {
	if cfg == nil {
		return nil, errors.New("missing parameter: cfg")
	}

	if cfg.ModelProcessor == nil {
		return nil, errors.New("missing parameter: cfg.ModelProcessor")
	}

	return &grpcImpl{
		modelProcessor: cfg.ModelProcessor,
	}, nil
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

	// process image
	detections, err := s.modelProcessor.ProcessFrame(img)
	if err != nil {
		return nil, err
	}

	// create response
	resp := &apiv1.DetectFrameResponse{
		Detections: make([]*apiv1.Detection, len(detections)),
	}

	for i, detection := range detections {
		resp.Detections[i] = &apiv1.Detection{
			Label:      detection.Label,
			X1:         float32(detection.X1),
			Y1:         float32(detection.Y1),
			X2:         float32(detection.X2),
			Y2:         float32(detection.Y2),
			Confidence: detection.Probability,
		}
	}

	return resp, nil
}
