FROM golang:1.18
RUN apt-get update && \
    apt-get install -y unzip curl jq vim unzip less \
     && rm -rf /var/lib/apt/lists/* 
WORKDIR /root
COPY . .
# download dependencies and prepare fresh installation
RUN go mod download
RUN go build .
RUN  mv ./docker-agent /usr/local/bin/ && chmod 777 /usr/local/bin/docker-agent


FROM docker:23.0.0-cli-alpine3.17
RUN apk --no-cache add ca-certificates
USER root
WORKDIR /home/app/
# setup user
RUN adduser -D app
COPY --chown=app:app --from=0 /usr/local/bin/docker-agent ./
USER app

