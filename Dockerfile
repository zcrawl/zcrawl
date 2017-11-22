FROM golang:1.9-alpine
RUN mkdir -p /usr/local/go/src/github.com/matiasinsaurralde
COPY . /usr/local/go/src/github.com/matiasinsaurralde/zcrawl-platform
WORKDIR /usr/local/go/src/github.com/matiasinsaurralde/zcrawl-platform
# RUN go get -u
RUN mkdir -p /opt/zcrawl-platform
WORKDIR /opt/zcrawl-platform
RUN go build -x github.com/matiasinsaurralde/zcrawl-platform/cmd/zcrawl-server
CMD ["/opt/zcrawl-platform/zcrawl-server"]
EXPOSE 9999
