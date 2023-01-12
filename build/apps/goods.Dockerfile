FROM golang:1.19.5-bullseye as builder
LABEL stage="builder"
ENV CGO_ENABLED 0
ENV GOOS linux
RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /goods
ADD go.mod ./
ADD go.sum ./
RUN go mod download
COPY . .
RUN go build -o app /cmd/main.go

FROM alpine
LABEL stage="deploy"
RUN apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /goods
COPY --from=builder /usr/share/zoneinfo/Europe/Moscow /usr/share/zoneinfo/Europe/Moscow
ENV TZ Europe/Moscow
COPY --from=builder /goods/app /goods/app
CMD [". /app"]