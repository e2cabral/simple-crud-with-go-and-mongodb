package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"crud-with-golang-and-mongodb/src/database"
	"crud-with-golang-and-mongodb/src/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection = database.Connect().Database("golang-tests").Collection("users")

// CreateProfile is responsible to create a new doc in mongodb
func CreateProfile(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var person models.User

	err := json.NewDecoder(req.Body).Decode(&person)

	if err != nil {
		fmt.Print(err)
	}

	result, err := userCollection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(res).Encode(result.InsertedID)
}

// GetUserProfile return all the users profile in mongodb
func GetUserProfile(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var person models.User

	err := json.NewDecoder(req.Body).Decode(&person)

	if err != nil {
		fmt.Print(err)
	}

	var result primitive.M

	newErr := userCollection.FindOne(context.TODO(), bson.D{{"name", person.Name}}).Decode(&result)

	if newErr != nil {
		fmt.Print(newErr)
	}

	json.NewEncoder(res).Encode(result)
}

// UpdateUserProfile this function updates an user profile
func UpdateUserProfile(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	type updateBody struct {
		Name string `json:"name"`
		City string `json:"city"`
	}

	var body updateBody

	err := json.NewDecoder(req.Body).Decode(&body)

	if err != nil {
		fmt.Print(err)
	}

	filter := bson.D{{"name", body.Name}}
	after := options.After

	returnOptions := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	update := bson.D{{"$set", bson.D{{"city", body.City}}}}

	updateResult := userCollection.FindOneAndUpdate(context.TODO(), filter, update, &returnOptions)

	var result primitive.M

	_ = updateResult.Decode(&result)

	json.NewEncoder(res).Encode(result)
}

// DeleteUserProfile this function deletes an user profile from mongodb
func DeleteUserProfile(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)["id"]

	_id, err := primitive.ObjectIDFromHex(params)

	if err != nil {
		fmt.Print(err)
	}

	opts := options.Delete().SetCollation(&options.Collation{})

	response, newErr := userCollection.DeleteOne(context.TODO(), bson.D{{"_id", _id}}, opts)

	if newErr != nil {
		log.Fatal(err)
	}

	json.NewEncoder(res).Encode(response.DeletedCount)
}

// GetAllUsersProfile this function returns all users profile registred
func GetAllUsersProfile(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var results []*models.User

	findOptions := options.Find()

	cur, err := userCollection.Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var user models.User
		err := cur.Decode(&user)

		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &user)
	}

	json.NewEncoder(res).Encode(results)
}
