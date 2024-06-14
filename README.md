# go-example-project

## Create DB

```shell
docker run --name go-example-postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:15.2

docker start go-example-postgres
docker stop go-example-postgres
docker rm go-example-postgres

docker exec -it go-example-postgres bash
psql -U postgres

CREATE DATABASE greenlight;
\c greenlight
CREATE ROLE greenlight WITH LOGIN PASSWORD 'password';
CREATE SCHEMA greenlight authorization greenlight;


psql -U greenlight greenlight
create table mytable(id integer);
\dt

CREATE EXTENSION IF NOT EXISTS citext;
```

## Setup Db

```shell
psql -U postgres -c 'SHOW config_file;'
```

## Connect Db

```
postgres://greenlight:password@localhost/greenlight?sslmode=disable
```

## Migrate

```shell
migrate -path=./migrations -database="postgres://greenlight:password@localhost/greenlight?sslmode=disable" up
```

## Start application

```shell
go run ./cmd/api    
```

## Stop application

```shell
pgrep -l api
pkill -SIGKILL api
pkill -SIGTERM api
```



