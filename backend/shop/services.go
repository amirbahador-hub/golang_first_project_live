package shop

import (
	"context"
	"digikala/initializers"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetProductService(title string) []byte{

	var result bson.M
	fmt.Println("WE ARE CALLING PING")
	err_ping := initializers.MONGO.Ping(context.TODO(), nil)
	fmt.Println(err_ping)
	err := initializers.ProductCollection.FindOne(context.TODO(), bson.D{{"title", title}}).
		Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
	}
	jsonData, _:= json.MarshalIndent(result, "", "    ")
	return jsonData
}

func ListProductService() []interface{}{
	filter := bson.D{{}}
	cursor, err := initializers.ProductCollection.Find(context.TODO(), filter)
		
	var results []interface{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	return results
}

func CreateProductService(request ProductRequest) error{
	fmt.Println(request)
	_, err := initializers.ProductCollection.InsertOne(context.TODO(), request)
	return err 
}


func CosumeProductService(request []byte) error{
	// product := string(*request)
	var d interface{}
	if err := json.Unmarshal(request, &d); err != nil {
        panic(err)
    }
	_, err := initializers.ProductCollection.InsertOne(context.TODO(), d)
	fmt.Println("===========")
	fmt.Println(err)
	fmt.Println("===========")
	return err 
}

// func ConsumeProduct(request []byte) {
	// var d []interface{}
	// if err := json.Unmarshal(request, &d); err != nil {
        // panic(err)
    // }
// 	CreateProductService(string(request))
// }