FROM labpositiva/golang:1.8

MAINTAINER Luis Mayta <@slovacus>

COPY ./glide.yaml /go/src/go-messages/
# COPY ./glide.lock /go/src/go-messages/

# set working directory to /app/
WORKDIR /go/src/go-messages/

# install golang dependencies
RUN apk add --no-cache --update \
  git

RUN glide install
