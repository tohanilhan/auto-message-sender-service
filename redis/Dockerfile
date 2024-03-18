# Dockerfile for 'klearis-redis-auth'
FROM redis:7.0-alpine


# set timezone
RUN apk add --no-cache tzdata

# copy config
COPY ./conf/redis.conf /usr/local/etc/redis/

CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]
