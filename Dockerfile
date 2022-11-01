### BUILD ###
FROM golang:1.19-alpine AS BUILD

WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go test ./... -cover
RUN go build -o ./out/aa-settlements ./main.go

### DEPLOY ###
FROM alpine:3.16
RUN apk add ca-certificates

COPY --from=BUILD /tmp/app/out/aa-settlements /app/aa-settlements
COPY --from=BUILD /tmp/app/locales /app/locales

WORKDIR /app

CMD ["./aa-settlements"]