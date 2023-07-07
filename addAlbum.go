package main

import "fmt"

func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title,artist,price) VALUES (?,?,?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("addAlbum %v", err)
	}

	return id, nil
}
