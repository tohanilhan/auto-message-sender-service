## Auto Message Sender Service

This project is a simple auto message sending service. System that automatically sends 2 messages retrieved from the database, which have not yet been sent, every 2 minutes.



## System Diagram

![Diagram](image.png)

## Kullanılan Teknolojiler

**Backend:** Go, Fiber

**Database:** PostgreSQL

**Cache:** Redis

**Deployment:** Docker, Docker Compose, Makefile



## Usage

First, clone the project:
```bash
git clone https://github.com/tohanilhan/auto-message-sender-service.git
````

Enter the project folder.
```bash
cd auto-message-sender-service;
```
Now, use Makefile to easily build and deploy this service.
```bash 
    make [$target]
```

#### Build, Deploy & Run Project

```bash 
    make deploy-all
    or
    make deploy-from-scratch
```

The only difference between **`make deploy-all`** and **`make deploy-from-scratch`** commands is that with make deploy-from-scratch, volumes are deleted before the deployment process begins.


#### Here are the list of all targets that can be used for this project


| **Target** | **Description**                       |
| :-------- | :-------------------------------- |
| `deploy-all:`      | Build and start all containers|
| `deploy-from-scratch:` | Build and start all containers with new volume|
| `remove-all:`      | Remove all containers and images|
| `up-all:`          | Start all containers|
| `down-all:`        | Stop all containers|
| `build-redis:`     | Build redis image|
| `redis-up:`        | Start redis container|
| `redis-down:`      | Stop redis container|
| `redis-logs:`      | Show redis logs|
| `redis-remove:`    | Remove redis container and image|
| `redis-deploy:`    | Build and start redis container with logs|
| `redis-deploy-with-new-volume:` | Build and start redis container with logs and new volume|
| `build-postgres:`  | Build postgres image|
| `postgres-up:`     | Start postgres container|
| `postgres-down:`   | Stop postgres container|
| `postgres-logs:`   | Show postgres logs|
| `postgres-remove:` | Remove postgres container and image|
| `postgres-deploy:` | Build and start postgres container with logs|
| `postgres-deploy-with-new-volume:` | Build and start postgres container with logs and new volume|
| `api-build:`       | Build api image|
| `api-up:`          | Start api container|
| `api-down:`        | Stop api container|
| `api-logs:`        | Show api logs|
| `api-remove:`      | Remove api container and image|
| `api-deploy:`      | Build and start api container with logs|
| `api-deploy-from-scratch:` | Build and start api container with logs and new volume|
| `scheduler-build:` | Build scheduler image|
| `scheduler-up:`    | Start scheduler container|
| `scheduler-down:`  | Stop scheduler container|
| `scheduler-logs:`  | Show scheduler logs|
| `scheduler-remove:`| Remove scheduler container and image|
| `scheduler-deploy:`| Build and start scheduler container with logs|
| `scheduler-deploy-from-scratch:` | Build and start scheduler container with logs and new volume|


## Tech Stack

**Backend:** Go, Fiber

**Database:** PostgreSQL

**Cache:** Redis

**Deployment:** Docker, Docker Compose, Makefile
