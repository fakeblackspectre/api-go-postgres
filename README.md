#docker commands

docker run --name some-postgres -e POSTGRES_USER=fakeblackspectre -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres

docker exec -it some-postgres bash

psql -U fakeblackspectre --password

CREATE DATABASE gorm;

\c gorm

docker stop some-postgres

docker rm some-postgres

#commands to run the api

go run .

air .


