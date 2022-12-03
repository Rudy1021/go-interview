package collections

import (
	"context"
	"encoding/json"
	"fmt"
	"go-interview/api/database"
	"go-interview/api/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UserText_c() {
	var c = database.SelectCollection()
	res, err := c.InsertOne(context.TODO(), bson.D{{"user_id", "user123"}, {"entertext", "test"}})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.InsertedID)
}
func UserText_r_all() {
	var c = database.SelectCollection()
	findOptions := options.Find()
	findOptions.SetLimit(2)

	// Here's an array in which you can store the decoded documents
	var results []*models.UserText

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := c.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem models.UserText
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
}

func UserText_r_one() {
	var c = database.SelectCollection()
	var result models.UserText
	err := c.FindOne(context.TODO(), bson.D{{"user_id", "user123"}}).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	output, err := json.MarshalIndent(result, "", "    ")
	fmt.Printf("%s\n", output)
}

func UserText_u() {

}
func UserText_d() {

}
