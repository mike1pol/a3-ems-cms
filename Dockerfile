FROM golang:latest
RUN mkdir /app
COPY . /go/src/github.com/mike1pol/rms/
WORKDIR /go/src/github.com/mike1pol/rms/
RUN make

CMD cd migrations && go run *.go && cd .. && /go/src/github.com/mike1pol/rms/main
