
FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN cd ./server && go mod tidy

RUN cd ./server && go build -o ../builds/Server.exe ./main.go

EXPOSE 5000

CMD ["/app/builds/Server.exe"]