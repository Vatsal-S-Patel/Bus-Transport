FROM golang:1.21 AS build_image
WORKDIR /exe
COPY . /exe/Backend
WORKDIR /exe/Backend
RUN go mod tidy
RUN go build
WORKDIR /exe
RUN mv /exe/Backend/busproject /exe/


FROM scratch
EXPOSE 8080
COPY --from=build_image /exe/busproject .
CMD ./busproject
#CMD ./app/busproject
