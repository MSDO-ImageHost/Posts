FROM golang:1.15 as builder
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go install /app/cmd/*


FROM scratch
COPY --from=builder /go/bin/app /usr/local/bin/
CMD ["app"]
