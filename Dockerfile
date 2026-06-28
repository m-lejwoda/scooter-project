From golang:1.26

WORKDIR /app

RUN go install github.com/air-verse/air@v1.65.3

CMD ["air"]
