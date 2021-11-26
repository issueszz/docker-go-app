FROM golang:1.16-alpine

WORKDIR /goapp
#RUN mkdir ./model
RUN mkdir ./services

COPY main.go ./
#COPY ./model ./model
COPY ./services ./services

COPY go.mod ./
COPY go.sum ./

RUN go mod download

RUN go build -o /main

CMD ["/main"]

#FROM gcr.io/distroless/base-debian10
#
#WORKDIR /
#
#COPY --from=build /main /main
#
#USER nonroot:nonroot
#
#ENTRYPOINT ["/main"]

