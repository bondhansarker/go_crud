FROM golang:alpine

WORKDIR /usr/src/app

COPY . .

RUN apk update && apk add curl && apk add bash
RUN echo "Downloading cosmtrek/air for live reload"
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

ENTRYPOINT ["air"]