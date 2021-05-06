FROM golang

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .

RUN  go mod init vendor/github.com/gorilla/mux

RUN go mod download github.com/gorilla/mux


RUN make build_linux

EXPOSE 8000

CMD ["./build/go-calc"]
