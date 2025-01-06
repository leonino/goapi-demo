FROM golang:1.23

WORKDIR //Users//Dell//go//apiGo

COPY . .

EXPOSE 8000

RUN go build -o main main.go

CMD ["./main"]