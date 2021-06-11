FROM golang:1.16 as builder
 
MAINTAINER Loïc SAUNIER
 
RUN apt update
RUN apt install -y prometheus prometheus-alertmanager

FROM golang:1.16 as app
 
COPY --from=builder promtool /usr/bin/promtool
COPY --from=builder amtool /usr/bin/amtool
