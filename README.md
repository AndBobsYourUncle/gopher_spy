# GopherSpy

GopherSpy is a simple program that connects to an RTSP stream, runs the frames from the stream through the YOLOv8m model, and serves up a live MJPEG stream with the bounding boxes drawn on the frames.

## Setup

You'll need ffmpeg, Go, onnxruntime, and the YOLOv8m model. You'll also need to be able to compile Go with CGO enabled.

### Go, onnxruntime, ffmpeg (Mac)

```bash
brew install onnxruntime
brew install go
brew install ffmpeg
```

## Usage

Example:
```bash
STREAM_URL=rtsp://192.168.1.20:8554/unicast \
ONNX_RUNTIME_LIB=/opt/homebrew/Cellar/onnxruntime/1.16.3/lib/libonnxruntime.1.16.3.dylib \
 go run . stream_and_detect
```

This starts up both the detection and streaming servers. Then open a browser to http://localhost:8080 to view the live MJPEG stream with detected bounding boxes and labels.

## How it Works

The program uses a Go package called `gortsplib` to connect to the RTSP stream. It then uses ffmpeg to decode the frames from the stream into image frames. The frames are then run through the YOLOv8m model using onnxruntime.

The frames are sent via a gRPC call to a separate server that runs the model. This lets us run the model in a separate program and process, which gives us maximum flexibility to use whatever programming language and framework best works for the model being used.