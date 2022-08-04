FROM golang:1.17-buster as builder

WORKDIR /app

RUN apt update && \
    apt upgrade -y && \
    apt install -y apt-utils upx

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make build && make compress


FROM scratch as final

WORKDIR /app

COPY --from=builder /app/build .

EXPOSE 50051

CMD ["./app"]