FROM golang:1.18-alpine
ENV PORT=8080

RUN mkdir app
COPY main.go app/

EXPOSE ${PORT}

CMD go run app/main.go