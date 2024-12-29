FROM golang:1.23

RUN mkdir /.cache && chmod 777 /.cache

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

RUN go install -tags 'postgres,mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.mod go.sum ./

RUN go mod download && go mod verify

USER 1000:1000

COPY . .

CMD ["air", "-c", ".air.toml"]