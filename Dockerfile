ARG OS="linux"
FROM docker:latest

MAINTAINER Loïc Saunier

RUN mkdir -p /configchecker
WORKDIR /configchecker

COPY LICENSE LICENSE/
COPY README ./
COPY testdata/ testdata/
COPY templates/ templates/
COPY main.go ./
COPY verify_test.go ./
COPY verify.go ./
COPY Makefile ./
COPY img/ img/
COPY data/ data/

EXPOSE 8181
VOLUME ["/configchecker"]

ENTRYPOINT ["make"]