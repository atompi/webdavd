FROM golang:1.18.0 as builder

WORKDIR /mysrc
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o webdavd

FROM scratch

WORKDIR /app
COPY --from=builder /mysrc/webdavd /app/webdavd
COPY examples/webdavd.yaml.example /app/webdavd.yaml

VOLUME /data
EXPOSE 8080

ENTRYPOINT ["/app/webdavd"]
