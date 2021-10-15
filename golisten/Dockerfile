# Start by building the application.
FROM golang:1.17 as build

WORKDIR /go/src/listen
COPY . .

RUN go get -d -v ./...
RUN GOOS=linux GOARCH=amd64 go build src/listen.go
RUN GOBIN=/go/bin go install -v ./...
RUN ls -l /go/bin

# Now copy it into our base image.
FROM redhat/ubi8-micro
RUN yum update -y
COPY --from=build /go/bin/listen /
USER 3001:3001

CMD [ "/listen" ]
