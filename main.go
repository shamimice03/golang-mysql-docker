package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// Created a db variable of *sql.DB type
// sql -> package name
// DB -> struct defined in the sql package
var db *sql.DB

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(reflect.TypeOf(db))
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected")

	// retrieve by artistName ( Mutiple Query )
	albums, err := albumByArtist("Warfezz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	// retrieve by ID ( single query)
	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	// add new album
	newAlbum := Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	}

	newAlbumId, err := addAlbum(newAlbum)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID of added album %v\n", newAlbumId)

	// delete by ID
	deleteStatus, err := deleteAlbum(3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number of Row Delete %v\n", deleteStatus)

	// update by ID
	updateAlb := Album{
		Title:  "Obak Valobasha",
		Artist: "Warfezz",
		Price:  60,
	}

	updateStatus, err := updateAlbum(4, updateAlb)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Row updated %v\n", updateStatus)

}
