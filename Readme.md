# Go Web API

This a template of GoLang Web API, using the Gin framework and Go ORM with PostgreSQL Database.

By Hidayat Arghandabi 2023

#### Implementations

Last Update 06/19/2023

- HTTP Request
- Authorization and Authentication
- Middleware (Auth, Logger)
- PostgreSQL DB
- Containerization
- .env
- Migration

## How to run using docker container

Run the `docker-compose up` command

## How to run/use this template without docker containers

1. Update DB Connection String in the .env file
2. Create migration located in migrate/migrate.go using `go run migrate/migrate.go`. It will create the posts and users table in the database.
3. Install packages `go mod tidy`
4. You can run the project from main.go using the vscode debugger or `CompileDaemon -command="./go-crud"` the Daemon tool, make sure it is installed.
5. Enjoy this project

## Initalization

`go mod init` creating a go mod file (like a node file for nodejs projects)

## Required Packages

1.  `go get github.com/githubnemo/CompileDaemon` and install it, so that it can be run as command-line tool `go install github.com/githubnemo/CompileDaemon` watch files for changes and rebuild
2.  `go get github.com/joho/godotenv` Easy to load environment variables
3.  `go get -u github.com/gin-gonic/gin` Gin Framework for Http Server
4.  `go get -u gorm.io/gorm` and `go get -u gorm.io/driver/postgres` Go ORM Library for Go
5.  `go get -u github.com/sirupsen/logrus` Package for logging in a Gin application

#### for JWT authentication

5. `go get -u golang.org/x/crypto/bcrypt` for cryptography visit https://pkg.go.dev/golang.org/x/crypto
6. `go get -u github.com/golang-jwt/jwt/v4` jwt package visit https://pkg.go.dev/github.com/golang-jwt/jwt or https://github.com/golang-jwt/jwt

## Running Development Server

`CompileDaemon -command="./go-crud"` Run the Daemon and given the package name

## Declaring Models

https://gorm.io/docs/models.html visit the link to see how to declare your models.
we are using the following model

```go

type User struct {
  gorm.Model
  Name string
}
// equals
type User struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
  Name string
```

to make auto `db.AutoMigrate(&Product{})` visit link https://gorm.io/docs/index.html

as migrate/migrate.go is created run the migration using
`go run migrate/migrate.go`, It will make the posts table in the database.

## Go ORM Create

visit the link https://gorm.io/docs/create.html for Go ORM

## Important Commands

- `docker-compose up --build`
