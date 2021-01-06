package datasource

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)


type mgoDB struct {
	DB mongo.Database
}

func (mDB *mgoDB) GetTable(collectionName string) Table {
	return &mgoCollection{
		col: mDB.DB.Collection(collectionName),
	}
}

type mgoCollection struct {
	col *mongo.Collection
}


func (mCol *mgoCollection)Insert(model interface{}) string{
	res,err := mCol.col.InsertOne(context.Background(),model)
	if err != nil {
		log.Fatal(err)
	}
	return res.InsertedID.(primitive.ObjectID).Hex()
}

func (mCol *mgoCollection)Query(opt map[string]interface{}) []interface{}{
	cursor, err := mCol.col.Find(context.Background(), opt)
	if err != nil {
		log.Fatal(err)
	}
	var models []interface{}
	cursor.All(context.Background(), &models)
	return models
}

func (mCol *mgoCollection)Update(opt map[string]interface{},des map[string]interface{}){
	mCol.col.UpdateMany(context.Background(),
		opt,bson.D{
			{"$set", des},
		})
}

func (mCol *mgoCollection)Delete(opt map[string]interface{}) []interface{}{
	models := mCol.Query(opt)
	mCol.col.DeleteMany(context.Background(),opt)
	return models
}

func NewMongoDB(dbName string) DataBase{

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		log.Fatal(err)
	}
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect to mongodb successfully!")

	return &mgoDB{
		DB: *client.Database(dbName),
	}
}