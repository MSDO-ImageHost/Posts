FROM golang:1.15 as build-stage

WORKDIR /app
COPY src/ .

RUN go mod tidy
RUN go mod download




FROM golang:1.15 as app-stage

COPY --from=build-stage /app/ /app/

EXPOSE 1000

WORKDIR /app
RUN go build -o app .
CMD /app/app
