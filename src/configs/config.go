package configs

import (
	"os"
)

var DGN string
var APIHost string
var MongoURI string
var DBName string

func init() {
	DGN = "local"
	APIHost = "localhost:5000"
	MongoURI = "mongodb://localhost:27017"
	DBName = "rakshit-dev"
	env := os.Getenv("DEPLOYMENT_GROUP_NAME")
	if env != "" {
		DGN = env
	}
	env = os.Getenv("API_HOST")
	if env != "" {
		APIHost = env
	}
	env = os.Getenv("MONGO_URI")
	if env != "" {
		MongoURI = env
	}
	env = os.Getenv("MONGO_DB_NAME")
	if env != "" {
		DBName = env
	}
}
