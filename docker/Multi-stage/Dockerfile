# build stage
FROM golang:1.12-alpine3.9 AS tung-97
ENV WDIR my-code
RUN mkdir -p  $HOME/$WDIR
ADD . $HOME/$WDIR
WORKDIR  $HOME/$WDIR
RUN GOOS=linux go build 
RUN ls
RUN echo $HOME/$WDIR
RUN echo $WDIR
RUN pwd

# final stage
FROM alpine:3.8
RUN mkdir -p /app/
WORKDIR /app
# Copy result binary go app to /app folder
COPY --from=tung-97 /my-code/example /app/
ENTRYPOINT ["./example"]
EXPOSE 8080


