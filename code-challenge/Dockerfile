# Go imagen base
FROM golang:1.21

# Stablish the working directory inside the container
WORKDIR /app

# Copy the files needed for the build
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project to the container
COPY . .

# Build the application binary
RUN go build -o /main ./cmd/api

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the applicationß
CMD ["/main"]
