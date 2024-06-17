FROM golang:1.22 as build
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o go_weather

FROM scratch
WORKDIR /app
COPY --from=build /app/go_weather .

ENTRYPOINT ["./go_weather"]