# Go GIN Web API

## Initalization

`go mod init` creating a go mod file (like a node file for nodejs projects)

## Packages

1.  `go get github.com/githubnemo/CompileDaemon` and `go install github.com/githubnemo/CompileDaemon` watch files for changes and rebuild
2.  `go get github.com/joho/godotenv` Easy to load environment variables
3.  `go get -u github.com/gin-gonic/gin` Gin Framework for Http Server
4.  `go get -u gorm.io/gorm` and `go get -u gorm.io/driver/postgres` Go ORM Library for Go

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
`go run migrate/migrate.go` , It will create the posts table in the database.
