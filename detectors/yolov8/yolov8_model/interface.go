package yolov8_model

type Model interface {
	ClassName(id int) string
	RunInference(input []float32) ([]float32, error)
}
