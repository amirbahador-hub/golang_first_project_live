package main

import (
	"digikala/mydb"
	"digikala/logger"
	"digikala/migrations"
	"digikala/apis"
	"net/http"
)


func main(){

	myslog := logger.GetLogger()
	http.HandleFunc("/", apis.GetRoot)
	http.HandleFunc("/create_product", apis.CreateProduct)



	db_connection := mydb.ConnectDatabase()
	defer db_connection.Close()

	migrations.CreateTables()
	myslog.Info("Start Listening to 8000")
	serverErr := http.ListenAndServe(":8000", nil)
	if serverErr != nil {
		myslog.Error("Something is Wrong")
	}
}