package gopsnr

import (
	"fmt"
	"math"

	"gocv.io/x/gocv"
)

const MAX float64 = 255

// Exec MSE/PSNR
func Exec(firstImg gocv.Mat, secondImg gocv.Mat) (mse float64, psnr float64, err error) {
	if firstImg.Rows() != secondImg.Rows() || firstImg.Cols() != secondImg.Cols() {
		return 0, 0, fmt.Errorf(
			"firstImg [rows: %v, cols: %v] secoundImg [rows: %v, cols: %v]",
			firstImg.Rows(), firstImg.Cols(), firstImg.Cols(), secondImg.Cols(),
		)
	}

	var mseSum float64
	for imgRows := 0; imgRows < firstImg.Rows(); imgRows++ {
		for imgCols := 0; imgCols < firstImg.Cols(); imgCols++ {
			distance := float64(firstImg.GetUCharAt(imgRows, imgCols)) - float64(secondImg.GetUCharAt(imgRows, imgCols))
			mseSum += math.Pow(distance, 2)
		}
	}

	temp := float64(firstImg.Rows() * firstImg.Cols())
	mse = mseSum / temp
	psnr = (20 * math.Log10(MAX)) - (10 * math.Log10(mse))

	return mse, psnr, nil
}
