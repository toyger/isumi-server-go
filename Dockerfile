FROM frolvlad/alpine-glibc:latest

COPY ./bin /usr/local/bin

ENTRYPOINT ./isumi

EXPOSE 8080
