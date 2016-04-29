FROM frolvlad/alpine-glibc:latest

ADD ./bin/isumi /usr/local/bin/isumi

ENTRYPOINT /usr/local/bin/isumi

EXPOSE 8080
