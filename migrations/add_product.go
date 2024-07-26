package migrations

import (
	"fmt"
	"digikala/logger"
	"digikala/mydb"
	_ "github.com/lib/pq"
)


func CreateTables(){

	globalDatabase := mydb.ConnectDatabase()
	myslog := logger.GetLogger()

	myslog.Info("create tables")
	rows, err := globalDatabase.Query(`create table if not exists Product (
  id INT PRIMARY KEY,
  title TEXT,
  price INT
)`)
	if err != nil {
		myslog.Error(err.Error())
	}
	fmt.Println(rows)
}