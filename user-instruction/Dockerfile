FROM golang:1.18-alpine

RUN mkdir /app

RUN addgroup -S cupcakezGroup
RUN adduser -S -D -h /app cupcakezUser cupcakezGroup
RUN chown -R cupcakezUser:cupcakezGroup /app

USER cupcakezUser

COPY main.go /app

EXPOSE 8080
CMD go run /app/main.go