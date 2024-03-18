.PHONY: all

all: help

## Help
help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

## deploy-all: Build and start all containers
deploy-all:
	make redis-deploy
	make postgres-deploy
	sleep 5; 
	make api-deploy
	make scheduler-deploy

## deploy-from-scratch: Build and start all containers with new volume
deploy-all-from-scratch:
	make redis-deploy-with-new-volume
	make postgres-deploy-with-new-volume
	sleep 5;
	make api-deploy-from-scratch
	make scheduler-deploy-from-scratch

## remove-all: Remove all containers and images
remove-all:
	make redis-remove
	make postgres-remove

## up-all: Start all containers
up-all:
	make redis-up
	make postgres-up

## down-all: Stop all containers
down-all:
	make redis-down
	make postgres-down

## build-redis: Build redis image
build-redis: 
	cd redis; docker build -f Dockerfile -t webhook-redis-img .

## redis-up: Start redis container	
redis-up: 
	docker-compose up webhook-redis-svc -d

## redis-down: Stop redis container
redis-down:
	docker-compose down webhook-redis-svc 

## redis-logs: Show redis logs
redis-logs:
	docker-compose logs webhook-redis-svc -f

## redis-remove: Remove redis container and image
redis-remove:
	docker-compose rm -fsv webhook-redis-svc
	docker rmi -f webhook-redis-img
	rm -rf webhook-redis-vol

## redis-deploy: Build and start redis container with logs	
redis-deploy:
	make build-redis;
	make redis-up;

## redis-deploy-with-new-volume: Build and start redis container with logs and new volume
redis-deploy-with-new-volume:
	make redis-remove;
	make build-redis;
	make redis-up;

## build-postgres: Build postgres image
build-postgres: 
	cd postgresql; docker build -f Dockerfile -t webhook-postgres-img .

## postgres-up: Start postgres container
postgres-up:
	docker-compose up webhook-postgres-svc -d 

## postgres-down: Stop postgres container
postgres-down:
	docker-compose down webhook-postgres-svc

## postgres-logs: Show postgres logs
postgres-logs:
	docker-compose logs webhook-postgres-svc 

## postgres-remove: Remove postgres container and image
postgres-remove:
	docker-compose rm -fsv webhook-postgres-svc
	docker rmi -f webhook-postgres-img
	rm -rf webhook-postgres-vol

## postgres-deploy: Build and start postgres container with logs
postgres-deploy:
	make build-postgres
	make postgres-up

## postgres-deploy-with-new-volume: Build and start postgres container with logs and new volume
postgres-deploy-with-new-volume:
	make postgres-remove
	make build-postgres
	make postgres-up

## api-build: Build api image
api-build:
	
	# build application
	
	cd api; go mod tidy
	cd api; go mod download
	cd api; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o prog .

	cd api; docker build -f Dockerfile -t webhook-api-img .

## api-up: Start api container
api-up:
	docker-compose up webhook-api-svc -d

## api-down: Stop api container
api-down:
	docker-compose down webhook-api-svc

## api-logs: Show api logs
api-logs:
	docker-compose logs webhook-api-svc -f

## api-remove: Remove api container and image
api-remove:
	docker-compose rm -fsv webhook-api-svc
	docker rmi -f webhook-api-img

## api-deploy: Build and start api container with logs
api-deploy:
	make api-build
	make api-up

## api-deploy-from-scratch: Build and start api container with logs and new volume
api-deploy-from-scratch:
	make api-remove
	make api-build
	make api-up

## scheduler-build: Build scheduler image
scheduler-build:
	
	# build application
	cd scheduler; go mod tidy
	cd scheduler; go mod download
	cd scheduler; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o prog .

	cd scheduler; docker build -f Dockerfile -t webhook-scheduler-img .

## scheduler-up: Start scheduler container
scheduler-up:
	docker-compose up webhook-scheduler-svc -d

## scheduler-down: Stop scheduler container
scheduler-down:
	docker-compose down webhook-scheduler-svc

## scheduler-logs: Show scheduler logs
scheduler-logs:
	docker-compose logs webhook-scheduler-svc -f

## scheduler-remove: Remove scheduler container and image
scheduler-remove:
	docker-compose rm -fsv webhook-scheduler-svc
	docker rmi -f webhook-scheduler-img

## scheduler-deploy: Build and start scheduler container with logs
scheduler-deploy:
	make scheduler-build
	make scheduler-up

## scheduler-deploy-from-scratch: Build and start scheduler container with logs and new volume
scheduler-deploy-from-scratch:
	make scheduler-remove
	make scheduler-build
	make scheduler-up
