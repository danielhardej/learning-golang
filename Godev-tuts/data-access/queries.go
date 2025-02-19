package main

import (
	"database/sql"
	"fmt"
	"log"
)

func showAllAlbums() {
	allAlbums, err := db.Query("SELECT * FROM album")
	if err != nil {
		log.Fatal(err)
	}
	for allAlbums.Next() {
		var id int
		var title string
		var artist string
		var price float32
		err = allAlbums.Scan(&id, &title, &artist, &price)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %q (%s) %.2f\n", id, title, artist, price)
	}
	defer allAlbums.Close()
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
	// check if album already exists
	var existingAlbum Album
	err := db.QueryRow("SELECT id, title, artist, price FROM album WHERE title = ? AND artist = ?", alb.Title, alb.Artist).
		Scan(&existingAlbum.ID, &existingAlbum.Title, &existingAlbum.Artist, &existingAlbum.Price)
	if err == nil {
		return 0, fmt.Errorf("addAlbum: album already exists: %s %s", existingAlbum.Artist, existingAlbum.Title)
	} else if err != sql.ErrNoRows {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
