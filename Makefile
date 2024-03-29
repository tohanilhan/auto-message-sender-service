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

## deploy-all-from-scratch: Build and start all containers with new volume
deploy-all-from-scratch:
	make down-all
	make redis-deploy-with-new-volume
	make postgres-deploy-with-new-volume
	sleep 5;
	make api-deploy-from-scratch
	make scheduler-deploy-from-scratch

## remove-all: Remove all containers and images
remove-all:
	make redis-remove
	make postgres-remove
	make api-remove
	make scheduler-remove

## up-all: Start all containers
up-all:
	make redis-up
	make postgres-up
	sleep 5;
	make api-up
	make scheduler-up

## down-all: Stop all containers
down-all:
	make redis-down
	make postgres-down
	make api-down
	make scheduler-down

## build-redis: Build redis image
build-redis: 
	cd redis; docker build -f Dockerfile -t redis-img .

## redis-up: Start redis container	
redis-up: 
	docker-compose up redis-svc -d

## redis-down: Stop redis container
redis-down:
	docker-compose down redis-svc 

## redis-logs: Show redis logs
redis-logs:
	docker-compose logs redis-svc -f

## redis-remove: Remove redis container and image
redis-remove:
	docker-compose rm -fsv redis-svc
	docker rmi -f redis-img
	rm -rf redis-vol

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
	cd postgresql; docker build -f Dockerfile -t postgres-img .

## postgres-up: Start postgres container
postgres-up:
	docker-compose up postgres-svc -d 

## postgres-down: Stop postgres container
postgres-down:
	docker-compose down postgres-svc

## postgres-logs: Show postgres logs
postgres-logs:
	docker-compose logs postgres-svc 

## postgres-remove: Remove postgres container and image
postgres-remove:
	docker-compose rm -fsv postgres-svc
	docker rmi -f postgres-img
	rm -rf postgres-vol

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

	cd api; docker build -f Dockerfile -t api-img .

## api-up: Start api container
api-up:
	docker-compose up api-svc -d

## api-down: Stop api container
api-down:
	docker-compose down api-svc

## api-logs: Show api logs
api-logs:
	docker-compose logs api-svc -f

## api-remove: Remove api container and image
api-remove:
	docker-compose rm -fsv api-svc
	docker rmi -f api-img

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

	cd scheduler; docker build -f Dockerfile -t scheduler-img .

## scheduler-up: Start scheduler container
scheduler-up:
	docker-compose up scheduler-svc -d

## scheduler-down: Stop scheduler container
scheduler-down:
	docker-compose down scheduler-svc

## scheduler-logs: Show scheduler logs
scheduler-logs:
	docker-compose logs scheduler-svc -f

## scheduler-remove: Remove scheduler container and image
scheduler-remove:
	docker-compose rm -fsv scheduler-svc
	docker rmi -f scheduler-img

## scheduler-deploy: Build and start scheduler container with logs
scheduler-deploy:
	make scheduler-build
	make scheduler-up

## scheduler-deploy-from-scratch: Build and start scheduler container with logs and new volume
scheduler-deploy-from-scratch:
	make scheduler-remove
	make scheduler-build
	make scheduler-up
