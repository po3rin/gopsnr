# gopsnr

this package calucurate MSE/PSNR.

## MSE/PSNR

PSNR is most easily defined via the mean squared error (MSE). Given a noise-free m×n monochrome image I and its noisy approximation K, MSE is defined as:

<img src="./static/3a34719b4f391dba26b3e8e4460b7595d62eece4.svg"></img>
<img src="./static/fc22801ed1232ff1231c4156b589de5c32063a8a.svg"></img>

## Quick Start

```bash
make run CMD="go run cmd/gopsnr/main.go img1.jpeg img2.jpeg"

compare: img1.jpeg & img2.jpeg
[MSE] value: +4.092628e-002 [PSNR] value: +6.201078e+001
```

## Reference
[画像処理の基本的なアルゴリズムをgolangで復習 １（平滑化）](https://qiita.com/ikeponsu/items/ccc346e08335a29a55e7#%E3%82%BD%E3%83%BC%E3%82%B9%E3%82%B3%E3%83%BC%E3%83%89-1)
[Image Quality Assessment through FSIM, SSIM, MSE and PSNR―A Comparative Study](https://file.scirp.org/Html/2-1730990_90911.htm)
