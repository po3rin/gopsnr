package main

import (
	"log"
	"os"

	"github.com/po3rin/gopsnr"
)

func main() {
	println("compare:", os.Args[1], "&", os.Args[2])

	mse, psnr, err := gopsnr.ExecWithFileName(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	println("[MSE] value:", mse, "[PSNR] value:", psnr)
}
