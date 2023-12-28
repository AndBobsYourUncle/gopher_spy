package yolov8_model

import (
	ort "github.com/yalue/onnxruntime_go"
	"os"
)

const (
	modelPath = "./detectors/yolov8/yolov8m.onnx"
)

type modelSession struct {
	Session *ort.AdvancedSession
	Input   *ort.Tensor[float32]
	Output  *ort.Tensor[float32]
}

type modelImpl struct {
	session *modelSession
}

type Config struct {
	ModelPath string
}

func New(cfg *Config) (*modelImpl, error) {
	session, err := initYolo8Session()
	if err != nil {
		return nil, err
	}

	return &modelImpl{
		session: session,
	}, nil
}

func (m *modelImpl) ClassName(id int) string {
	if id < 0 || id >= len(yoloClasses) {
		return ""
	}

	return yoloClasses[id]
}

func (m *modelImpl) RunInference(input []float32) ([]float32, error) {
	inTensor := m.session.Input.GetData()
	copy(inTensor, input)
	err := m.session.Session.Run()
	if err != nil {
		return nil, err
	}
	return m.session.Output.GetData(), nil
}

// Array of YOLOv8 class labels
var yoloClasses = []string{
	"person", "bicycle", "car", "motorcycle", "airplane", "bus", "train", "truck", "boat",
	"traffic light", "fire hydrant", "stop sign", "parking meter", "bench", "bird", "cat", "dog", "horse",
	"sheep", "cow", "elephant", "bear", "zebra", "giraffe", "backpack", "umbrella", "handbag", "tie",
	"suitcase", "frisbee", "skis", "snowboard", "sports ball", "kite", "baseball bat", "baseball glove",
	"skateboard", "surfboard", "tennis racket", "bottle", "wine glass", "cup", "fork", "knife", "spoon",
	"bowl", "banana", "apple", "sandwich", "orange", "broccoli", "carrot", "hot dog", "pizza", "donut",
	"cake", "chair", "couch", "potted plant", "bed", "dining table", "toilet", "tv", "laptop", "mouse",
	"remote", "keyboard", "cell phone", "microwave", "oven", "toaster", "sink", "refrigerator", "book",
	"clock", "vase", "scissors", "teddy bear", "hair drier", "toothbrush",
}

func getSharedLibPath() string {
	return os.Getenv("ONNX_RUNTIME_LIB")
}

func initYolo8Session() (*modelSession, error) {
	ort.SetSharedLibraryPath(getSharedLibPath())
	err := ort.InitializeEnvironment()
	if err != nil {
		return &modelSession{}, err
	}

	input := make([]float32, 1*3*640*640)

	inputShape := ort.NewShape(1, 3, 640, 640)
	inputTensor, err := ort.NewTensor(inputShape, input)
	if err != nil {
		return &modelSession{}, err
	}

	outputShape := ort.NewShape(1, 84, 8400)
	outputTensor, err := ort.NewEmptyTensor[float32](outputShape)
	if err != nil {
		return &modelSession{}, err
	}

	options, e := ort.NewSessionOptions()
	if e != nil {
		return &modelSession{}, err
	}

	session, err := ort.NewAdvancedSession(modelPath,
		[]string{"images"}, []string{"output0"},
		[]ort.ArbitraryTensor{inputTensor}, []ort.ArbitraryTensor{outputTensor}, options)

	if err != nil {
		return &modelSession{}, err
	}

	modelSes := &modelSession{
		Session: session,
		Input:   inputTensor,
		Output:  outputTensor,
	}

	return modelSes, err
}
