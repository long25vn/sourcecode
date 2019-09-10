FROM golang:alpine 
RUN mkdir -p  $GOPATH/src/my-code
ADD . $GOPATH/src/my-code
WORKDIR  $GOPATH/src/my-code
RUN GOOS=linux go build
RUN ls
RUN pwd
ENTRYPOINT ["./my-code"]
EXPOSE 80

