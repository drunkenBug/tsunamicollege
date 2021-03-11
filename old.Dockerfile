FROM golang:1.14.0
RUN mkdir /app
ADD ./backend /app
WORKDIR /app
RUN go build -o main server.go
CMD ["go","run","server.go"]

FROM node:lts-alpine
RUN npm install -g http-server
WORKDIR /app
COPY frontend/package*.json ./
RUN npm install

COPY frontend/ .

EXPOSE 8080
CMD ["npm","run","serve"]
