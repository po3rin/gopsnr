package gopsnr_test

import (
	"testing"

	"github.com/po3rin/gopsnr"
	"gocv.io/x/gocv"
)

func TestExec(t *testing.T) {
	tests := []struct {
		name      string
		filepath1 string
		filepath2 string
		wantMSE   float64
		wantPSNR  float64
	}{
		{
			name:      "compare jpeg",
			filepath1: "img/gopher.jpeg",
			filepath2: "img/gopher_golden.jpeg",
			wantMSE:   0.04092627599243856,
			wantPSNR:  62.01078132213712,
		},
	}

	for _, tt := range tests {
		img1 := gocv.IMRead(tt.filepath1, gocv.IMReadColor)
		img2 := gocv.IMRead(tt.filepath2, gocv.IMReadColor)
		mse, psnr, err := gopsnr.Exec(img1, img2)
		if err != nil {
			t.Fatalf("got unexpected error: %v", err)
		}
		if mse != tt.wantMSE || psnr != tt.wantPSNR {
			t.Errorf(
				"unexpected result\n[MSE]want: %v, got: %v\n[PSNR]want: %v, got: %v\n",
				tt.wantMSE, mse, tt.wantPSNR, psnr,
			)
		}
	}
}
