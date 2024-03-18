# Dockerfile for 'message-sender-service-postgresl' database
FROM postgres:alpine3.19

# Environments
ENV POSTGRES_USER=management_user
ENV POSTGRES_PASSWORD=HwcGNMU0KTpABxgJLPxfHPHZKgEYUqRE
ENV POSTGRES_DB=message_sender_service_db
ENV POSTGRES_SCHEMA=message_sender_service_schema

# copy init sql
COPY ./sql/init.sql /docker-entrypoint-initdb.d
