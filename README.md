# jitapi

`jitapi` is a package for Go, designed to fetch data from the database. It also provides a multiple functions to save the data in a csv or a json file.

- **Connect**:  Connect is used to connect to database. It will return an error for better error handling.
- **Query**: Query function returns the data that is fetch using the query passed while calling the function. It also returns an error.
- **DbConfig**: Takes the Configuration of database connection.
- **Postgres**: Generates a connection string with the given parameters and creates a connection pool.
- **Connect**: Connects to the database

## Installation 

To install `jitapi`, run:

```sh
go get github.com/parthvinchhi/jitapi
```

### Usage
Here's an example of how to use `jitapi`:

```
package main

import (
	"fmt"
	"log"

	"github.com/parthvinchhi/jitapi"
)

func main() {
	dbconfig := jitapi.DbConfig{
		// DBType:     "db_type",
		DBName:     "db_name",
		DBPort:     "db_port",
		DBHost:     "db_host",
		DBUser:     "db_user",
		DBPassword: "db_password",
		DBSslMode:  "db_sslmode",
	}

	// To connect to postgresql
	connectionString := jitapi.Postgres{
		Config: dbconfig,
	}

	if err := connectionString.Connect(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	query := "your sql query"

	data, err := connectionString.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
```