FROM golang:latest
COPY . ./bin/pe-sample-api
WORKDIR /go/bin/pe-sample-api
RUN ls -l
ENTRYPOINT /go/bin/pe-sample-api/sample-api

EXPOSE 8080
