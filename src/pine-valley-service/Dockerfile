# Latest golang image on apline linux
FROM golang:1.22-alpine

# Work directory
WORKDIR /app

# Installing dependencies
COPY src/ ./
RUN go mod download

# Copying all the files
COPY . .

# Starting our application
CMD ["go", "run", "main.go"]

# Exposing server port
EXPOSE 3002