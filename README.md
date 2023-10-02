# RSS Feed Aggregator
This project represents the remarkable outcome of [Wags Lane's](https://github.com/wagslane) diligent efforts in learning GoLang by releasing ["Learn Go"](https://www.youtube.com/watch?v=un6ZyFkqFKo) in collaboration with [FreeCodeCamp](https://www.freecodecamp.org/news/learn-golang-handbook/).

## Run dependencies and start development
To streamline the setup process for this project, which utilizes `Docker Compose`, you can conveniently use the following command. It will ensure that all the prerequisites, such as `postgres` and `pgAdmin`, are properly configured for seamless development:

```shell
docker-compose up -d
```

By executing this command, the required services will be initialized and made available in the background. This includes setting up the `postgres` database and the `pgAdmin` administration tool. The `-d` flag ensures that the services run in detached mode, allowing you to continue working in your terminal without interruption.

## Useful Commands
| Title | Command|
|---|---|
| Run and build | `go build && ./rssagg.exe` |
| Initializes and writes a new go.mod file | `go mod init` |
| Match `go.mod` file with source code | `go mod tidy`|
| Construct `vendors` and copy all packages | `go mod vendor` |
| Installing package | `go install github.com/joho/godotenv`, `go get github.com/lib/pq`|

## Working with database and migrations
For handling and managing migrations in this project, we are using the [`goose`](https://github.com/pressly/goose) package. To install goose, run the following command:

```shell
go install github.com/pressly/goose/v3/cmd/goose@latest
```

To execute an "up" migration on your desired database, use the following command (note that you can find the connection string in the `.env` file of the project):

```shell
goose postgres YOUR_DB_CONNECTION_STRING up
```

Here are some important notes to consider when working with migrations:

1. Make sure to run the migrations command inside the `migrations` folder.
2. To perform a downgrade migration on your database, change the last parameter of the previous command to `down`.

To generate type-safe Go code based on our migrations, we use SQLC. Here's the command to run [`sqlc`](https://docs.sqlc.dev/en/stable/overview/install.html) (assuming you have Docker installed):

```shell
docker run --rm -v $(pwd):/src -w /src sqlc/sqlc generate
```

Note: This assumes you are executing the commands in a terminal or command prompt.