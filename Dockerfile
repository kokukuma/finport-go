FROM golang:latest as builder

WORKDIR /go/src/github.com/kokukuma/finport-go

RUN go get github.com/golang/dep/cmd/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -v -vendor-only

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go install -v \
            -ldflags="-w -s" \
            -ldflags "-X main.version=1.0" \
            -ldflags "-X main.serviceName=finport-go" \
            github.com/kokukuma/finport-go/cmd/server

# runtime image
FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/bin/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]
