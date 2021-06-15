FROM elsaunier/go-promtool-amtool:latest as builder

MAINTAINER Loïc Saunier

RUN mkdir -p configChecker
COPY . /configChecker
WORKDIR /configChecker

FROM ubuntu:20.04 as app

COPY --from=builder /usr/bin/promtool /usr/bin/promtool 
COPY --from=builder /usr/bin/amtool /usr/bin/amtool
RUN mkdir -p app/
COPY --from=builder /configChecker/build /app/configChecker
COPY --from=builder /configChecker/templates /app/templates

EXPOSE 8181

WORKDIR /usr/bin
ENTRYPOINT ["/app/configChecker"]
