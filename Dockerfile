FROM golang:alpine

WORKDIR /src

COPY . .

RUN ls -l

RUN go clean --modcache

RUN go build -o main .

EXPOSE 4000

ENTRYPOINT ["./main"]