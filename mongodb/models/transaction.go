package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type InnerTransaction struct{
	Date time.Time `json:"date" bson:"date,required"`
	Amount int `json: "amount" bson:"amount,required"`
	TransactionCode string `json:"transaction_code" bson:"transaction_code,required"`
	Symbol string `json: "symbol" bson:"symbol,required"`
	Price string `json: "price" bson:"price,required"`
	Total string `json:"total" bson:"total,required"`

}
type Transaction struct{
	ID primitive.ObjectID`json:"id" bson:"_id"`
	AccountID  int `json:"account_id" bson:"account_id,required"`
	TransactionCount int `json:"transaction_count" bson:"transaction_count,required"`
	Bucketstartdate time.Time `json:"bucket_start_date" bson:"bucket_start_date,required"`
	Bucketenddate time.Time `json:"bucket_end_date" bson:"bucket_end_date,required"`
    //Transactions []InnerTransaction `json:"transactions" bson:"transactions,required"`
}

