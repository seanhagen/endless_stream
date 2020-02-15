FROM alpine:latest
MAINTAINER Sean Hagen <sean.hagen@gmail.com>

RUN apk --no-cache add ca-certificates

WORKDIR /app
ADD server /app/server
ADD config.yaml /etc/es/config.yaml

ENV PORT 3000

CMD ["./server"]
