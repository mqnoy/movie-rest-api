# Movie REST API

Movie Database api


## Prerequisites

1. [Golang v1.20+](https://go.dev/doc/install) , currently i used version 1.21.1
1. MariaDB v.10.2+
1. [migrate](https://github.com/golang-migrate/migrate)
1. Git


### Clone the project

```
$ git clone git@github.com:mqnoy/movie-rest-api.git
$ cd movie-rest-api
$ go get
```



### Get Started
1. Open the project directory with code editor eg: [Visual Studio Code](https://code.visualstudio.com/download)
1. Open the terminal 
1. Install dependencies type in terminal `go get`
1. Navigate to your mysql client then create database eg: `movie_db`
1. Back to terminal in vscode then run migration 
   ```
    migrate -path migration/script -database "mysql://your-db-username:your-db-password@tcp(127.0.0.1:3306)/movie_db" up
   ```
1. Copy `.env.example` to `.env` in root project directory
1. Edit variable on `.env` for detail see the table
   | Variabel          |        description        |  Example value |
   |-------------------|---------------------------|------    |
   | SERVER_HOST       |  Server listen all interface     | 0.0.0.0    |
   | SERVER_PORT       |  Server listen port       |  8080 |
   | MYSQL_HOST        | mysql host                |  localhost |
   | MYSQL_PORT        | mysql port                |    3306 |
   | MYSQL_USERNAME    | user mysql                |    root |
   | MYSQL_PASSWORD    | password user mysql       |    your-password-root |
   | MYSQL_DB_NAME     | database name             |    movie_db |

1. Run service `go run main.go`


### Library that i use
- [Viper](https://github.com/spf13/viper) - for handle configuration
- [GORM](https://github.com/go-gorm/gorm) - Orm for golang with [mysql driver](gorm.io/driver/mysql)
- [gin-gonic](https://github.com/gin-gonic/gin) - http server (routing, middleware) 
- [validator](https://pkg.go.dev/github.com/go-playground/validator/v10@v10.16.0) - validator tag based
- [zerolog](https://pkg.go.dev/github.com/rs/zerolog@v1.31.0) - logging


### TODO 
- [ ] Support with unit test and mocking

### Contributor
- [Rifky](https://github.com/mqnoy/)