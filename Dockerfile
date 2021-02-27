FROM golang:1.16.0-alpine3.13

WORKDIR go/src/app
COPY . .

ARG MONGO_ENDPOINT
ARG MONGO_PORT
ARG MONGO_USERNAME
ARG MONGO_PASSWORD
ARG SECRET
ARG BASE_URL

ENV MONGO_ENDPOINT="mongo"
ENV MONGO_PORT="27017"
ENV MONGO_USERNAME="singh"
ENV MONGO_PASSWORD="12345678"
ENV SECRET="LOGICISKING"
ENV BASE_URL="http://127.0.0.1:8000"

RUN go get
RUN go build

CMD ["./pdf-backend"]