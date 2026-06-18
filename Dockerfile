From golang:1.26

WORKDIR /app

COPY /scooter-prj/go.mod ./

RUN go mod download

CMD ["air"]
