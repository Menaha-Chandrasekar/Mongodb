package services

import (
	"context"
	"fmt"
	"mongodb/config"
	"mongodb/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Product() *mongo.Collection{
	client,_:=config.ConnectDataBase()
	return config.GetCollection(client,"sample_analytics","transactions")
	
}
func FindTransaction()([]*models.Transaction,error){
	ctx, _:=context.WithTimeout(context.Background(),100*time.Second)
	//filter:=bson.D{} # D for Document if query is plane
	filter:=bson.M{"transaction_count":bson.D{{"$gt",85},{"$lt",90}}}//M is for Map, query will be having nested operators
	options:=options.Find().SetSort(bson.D{{"transaction_count",1}}).SetSkip(30).SetLimit(10)
	result,err:=Product().Find(ctx,filter,options)
	if err!=nil{
		fmt.Println(err.Error())
		return nil,err
	}else{
		//do something
		//bulidthe array of products for the cursor that we received
		var products []*models.Transaction

	for result.Next(ctx){
		product:=&models.Transaction{}
		err:=result.Decode(product)

		if err!=nil{
			return nil,err
		}
		products=append(products,product)
	}
	if err:=result.Err();err!=nil{
		return nil,err
	}
	if len(products)==0{
		return []*models.Transaction{},nil
	}
	return products,nil
	}
}
