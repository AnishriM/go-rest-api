**go-rest-api**
----
[Note: This repo is under progress]


The aim of this project is to create template layout and project for getting started wit REST API development using golang.

#### packages used:
1. For database: postgres
2. For database connection & handling: gorm
3. For logging: logrus
4. For api development: gorilla

#### Requirement for application:
1. Linux system
2. Docker
3. golang 1.19 version

#### How to run application on local setup?
1. create docker container for DB using following command.
```
docker run --name some-postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
```
2. Export following environment variables.
```
export DB_USERNAME=postgres
export DB_PASSWORD=postgres
export DB_HOST=localhost
export DB_PORT=5432
export DB_TABLE=postgres
```
3. execute below command to start the application.
```
go run cmd/server/main.go
```


![Number of Closed PRs](https://img.shields.io/github/issues-pr-closed-raw/AnishriM/go-rest-api)
![Repo Size](https://img.shields.io/github/repo-size/AnishriM/go-rest-api)
----
