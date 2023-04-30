#Build stage
FROM golang:1.19-alpine3.16 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main main.go


# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY scripts/start.sh .
# wait for postgres to be ready
COPY scripts/wait-for.sh . 
COPY db/migration ./db/migration

EXPOSE 8080
# CMD run whenever app start
CMD [ "/app/main" ] 
ENTRYPOINT [ "/app/start.sh" ]