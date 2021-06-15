FROM elsaunier/go-promtool-amtool:latest as builder

MAINTAINER Loïc Saunier

ADD verify.go .
ADD main.go .
ADD /templates/ .

RUN go build

FROM scratch as app

COPY --from=builder /usr/bin/promtool /usr/bin/promtool 
COPY --from=builder /usr/bin/amtool /usr/bin/amtool
COPY --from=builder /configChecker /usr/bin/configChecker

EXPOSE 8181
VOLUME ["/usr/bin"]

WORKDIR /usr/bin
ENTRYPOINT ["./configChecker"]