FROM golang:alpine AS builder
ADD . /app/
WORKDIR /app
RUN go build -o app

# -----

FROM alpine
ADD . /app/
WORKDIR /app
# copy dependent files
COPY --from=builder /go/pkg/mod/github.com/maxisme/ /go/pkg/mod/github.com/maxisme/
COPY --from=builder /app/app /app/app

HEALTHCHECK CMD curl --fail http://localhost:8080/ || exit 1
CMD ["./app"]