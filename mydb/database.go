package mydb

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)
const (
	host     = "localhost"
	port     = 5432
	user     = "citizix_user"
	password = "S3cret"
	dbname   = "citizix_db"
  )

var globalDatabase *sql.DB

func GetDatabase() *sql.DB{
	return globalDatabase
}

func ConnectDatabase() *sql.DB{

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
  		panic(err)
	}
	err = db.Ping()
	if err != nil {
 	 panic(err)
	}

	globalDatabase = db
	return db
}