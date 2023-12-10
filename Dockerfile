FROM golang:1.21.3-alpine
ENV CGO_ENABLED=0

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o bin/goose .

EXPOSE $PORT

CMD [ "./bin/goose" ]