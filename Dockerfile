FROM golang:1.19-alpine AS BUILD

WORKDIR /app

COPY go.mod /app

COPY go.sum /app

RUN go mod download

COPY . /app

RUN go build -o jamie-service

FROM alpine:3.17

COPY --from=BUILD /app/jamie-service /bin/

CMD [ "jamie-service" ] 



