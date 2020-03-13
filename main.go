package main

import (
	"fmt"
	"image/color"
	"log"

	"gocv.io/x/gocv"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatalf("error opening web cam: %v", err)
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	window := gocv.NewWindow("webcamwindow")
	defer window.Close()

	faceCascade := "haarcascades/haarcascade_frontalface_default.xml"
	eyeCascade := "haarcascades/haarcascade_eye.xml"

	classifierFace := gocv.NewCascadeClassifier()
	classifierFace.Load(faceCascade)
	defer classifierFace.Close()

	classifierEye := gocv.NewCascadeClassifier()
	classifierEye.Load(eyeCascade)
	defer classifierEye.Close()

	color := color.RGBA{0, 255, 0, 0}
	for true {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Println("Unable to read from the device")
			continue
		}

		faces := classifierFace.DetectMultiScale(img)
		for _, f := range faces {
			fmt.Println("face detected", f)
			gocv.Rectangle(&img, f, color, 3)

			eyes := classifierEye.DetectMultiScale(img)
			for _, e := range eyes {
				fmt.Println("eye detected", e)
				gocv.Rectangle(&img, e, color, 3)
			}
		}

		window.IMShow(img)
		window.WaitKey(50)
	}
}
