#docker commands

docker run --name some-postgres -e POSTGRES_USER=fakeblackspectre -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres