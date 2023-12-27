package main

import (
	"github.com/bluenviron/gortsplib/v4"
	"github.com/bluenviron/gortsplib/v4/pkg/base"
	"github.com/bluenviron/gortsplib/v4/pkg/format"
	"github.com/bluenviron/gortsplib/v4/pkg/format/rtph264"
	"github.com/nsmith5/mjpeg"
	"github.com/pion/rtp"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

func main() {
	c := gortsplib.Client{}

	// parse URL
	u, err := base.ParseURL("rtsp://192.168.1.20:8554/unicast")
	if err != nil {
		log.Fatalf("Error parsing URL: %v", err)
	}

	// connect to the server
	err = c.Start(u.Scheme, u.Host)
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer c.Close()

	// find available medias
	desc, _, err := c.Describe(u)
	if err != nil {
		log.Fatalf("Error describing stream: %v", err)
	}

	// find the H264 media and format
	var forma *format.H264
	medi := desc.FindFormat(&forma)
	if medi == nil {
		log.Fatalf("H264 media not found")
	}

	// setup RTP/H264 -> H264 decoder
	rtpDec, err := forma.CreateDecoder()
	if err != nil {
		log.Fatalf("Error creating decoder: %v", err)
	}

	// setup H264 -> raw frames decoder
	frameDec := &h264Decoder{}
	frameDec.initialize()
	if err != nil {
		log.Fatalf("Error creating decoder: %v", err)
	}
	defer frameDec.close()

	// if SPS and PPS are present into the SDP, send them to the decoder
	if forma.SPS != nil {
		frameDec.decode(forma.SPS)
	}
	if forma.PPS != nil {
		frameDec.decode(forma.PPS)
	}

	// setup a single media
	_, err = c.Setup(desc.BaseURL, medi, 0, 0)
	if err != nil {
		log.Fatalf("Error setting up stream: %v", err)
	}

	imgFrames := make(chan image.Image)

	// called when a RTP packet arrives
	c.OnPacketRTP(medi, forma, func(pkt *rtp.Packet) {
		// decode timestamp
		_, ok := c.PacketPTS(medi, pkt)
		if !ok {
			//log.Printf("waiting for timestamp")
			return
		}

		// extract access units from RTP packets
		au, err := rtpDec.Decode(pkt)
		if err != nil {
			if err != rtph264.ErrNonStartingPacketAndNoPrevious && err != rtph264.ErrMorePacketsNeeded {
				log.Printf("ERR: %v", err)
			}
			return
		}

		for _, nalu := range au {
			// convert NALUs into RGBA frames
			img, err := frameDec.decode(nalu)
			if err != nil {
				log.Fatalf("Error decoding frame: %v", err)
			}

			// wait for a frame
			if img == nil {
				continue
			}

			imgFrames <- img

			//log.Printf("decoded frame with PTS %v and size %v", pts, img.Bounds().Max)
		}
	})

	// start playing
	_, err = c.Play(nil)
	if err != nil {
		log.Fatalf("Error opening stream: %v", err)
	}

	detectedImg := make(chan image.Image)

	processingImage := atomic.Bool{}

	lastTime := time.Now()
	fps := 0.0

	latestDetections := []*Detection{}

	go func() {
		for {
			img := <-detectedImg

			t1 := time.Now()

			detections, detErr := detect_objects_on_image(img)
			if detErr != nil {
				log.Printf("error detecting objects: %v", detErr)
			}

			for _, detection := range detections {
				log.Printf("detected: %+v", detection)
			}

			latestDetections = detections

			detectionFps := 1 / time.Since(t1).Seconds()

			log.Printf("stream fps: %.1f, detection fps: %.1f", fps, detectionFps)

			processingImage.Store(false)
		}
	}()

	stream := mjpeg.Handler{
		Next: func() (image.Image, error) {
			img := <-imgFrames

			fps = 1 / time.Since(lastTime).Seconds()
			lastTime = time.Now()

			if !processingImage.Load() {
				processingImage.Store(true)
				detectedImg <- img
			}

			// draw bounding boxes
			for _, detection := range latestDetections {
				// draw rectangle
				drawHollowRectangle(
					img.(*image.RGBA),
					int(detection.X1),
					int(detection.Y1),
					int(detection.X2-detection.X1),
					int(detection.Y2-detection.Y1),
					2,
				)

				// add label
				addLabel(
					img.(*image.RGBA),
					int(detection.X1),
					int(detection.Y1),
					detection.Label,
				)
			}

			return img, nil
		},
		Options: &jpeg.Options{Quality: 80},
	}

	mux := http.NewServeMux()
	mux.Handle("/stream", stream)

	log.Println(http.ListenAndServe(":8080", mux))

	c.Close()
}

func drawHollowRectangle(img *image.RGBA, x, y, width, height, borderWidth int) {
	// bright green color
	rectColor := color.RGBA{0, 255, 0, 255}

	// Draw top and bottom borders
	for i := 0; i < borderWidth; i++ {
		for j := 0; j < width; j++ {
			img.Set(x+j, y+i, rectColor)          // Top border
			img.Set(x+j, y+height-i-1, rectColor) // Bottom border
		}
	}

	// Draw left and right borders
	for i := borderWidth; i < height-borderWidth; i++ {
		img.Set(x, y+i, rectColor)         // Left border
		img.Set(x+width-1, y+i, rectColor) // Right border
	}
}

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{0, 255, 0, 255}
	point := fixed.Point26_6{fixed.I(x), fixed.I(y)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
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
