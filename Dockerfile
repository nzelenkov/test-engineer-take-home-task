# build stage
FROM golang:1.12.5-alpine3.9 AS builder
LABEL Author="Nikola.Zelenkov <z@aapack.live>"
LABEL stage=builder
ENV APPNAME=$APPNAME
ENV DOMAIN=$DOMAIN
ENV GOBIN=/go/bin
ENV GOPATH=/go
ENV CGO_ENABLED=0
ENV GOOS=linux
COPY . ${APPNAME}
RUN apk update && apk add git && cd ${APPNAME} \
go get -d -v && go build -v && go install -v && go test -v ./...

# final stage
FROM scratch
ENV APPNAME=$APPNAME
WORKDIR /usr/local/bin
COPY --from=builder /go/bin/$APPNAME .
CMD ["${APPNAME}"]