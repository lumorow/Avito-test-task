FROM golang:1.21

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN chmod +x wait-for-it.sh
RUN go mod download
RUN go build -o user-segmentation ./cmd/server.go

CMD ["./user-segmentation"]
