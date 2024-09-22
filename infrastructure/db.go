package main

import (
    "fmt"
    "kitashiruAPI/db"
)

func main() {
    dbConn := db.NewDB()
    defer fmt.Println("Successfully Connected to Database")
    defer db.CloseDB(dbConn)
}