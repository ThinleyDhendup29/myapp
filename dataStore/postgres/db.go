package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// db details
const (
	postgres_host     = "dpg-cp5maei1hbls73fi6aag-a.singapore-postgres.render.com"
	postgres_port     = 5432
	postgres_user     = "postgres_admin"
	postgres_password = "CtzYEL9RdVK3EFk3xyRYUyKsSQbn0wce"
	postgres_dbname   = "my_db_c7dx"
)

// create pointer variable Db which points to sql driver
var Db *sql.DB

// init() is always called before main()
func init() {
	//creating the connection string
	db_info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname)

	var err error
	//open connection to database
	Db, err = sql.Open("postgres", db_info)

	if err != nil {
		panic(err)

	} else {
		log.Println("Database successfully configured")
	}
}
