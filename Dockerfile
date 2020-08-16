# build stage
FROM golang:alpine AS build-env
RUN apk add build-base
ADD . /src
RUN apk add build-base && cd /src && go build -o main

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/main /app/
ENTRYPOINT ./main
