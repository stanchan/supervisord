FROM stanchan/go:1.13.3 AS builder

RUN apk add --no-cache curl make gcc g++ python linux-headers binutils-gold gnupg libstdc++ git

RUN go get -v -u github.com/stanchan/supervisord

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-linkmode external -extldflags -static" -o /usr/local/bin/supervisord github.com/stanchan/supervisord

FROM scratch

COPY --from=builder /usr/local/bin/supervisord /usr/local/bin/supervisord

ENTRYPOINT ["/usr/local/bin/supervisord"]
