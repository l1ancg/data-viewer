FROM golang:1.21.1-alpine3.18 AS GB
COPY . /app
WORKDIR /app
RUN go version
RUN go install
RUN go build -o data-viewer

#FROM jitesoft/sqlite:latest
FROM alpine:3.12
RUN apk update \
    && apk add sqlite \
    && apk add socat
COPY --from=GB /app/data-viewer ./
RUN sqlite3 data-viewer.db
CMD ./data-viewer
