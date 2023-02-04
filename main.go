package main

import (
	"fmt"
	"math"

	"github.com/montanaflynn/stats"
	"gocv.io/x/gocv"
)

func main() {
	img := gocv.IMRead("blur.jpg", gocv.IMReadColor)

	gray := gocv.NewMat()

	gocv.CvtColor(img, &gray, gocv.ColorBGRToGray)
	img.Close()

	laplacian := gocv.NewMat()

	gocv.Laplacian(gray, &laplacian, gocv.MatTypeCV64F, 1, 1, 0, gocv.BorderDefault)
	gray.Close()

	l2, err := laplacian.DataPtrFloat64()
	if err != nil {
		panic(err)
	}

	v2, err := stats.Variance(stats.Float64Data(l2))
	if err != nil {
		panic(err)
	}

	meanStdDev := gocv.NewMat()

	stdDev := gocv.NewMat()

	gocv.MeanStdDev(laplacian, &meanStdDev, &stdDev)
	laplacian.Close()
	meanStdDev.Close()

	sd, err := stdDev.DataPtrFloat64()
	if err != nil {
		panic(err)
	}

	stdDev.Close()

	variance := math.Pow(sd[0], 2)

	fmt.Println("Variance:", variance)

	fmt.Println("Stats variance:", v2)
}
