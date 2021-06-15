FROM goland:1.16 as builder

MAINTAINER Loïc Saunier

COPY . .
RUN go build

FROM scratch as app

COPY --from=builder /usr/bin/promtool /usr/bin/promtool 
COPY --from=builder /usr/bin/amtool /usr/bin/amtool
COPY --from=builder /configChecker /usr/bin/configChecker

EXPOSE 8181
VOLUME ["/usr/bin"]

WORKDIR /usr/bin
ENTRYPOINT ["./configChecker"]