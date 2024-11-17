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

RUN  apt-get install -y curl 

# Run any required initialization or dependencies installation (optional)
RUN go mod tidy


EXPOSE 8080
EXPOSE 3000

# Define the default command to run
CMD ["go", "run", "main.go"]

