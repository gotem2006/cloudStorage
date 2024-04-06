FROM golang:1.22

WORKDIR /app


RUN go mod init app

RUN go mod tidy




COPY . .


RUN go build cmd/app/main.go


CMD ["./main"]