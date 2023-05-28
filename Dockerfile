FROM golang:1.19-alpine

WORKDIR /app
COPY . .
RUN apk add build-base

RUN go mod download
RUN cd /app/cmd && go build -o /bloggingsystem

EXPOSE 3000

CMD [ "/bloggingsystem" ]
