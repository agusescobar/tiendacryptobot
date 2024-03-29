FROM golang:1.18.1

# Move to working directory '/build'
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main cmd/bot/main.go

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .
COPY .env /dist

# Command to run when starting the container
CMD ["/dist/main"]