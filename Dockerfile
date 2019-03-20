FROM golang:1.8

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/minhajuddinkhan/muntaha
COPY . .
RUN dep ensure -v
#RUN go install github.com/minhajuddinkhan/muntaha/...

CMD ["go", "run", "bin/muntaha/main.go", "serve"]
