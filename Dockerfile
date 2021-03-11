FROM golang:1.14.0
RUN mkdir /app
COPY ./backend /app
WORKDIR /app
RUN ls
#RUN go build -o main server.go
RUN go run server.go &
EXPOSE 3080

FROM node:lts-alpine
RUN npm install -g http-server
WORKDIR /app
COPY frontend/package*.json ./
RUN npm install

COPY frontend/ .

EXPOSE 8080
CMD ["npm","run","serve"]
