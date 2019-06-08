package main

import (
	"log"
	"os"

	"github.com/po3rin/gopsnr"
	"gocv.io/x/gocv"
)

func main() {
	img1 := gocv.IMRead(os.Args[1], gocv.IMReadColor)
	img2 := gocv.IMRead(os.Args[2], gocv.IMReadColor)

	println("compare:", os.Args[1], "&", os.Args[2])

	mse, psnr, err := gopsnr.Exec(img1, img2)
	if err != nil {
		log.Fatal(err)
	}
	println("[MSE] value:", mse, "[PSNR] value:", psnr)
}
