# syntax=docker/dockerfile:1
FROM golang:1.16-alpine AS stage

WORKDIR /microservice

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
COPY ./database/ ./database/

#RUN CGO_ENABLED=0 go build -a -installsuffix nocgo -o /microservice
RUN go build
EXPOSE 8005
#RUN chmod +x ./microservice
CMD ["./microservice"]