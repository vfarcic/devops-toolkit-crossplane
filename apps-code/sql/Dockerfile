FROM golang:1.17 AS build
ADD . /src
WORKDIR /src
RUN go get -d -v -t
RUN go build -v -o silly-sql-app

FROM alpine
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
EXPOSE 8080
ENV VERSION 1.0.5
CMD ["silly-sql-app"]
COPY --from=build /src/silly-sql-app /usr/local/bin/silly-sql-app
RUN chmod +x /usr/local/bin/silly-sql-app
