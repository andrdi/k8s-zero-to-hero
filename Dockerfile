FROM golang:1.18-alpine AS build
COPY . /app
WORKDIR /app
RUN go mod download
RUN go build -o /goapp cmd/main.go

FROM alpine:3.18 as backend
WORKDIR /
COPY --from=build /goapp /goapp
COPY ./templates /templates
COPY ./static /static
EXPOSE 8080
ENTRYPOINT ["/goapp"]

FROM nginx:stable-alpine as frontend
WORKDIR /www
COPY static /www/static/
COPY .werf/nginx.conf /etc/nginx/nginx.conf
