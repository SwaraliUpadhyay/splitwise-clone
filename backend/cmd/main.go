package main

import (
    "log"
    "splitwise-clone/db"
    "splitwise-clone/server"

    "github.com/joho/godotenv"
)

func main() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Connect to MongoDB
    db.Connect()

    // Start server
    server.Start()
}