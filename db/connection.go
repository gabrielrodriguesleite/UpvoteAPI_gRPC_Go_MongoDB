package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB_USER = ""        // os.Getenv("DB_USER")
var DB_PASS = ""        // os.Getenv("DB_PASS")
var DB_HOST = ""        // os.Getenv("DB_HOST")
var DB_PORT = ""        // os.Getenv("DB_PORT")
var DB_CLUSTER_URI = "" // os.Getenv("DB_CLUSTER")
var MONGO_URI = ""

func configMongoURI() {
	gotenv.Load()
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_CLUSTER_URI = os.Getenv("DB_CLUSTER_URI")

	if DB_USER == "" {
		DB_USER = "root"
	}
	if DB_PASS == "" {
		DB_PASS = "12345678"
	}
	if DB_CLUSTER_URI != "" {
		DB_CLUSTER_URI = fmt.Sprintf("%s:%s", DB_HOST, DB_PORT)
		MONGO_URI = fmt.Sprintf("mongodb+srv://%s:%s@%s", DB_USER, DB_PASS, DB_CLUSTER_URI)
	} else {
		if DB_HOST == "" {
			DB_HOST = "localhost"
		}
		if DB_PORT == "" {
			DB_PORT = "27017"
		}
		MONGO_URI = fmt.Sprintf("mongodb://%s:%s@%s:%s", DB_USER, DB_PASS, DB_HOST, DB_PORT)
		print(MONGO_URI)
	}
}

// baseado no exemplo do cloud.mongodb na versão do driver 1.6+
func getConnection() (client *mongo.Client, ctx context.Context) {
	configMongoURI()

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(MONGO_URI).SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return
}
