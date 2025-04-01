# Purplish Project - Items

Micro for the items functionality

## Setting up the project

Follow the next instructions to get the micro up and running

### Dependencies

Make sure you have the next dependencies on the target system:

- sqlite3
- go

### Deploying

Run the next commands on your system:

```sh
./.bin/run-migrations.sh
go mod tidy
cp -rvf ./.env{.example,} # Edit the .env file if needed.
go run cmd/api/main.go
```
