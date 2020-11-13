FROM golang:1.15 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go install ./cmd/*


FROM scratch
COPY --from=builder /go/bin/app /usr/local/bin/
EXPOSE 1000
CMD ["app"]
