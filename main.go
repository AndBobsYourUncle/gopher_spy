package main

import (
	"github.com/bluenviron/gortsplib/v4"
	"github.com/bluenviron/gortsplib/v4/pkg/base"
	"github.com/bluenviron/gortsplib/v4/pkg/format"
	"github.com/bluenviron/gortsplib/v4/pkg/format/rtph264"
	"github.com/nsmith5/mjpeg"
	"github.com/pion/rtp"
	"image"
	"image/jpeg"
	"log"
	"net/http"
)

// This example shows how to
// 1. connect to a RTSP server
// 2. read all media streams on a path.

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
		pts, ok := c.PacketPTS(medi, pkt)
		if !ok {
			log.Printf("waiting for timestamp")
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

			log.Printf("decoded frame with PTS %v and size %v", pts, img.Bounds().Max)
		}
	})

	// start playing
	_, err = c.Play(nil)
	if err != nil {
		log.Fatalf("Error opening stream: %v", err)
	}

	stream := mjpeg.Handler{
		Next: func() (image.Image, error) {
			img := <-imgFrames

			return img, nil
		},
		Options: &jpeg.Options{Quality: 80},
	}

	mux := http.NewServeMux()
	mux.Handle("/stream", stream)

	log.Println(http.ListenAndServe(":8080", mux))

	c.Close()
}
