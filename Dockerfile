# Start from base image 1.16.4:
FROM golang:1.16.4

ENV CACHE_MEMORY_MANAGEMENT_ENABLED=true

# Configure the repo url so we can configure our work directory:
ENV REPO_URL=github.com/iloncar89/calculation-api

# Setup out $GOPATH
ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/$REPO_URL

# /app/src/github.com/iloncar89/calculation-api/src

# Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN go test ./...

# Build the Go app
RUN go build -o ./out/calculation-api .

# Expose port 8080 to the world:
EXPOSE 8080

# Run the binary program produced by go install
CMD ["./out/calculation-api"]