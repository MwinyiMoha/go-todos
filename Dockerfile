FROM golang:1.24-bookworm AS builder

ARG UPX_VERSION=4.2.4
ARG USER_ID=10001

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends apt-utils xz-utils ca-certificates && \
    curl -Ls https://github.com/upx/upx/releases/download/v${UPX_VERSION}/upx-${UPX_VERSION}-amd64_linux.tar.xz -o - | tar xvJf - -C /tmp && \
    cp /tmp/upx-${UPX_VERSION}-amd64_linux/upx /usr/local/bin/ && \
    chmod +x /usr/local/bin/upx && \
    apt-get remove -y xz-utils && \
    rm -rf /var/lib/apt/lists/* && \
    useradd -u ${USER_ID} app

COPY . .

RUN make build && \
    make compress_binary && \
    make test_binary


FROM scratch AS final

ARG WEB_PORT=8080

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /app/build .

EXPOSE ${WEB_PORT}

USER app

ENTRYPOINT [ "./go-todos" ]