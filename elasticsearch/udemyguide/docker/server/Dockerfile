FROM golang:1.23.1-bullseye

RUN mkdir /app
WORKDIR /app
COPY . .

RUN go install github.com/air-verse/air@latest
RUN go mod download

CMD ["air", "-c", ".air.toml"]
