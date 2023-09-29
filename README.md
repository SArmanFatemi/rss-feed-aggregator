## Useful Commands
| Title | Command|
|---|---|
| Run and build | `go build && ./rssagg.exe` |
| initializes and writes a new go.mod file | `go mod init` |
| Match `go.mod` file with source code | `go mod tidy`|
| Construct `vendors` and copy all packages | `go mod vendor` |
| Installing package | `go github.com/joho/godotenv`, `go get github.com/lib/pq`|

## Required and command tools for work with database
| Title | Command | Note|
|---|---|---|
|Install `goose`|`go install github.com/pressly/goose/v3/cmd/goose@latestH`|
|Run migrations - up| `goose postgres YOUR_DB_CONNECTION_STRING up`| Should run this command inside `migrations` folder|
|Run migrations - down| `goose postgres YOUR_DB_CONNECTION_STRING down`| Should run this command inside `migrations` folder|
|Generate type-safe go code based on queries using `sqlc`|`docker run --rm -v $(pwd):/src -w /src sqlc/sqlc generate`| 
## Run dependencies and start development
`docker-compose up -d`
