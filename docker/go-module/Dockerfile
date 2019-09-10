FROM golang:alpine 
RUN mkdir -p  $HOME/src/my-code
ADD . $HOME/src/my-code
WORKDIR  $HOME/src/my-code
RUN GOOS=linux go build
RUN ls
RUN pwd
ARG PASSWORD="first"
RUN echo $PASSWORD
CMD ["./example"]
EXPOSE 8080
