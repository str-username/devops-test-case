ARG BUILD_IMAGE="golang:latest"

FROM $BUILD_IMAGE as build

WORKDIR /build

COPY app/ .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -a -installsuffix cgo -o server main.go


FROM alpine:3.18

WORKDIR /app

COPY --from=build /build/server .

USER nobody:nobody

ENTRYPOINT ["./server"]