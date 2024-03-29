- command to run postgres on docker:
```
    docker run --name postgres -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:alpine3.19
```

- running makefile
```
    make postgresinit
```

- command to get list image
```
    docker ps
```

- command to create database on docker postgres with name go-chat:
```
    docker exec -it postgres createdb --username=root --owner=root go-chat
```

- run postgres docker console:
```
    docker exec -it postgres psql
```

> create model through db.go on db folder

> add underscore before module to prevent unused module in import being deleted after saving

> Install golang-migrate to migrate your database

- automatically install the depedency
```
    go mod tidy
```

- create migration with extension sql with table name add_user_table to folder db/migrations
```
    migrate create -ext sql -dir db/migrations add_user_table
```

```
- up file are for creating table from migration
- down file are for go backwards and deleting the table created by migrations
```
- specify the path for migration
```
    migrate -path <migrations folder locations> -database "postgres://<username>:<password>@<server-that-run>:<port-that-server-run>" -verbose up
```
**Ex**:
```
    migrate -path db/migrations -database "postgres://root:password@localhost:5433" -verbose up
```

> server/internal/users/ are model directories for table users(It's you that need to create it)
> users_repository.go serve as packages that run sql queries for users that you need to create
> users.go serve as models for table users that you need to create

- Go syntax breakdown
```
// syntax: func (<receiver AKA inheritence>) <name function>(<parameter>) (<return value>,<return value>)
func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error)
```