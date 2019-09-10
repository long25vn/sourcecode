FROM golang:alpine 
RUN mkdir -p  $HOME/src/my-code
ADD . $HOME/src/my-code
WORKDIR  $HOME/src/my-code
RUN apk add git
RUN GOOS=linux go build
RUN ls
RUN pwd
CMD ["./example"]
EXPOSE 8080
