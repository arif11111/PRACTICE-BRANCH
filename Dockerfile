FROM golang

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# select the working directory where the source code will be copied
WORKDIR /app

# copying the source code files in the working directory
COPY  ./src .

# downloading the module dependecny package requried for the application
RUN go mod download github.com/gorilla/mux

# running the cmake file command which basically build the golang application
RUN make build_linux

#exposing the port from where the application will  be accessible
EXPOSE 8000

#giving the path to the go binary file for execution
CMD ["./build/go-calc"]
