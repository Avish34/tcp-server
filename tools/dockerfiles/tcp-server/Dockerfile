FROM golang:1.19

# Set destination for COPY
WORKDIR /app

# Copy everything
COPY ../../. .

# Downloading go mod dependencies
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /tcp-server

# Expose port 8080, as tcp-server listen on 8080
EXPOSE 8080

# Run
CMD [ "/tcp-server" ]