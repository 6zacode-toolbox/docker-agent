FROM cgr.dev/chainguard/go:1.19
RUN apt-get update && \
    apt-get install -y unzip curl jq vim unzip less \
     && rm -rf /var/lib/apt/lists/* 
WORKDIR /root
COPY . .
# download dependencies and prepare fresh installation
RUN go mod download
RUN go build .
RUN  mv ./docker-agent /usr/local/bin/ && chmod 777 /usr/local/bin/docker-agent


FROM docker:23.0.0
RUN apk --no-cache add ca-certificates jq git
USER root
WORKDIR /home/app/
# setup user
RUN adduser -D app
COPY --chown=app:app --from=0 /usr/local/bin/docker-agent ./
COPY --chown=app:app ./scripts ./
RUN chmod +x *.sh
USER app


