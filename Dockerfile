FROM golang:1.16 as builder
 
MAINTAINER Loïc SAUNIER
 
RUN apt update
RUN apt install -y prometheus prometheus-alertmanager
RUN whereis promtool

FROM golang:1.16 as app

COPY --from=builder /usr/bin/promtool /usr/bin/promtool
COPY --from=builder /usr/bin/amtool /usr/bin/amtool
