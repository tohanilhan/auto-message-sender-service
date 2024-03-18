# Dockerfile for 'klearis-redis-auth'
FROM redis:alpine3.19


# set timezone
RUN apk add --no-cache tzdata

# copy config
COPY ./conf/redis.conf /usr/local/etc/redis/

CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]
