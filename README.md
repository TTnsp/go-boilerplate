# go boilerplate
This project is a boilerplate to write API with golang. It's manage :
- connection to postgres SQL database
- database schema migration at startup
- configuration with yaml file and environments variables
- exposing endpoints for API
- authentication with JWT token

## Get started
Start postgres SQL server
``` bash
docker run --rm -p 5432:5432 -e POSTGRES_PASSWORD=test -e POSTGRES_USER=test postgres
```

Start application
``` bash
go run .
```

## Structure

``` bash
.
├── auth
│   ├── bcrypt.go
│   └── jwt.go
├── configuration           
│   └── configuration.go    
├── controllers             # Add your API controllers here
│   ├── foo.go
│   └── user.go
├── models                  # Add your models for API and database here
│   ├── foo.go
│   └── users.go
├── repositories            # Postgres SQL connection
│   └── postgres.go
├── config.yml              # Add your configuration here
├── go.mod
├── go.sum
├── main.go
├── README.md
└── router.go               # Add your endpoint here
```