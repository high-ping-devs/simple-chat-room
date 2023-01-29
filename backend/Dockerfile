FROM golang:latest

WORKDIR /usr/src/app

COPY . .
RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin/app .

EXPOSE 8080

CMD ["app"]