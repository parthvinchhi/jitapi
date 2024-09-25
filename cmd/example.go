package main

import (
	"fmt"
	"log"

	"github.com/parthvinchhi/jitapi"
)

func main() {
	dbconfig := jitapi.DbConfig{
		DBType:     "db_type",
		DBName:     "db_name",
		DBPort:     "db_port",
		DBHost:     "db_host",
		DBUser:     "db_user",
		DBPassword: "db_password",
		DBSslMode:  "db_sslmode",
	}

	connectionString := jitapi.Postgres{
		Config: dbconfig,
	}

	if err := connectionString.Connect(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	query := "you sql query"

	data, err := connectionString.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
