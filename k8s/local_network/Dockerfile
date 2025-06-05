FROM alpine:latest

RUN apk update && apk add --no-cache dnsmasq

COPY dnsmasq.conf /etc/dnsmasq.conf

EXPOSE 53/udp

CMD ["dnsmasq", "-k"]
