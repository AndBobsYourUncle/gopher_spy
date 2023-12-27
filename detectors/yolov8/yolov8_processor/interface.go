package yolov8_processor

import "image"

type Processor interface {
	ProcessFrame(img image.Image) ([]*Detection, error)
}
