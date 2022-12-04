package collections

import (
	"context"
	"fmt"
	"go-interview/api/database"

	"go.mongodb.org/mongo-driver/bson"
)

func UserText_c(user_id string, entertext string) {
	var c = database.SelectCollection()
	_, err := c.InsertOne(context.TODO(), bson.D{{"user_id", user_id}, {"entertext", entertext}})
	if err != nil {
		fmt.Println(err)
	}
}
func UserText_r_all() {
	//var c = database.SelectCollection()

}

func UserText_r_one(id string) string {
	var c = database.SelectCollection()
	filter := bson.D{{"user_id", id}}
	cursor, err := c.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	output := ""
	for _, result := range results {
		output = output + fmt.Sprintf("%s", result)
	}
	return output
}

func UserText_u() {

}
func UserText_d() {

}
