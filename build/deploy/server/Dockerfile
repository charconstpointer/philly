FROM golang:1.14

COPY . /app
WORKDIR /app

ENV GO111MODULE=on

RUN apt-get update && apt-get install build-essential -y
RUN cd cmd/server && GOOS=linux go build -o serv
EXPOSE 8080
ENTRYPOINT [ "./cmd/server/serv" ]  