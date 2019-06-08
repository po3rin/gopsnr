package gopsnr_test

import (
	"testing"

	"github.com/po3rin/gopsnr"
	"gocv.io/x/gocv"
)

func TestExecWithMat(t *testing.T) {
	tests := []struct {
		name     string
		mat1     gocv.Mat
		mat2     gocv.Mat
		wantMSE  float64
		wantPSNR float64
	}{
		{
			name:     "compare jpeg",
			mat1:     gocv.IMRead("img/gopher.jpeg", gocv.IMReadColor),
			mat2:     gocv.IMRead("img/gopher_golden.jpeg", gocv.IMReadColor),
			wantMSE:  0.04092627599243856,
			wantPSNR: 62.01078132213712,
		},
	}

	for _, tt := range tests {
		mse, psnr, err := gopsnr.ExecWithMat(tt.mat1, tt.mat2)
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

func TestExecWithFileName(t *testing.T) {
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
		mse, psnr, err := gopsnr.ExecWithFileName(tt.filepath1, tt.filepath2)
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
