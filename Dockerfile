FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . /go/src/
WORKDIR /go/src/
RUN make

# runtime image
FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
