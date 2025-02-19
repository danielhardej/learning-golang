package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	fmt.Println("Show all albums...")
	showAllAlbums()

	fmt.Println("Getting all albums by John Coltrane...")
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	for _, alb := range albums {
		fmt.Printf("%d: %q (%s) %.2f\n", alb.ID, alb.Title, alb.Artist, alb.Price)
	}

	// Query for a single row
	// Hard-code ID 2 here to test the query.
	fmt.Println("Getting album by ID=2...")
	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d: %q (%s) %.2f\n", alb.ID, alb.Title, alb.Artist, alb.Price)

	// Add data...
	fmt.Println("Adding album...")
	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Print(err)
	} else {
		fmt.Println("New record added with ID:", albID)
	}
}
