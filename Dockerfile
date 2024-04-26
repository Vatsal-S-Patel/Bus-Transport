FROM golang:1.21 AS build_image

WORKDIR /exe

COPY ./Backend /exe/Backend

WORKDIR /exe/Backend

RUN go mod tidy

RUN go build

WORKDIR /exe

RUN mv /exe/Backend/busproject /exe/


FROM ubuntu:latest AS final_image

EXPOSE 8080

COPY --from=build_image /exe/busproject .

CMD ./busproject