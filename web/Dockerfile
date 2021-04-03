FROM golang:1.16
WORKDIR /go/src
COPY . .
RUN go build
ENV GODO_API http://api:3000
ENV GODO_ADDR :3001
ENTRYPOINT [ "./web" ]
