package main

import (
	"context"
	"fmt"
	"os"

	"benni/db"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	fmt.Println("Hello World!")

	db.InitDB()
	conn = db.GetDB()

	defer conn.Close(context.Background())

	var err error
	var id int
	var artist string
	var title string
	var price float32
	err = conn.QueryRow(context.Background(), "select * from albums").Scan(&id, &title, &artist, &price)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%d %v %v %v\n", id, title, artist, price)

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		fmt.Println("%v", err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	addAlbum(Album{
		Title:  "Bennis Album",
		Artist: "Benni",
		Price:  99.99,
	})

	fmt.Printf("%d %v %v %v\n", id, title, artist, price)

	albums, err = albumsByArtist("Benni")
	if err != nil {
		fmt.Println("%v", err)
	}
	fmt.Printf("Albums found: %v\n", albums)

}

func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	url := "postgres://postgres:postgres@localhost:5432/records"
	conn, err := pgx.Connect(context.Background(), url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	rows, err := conn.Query(context.Background(), "SELECT * FROM albums WHERE artist LIKE $1", name)
	defer conn.Close(context.Background())

	if err != nil {
		return nil, fmt.Errorf("query failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("query failed: %v", err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

func addAlbum(alb Album) (int64, error) {
	url := "postgres://postgres:postgres@localhost:5432/records"
	conn, err := pgx.Connect(context.Background(), url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	result, err := conn.Exec(context.Background(), "INSERT INTO albums (title, artist, price) VALUES ($1, $2, $3)", alb.Title, alb.Artist, alb.Price)
	defer conn.Close(context.Background())

	if err != nil {
		fmt.Fprintf(os.Stderr, "addAlbum: %v\n", err)
		os.Exit(1)
	}

	return result.RowsAffected(), err
}
