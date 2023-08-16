package main

import (
	//"context"
	"fmt"
	//"mongodb/constants"
	// "mongodb/models"
	"mongodb/services"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	// "honnef.co/go/tools/config"
)
	
var (//mongoclient will act as a global variable
    mongoClient *mongo.Client
)
func main(){
    fmt.Println("MongoDB successfully connected...")
    // products:=[]interface{}{
    //     models.Product{ID:primitive.NewObjectID(),Name:"Oneplus",Price:100000,Description:"Budget Phone"},
    //     models.Product{ID:primitive.NewObjectID(),Name:"Vivo",Price:100000,Description:"China Phone"} }
    // services.InsertProductlist(products)
    //client:=config.ConnectDatabase()
    //collection:=config
	products,_:=services.FindRes()
    for _,product:=range products{
        fmt.Println(product.Name)
    }
}

