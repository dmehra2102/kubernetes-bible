# Stage 1 : Build the Go application
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2 : Create the final, minimal image
FROM alpine:3.18
WORKDIR /root/
COPY --from=builder /app/main /root/
EXPOSE 8080
CMD [ "./main" ]