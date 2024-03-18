# Dockerfile for 'klearis-cron-manager'
FROM alpine:3.12

# create work directory
RUN mkdir /prog

# switch to work directory
WORKDIR /prog

# copy application binary
COPY .env .

COPY prog .

# run application
ENTRYPOINT ["/prog/prog"]
