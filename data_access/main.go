package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	// Get a database handle.
	db, err := sql.Open("sqlite3", "../.data/recordings.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums found: %v\n", albums)

	// Hard-code ID 2 here to test the query.
	alb, err := albumByID(2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album found: %v\n", alb)

	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID of added album: %v\n", albID)

	rowsAffected, err := removeAlbum(albID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Rows affected of removed album: %v\n", rowsAffected)
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	// Get a database handle.
	db, err := sql.Open("sqlite3", "../.data/recordings.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

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
	// Get a database handle.
	db, err := sql.Open("sqlite3", "../.data/recordings.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

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
	// Get a database handle.
	db, err := sql.Open("sqlite3", "../.data/recordings.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

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

func removeAlbum(albID int64) (int64, error) {
	db, err := sql.Open("sqlite3", "../.data/recordings.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	result, err := db.Exec("DELETE FROM album WHERE id = ?", albID)

	if err != nil {
		return 0, fmt.Errorf("removeAlbum: %v", err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, fmt.Errorf("removeAlbum: %v", err)
	}

	return rowsAffected, nil
}
