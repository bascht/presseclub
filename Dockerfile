FROM golang:1.17-alpine3.15 AS presseclub

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /usr/local/bin/presseclub

FROM capsulecode/singlefile
COPY --from=presseclub /usr/local/bin/presseclub /usr/local/bin/presseclub
COPY entrypoint.sh /usr/local/bin/entrypoint.sh

EXPOSE 3000

ENTRYPOINT "entrypoint.sh"
