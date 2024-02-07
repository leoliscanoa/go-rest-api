FROM golang:1.21.7-alpine
LABEL authors="lliscano"

RUN mkdir /app
ADD . /app
WORKDIR /app

ENV DB_HOST ""
ENV DB_USER ""
ENV DB_PASSWORD ""
ENV DB_NAME ""
ENV DB_PORT ""

RUN go build -o main .
CMD ["/app/main"]