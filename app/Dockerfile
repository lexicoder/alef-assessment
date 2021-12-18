FROM golang:alpine as builder
WORKDIR /app 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .
FROM scratch
WORKDIR /app
COPY --from=builder /app/assessment /usr/bin/
ADD views  /app/views
ADD assets /app/assets
ENTRYPOINT ["assessment"]
