# Dockerfile for 'api'
FROM alpine:3.12

# create work directory
RUN mkdir /prog

# switch to work directory
WORKDIR /prog

# copy application artifacts to current directory
COPY .env .

COPY prog .

# expose port
EXPOSE ${API_PORT}

# run application
ENTRYPOINT ["/prog/prog"]
