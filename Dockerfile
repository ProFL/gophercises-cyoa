FROM golang:1.16-alpine AS build

WORKDIR /opt/cyoa
COPY . .
RUN go build -o cyoa

FROM build AS dev

RUN apk add nodejs npm \
    && npm i -g nodemon

CMD ["nodemon", "main.go"]

FROM alpine:latest

WORKDIR /opt/cyoa
COPY --from=build /opt/cyoa/cyoa ./cyoa

EXPOSE 8080
CMD ["./cyoa"]
