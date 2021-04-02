FROM golang:1.16
RUN apt-get update && apt-get install -y sqlite3
WORKDIR /go/src
COPY ./database ./database
RUN sqlite3 --init ./database/init.sql ./database/godo.sqlite3
COPY . .
RUN go build
ENV GODO_ADDR :3000
ENV GODO_DATABASE /go/src/database/godo.sqlite3
ENTRYPOINT [ "./api" ]
