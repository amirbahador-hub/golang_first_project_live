package apis

import (
	"net/http"
	"digikala/models"
	"fmt"
	"encoding/json"
	"digikala/mydb"
)



func GetRoot(w http.ResponseWriter, r *http.Request){
	
	var products []models.Product
	globalDatabase := mydb.GetDatabase()
	rows, err := globalDatabase.Query("select id, title, price from Product")
	fmt.Println(err)
	defer rows.Close()
	for rows.Next() {
		product1 := models.Product{}
		err := rows.Scan(&product1.Id, &product1.Title, &product1.Price)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(product1)
		products = append(products, product1)
	}
	err = rows.Err()
	fmt.Println(err)
	fmt.Println("GET ROOT")
	fmt.Println(rows)
	response, _ := json.Marshal(products)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	// io.WriteString(w, string(response))
}

func CreateProduct(w http.ResponseWriter, r *http.Request){
	body := json.NewDecoder(r.Body)
	p := &models.Product{}
	err := body.Decode(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	globalDatabase := mydb.GetDatabase()
	rows, err := globalDatabase.Query(`insert into Product VALUES ($1, $2, $3)`, p.Id, p.Title, p.Price)
	fmt.Println(err)
	fmt.Println(rows)
}

