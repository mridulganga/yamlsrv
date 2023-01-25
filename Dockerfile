FROM golang:alpine as builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -tags netgo -a -v -installsuffix cgo -o bin/yamlsrv main.go


FROM alpine:3
RUN apk update \
    && apk add --no-cache curl wget \
    && apk add --no-cache ca-certificates \
    && update-ca-certificates 2>/dev/null || true

COPY --from=builder /app/bin/yamlsrv /yamlsrv
COPY ./main.yaml /main.yaml

CMD ["/yamlsrv"]

ENV PORT=3000
EXPOSE $PORT