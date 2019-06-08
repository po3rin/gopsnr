PHONY: run
run:
	docker run --rm -v ${PWD}:/go/src/github.com/po3rin/gopsnr \
	gocv-playground /bin/sh -c "${CMD}"
build:
	docker build ./ -t gocv-playground
