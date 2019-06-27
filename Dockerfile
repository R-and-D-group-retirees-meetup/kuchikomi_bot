FROM golang:1.12.6-alpine3.10
ENV LINE_ACCESS_TOKEN=""
COPY . /root/kuchikomi_bot
WORKDIR /root/kuchikomi_bot
RUN apk update \
    && apk add git  \
    && go build
CMD ./kuchikomi_bot