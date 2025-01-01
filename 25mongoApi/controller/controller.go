package controller

import (
	"context"
	"fmt"
	"mongoapi/model"
	"mongoapi/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*

*/

/*

*/

const connectionString = ""
const dbName = "netflix"
const colName = "watchlist"

// Most importent!

var collection *mongo.Collection

// connect with mongoDB

func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	utils.CheckTheError(err)
	fmt.Println("MongoDB connection success..!")

	collection = client.Database(dbName).Collection(colName)
	// collection instance
	fmt.Println("Collection instance is ready..!")
}

// 	MONGODEB helpler - file's

// insert 1 record
func insertOneMovie(movie model.Netflix) {
	insertOne, err := collection.InsertOne(context.Background(), movie)
	utils.CheckTheError(err)

	fmt.Println("Insert 1 movie in db with id: ", insertOne.InsertedID)
}

// update 1 record
func updateOneMove(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	utils.CheckTheError(err)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	utils.CheckTheError(err)

	fmt.Println("Modified count: ", updateResult.ModifiedCount)
}

// delete 1 record
func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	utils.CheckTheError(err)
	fillter := bson.M{"_id": id}
	deleteResult, err := collection.DeleteOne(context.Background(), fillter, nil)
	utils.CheckTheError(err)
	fmt.Println("Deleted count: ", deleteResult.DeletedCount)
}

// delete all record in mongoDB
func deleteAllMovie() {
	// fillter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	utils.CheckTheError(err)
	fmt.Println("Deleted count: ", deleteResult.DeletedCount)
}

// get all record in mongoDB
func getAllMovie() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	utils.CheckTheError(err)

	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		utils.CheckTheError(err)
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}
