FROM golang:latest AS builder
LABEL authors="vbg911"
WORKDIR /app
COPY . .
WORKDIR /app/cmd/parser
RUN go build -o task.exe .

FROM ubuntu:latest
COPY --from=builder /app/cmd/parser/task.exe task.exe
CMD ["task.exe"]
