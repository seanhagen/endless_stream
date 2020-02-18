FROM alpine:latest
MAINTAINER Sean Hagen <sean.hagen@gmail.com>

RUN apk --no-cache add ca-certificates

ADD server /server
ADD config.yaml /etc/es/config.yaml

ENV PORT 8000
EXPOSE 8000

ENTRYPOINT ["/server"]
