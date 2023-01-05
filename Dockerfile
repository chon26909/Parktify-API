FROM golang:alpine

WORKDIR /src

COPY . .

RUN go mod download 

RUN go build main.go

EXPOSE 4000

ENTRYPOINT ["./main"]