# Start by building the application.
FROM golang:1.17 as build

WORKDIR /go/src/listen
COPY . .

RUN go get -d -v ./...
RUN GOOS=linux GOARCH=amd64 go build src/listen.go
RUN GOBIN=/go/bin go install -v ./...
RUN ls -l /go/bin

# Now copy it into our base image.
FROM redhat/ubi8

RUN dnf -y update 
RUN dnf -y install python39

RUN pip3 install --upgrade pip
RUN pip3 install ansible-core
RUN pip3 install ansible-runner

RUN adduser runner

RUN for dir in \
      /home/runner \
      /home/runner/.ansible \
      /home/runner/.ansible/tmp \
      /runner \
      /home/runner \
      /runner/env \
      /runner/inventory \
      /runner/project \
      /runner/artifacts ; \
    do mkdir -m 0775 -p $dir ; chmod -R g+rwx $dir ; chgrp -R root $dir ; done && \
    for file in \
      /home/runner/.ansible/galaxy_token \
      /etc/passwd \
      /etc/group ; \
    do touch $file ; chmod g+rw $file ; chgrp root $file ; done
ENV HOME=/home/runner
ENV USERNAME=runner

COPY --from=build /go/src/listen/listen /usr/bin

CMD [ "/usr/bin/listen" ]
