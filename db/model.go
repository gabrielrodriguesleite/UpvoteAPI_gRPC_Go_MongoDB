package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func Insert(collection string, data any) error {

	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database("upvote").Collection(collection)

	_, err := c.InsertOne(context.Background(), data)

	return err

}

func List(collection string) (list []Vote, err error) {

	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database("upvote").Collection(collection)

	cur, err := c.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Println(err)
	}

	err = cur.All(context.Background(), &list)

	return

}
