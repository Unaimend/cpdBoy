# Use the official Go image for Go 1.21
FROM golang:1.21

# Set environment variables for Go
ENV GO111MODULE=on \
    GOPATH=/go \
    PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

# Set the working directory inside the container
WORKDIR /app

# Copy the current project files to the container
COPY . /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    python3 \
    python3-pip \
    sqlite3 && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Run any required initialization or dependencies installation (optional)
RUN go mod tidy
RUN python3 generate_sqlite.py
RUN go get github.com/mattn/go-sqlite3


EXPOSE 3000

# Define the default command to run
CMD ["go", "run", "main.go"]

