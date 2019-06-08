FROM denismakogon/gocv-alpine:4.0.1-buildstage as build-stage

RUN go get -u -d gocv.io/x/gocv
RUN cd $GOPATH/src/gocv.io/x/gocv && go build -o $GOPATH/bin/gocv-version ./cmd/version/main.go
WORKDIR /go/src/github.com/po3rin/gopsnr
