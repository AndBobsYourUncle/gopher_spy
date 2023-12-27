package yolov8_processor

import (
	"errors"
	"gopher_spy/detectors/yolov8m/yolov8_model"
	"image"
)

type processorImpl struct {
	yoloModel yolov8_model.Model
}

type Config struct {
	YoloModel yolov8_model.Model
}

func New(cfg *Config) (*processorImpl, error) {
	if cfg == nil {
		return nil, errors.New("missing parameter: cfg")
	}

	if cfg.YoloModel == nil {
		return nil, errors.New("missing parameter: cfg.YoloModel")
	}

	return &processorImpl{
		yoloModel: cfg.YoloModel,
	}, nil
}

func (p *processorImpl) ProcessFrame(img image.Image) ([]*Detection, error) {
	input, img_width, img_height := prepare_input(img)
	output, err := p.yoloModel.RunInference(input)
	if err != nil {
		return nil, err
	}

	data := p.process_output(output, img_width, img_height)

	return data, nil
}
