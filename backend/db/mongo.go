package db

import (
    "context"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() {
    uri := os.Getenv("MONGO_URI")

    // Set a 10 second timeout — don't wait forever if DB is unreachable
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal("MongoDB connection failed:", err)
    }

    // Ping to confirm connection actually works
    if err = client.Ping(ctx, nil); err != nil {
        log.Fatal("MongoDB ping failed:", err)
    }

    Client = client
    log.Println("Connected to MongoDB!")
}

// Helper to get a collection cleanly anywhere in the app
func GetCollection(name string) *mongo.Collection {
    return Client.Database(os.Getenv("DB_NAME")).Collection(name)
}