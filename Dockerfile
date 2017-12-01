FROM golang:1.9-alpine
RUN mkdir -p /usr/local/go/src/github.com/zcrawl
COPY . /usr/local/go/src/github.com/zcrawl/zcrawl
WORKDIR /usr/local/go/src/github.com/zcrawl/zcrawl
# RUN go get -u
RUN mkdir -p /opt/zcrawl
WORKDIR /opt/zcrawl
RUN go build -x github.com/zcrawl/zcrawl/cmd/zcrawl-server
CMD ["/opt/zcrawl/zcrawl-server"]
EXPOSE 9999
