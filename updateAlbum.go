package main

import "fmt"

func updateAlbum(id int64, album Album) (int64, error) {
	result, err := db.Exec("UPDATE album SET title=?, artist=?, price=? WHERE id=?", album.Title, album.Artist, album.Price, id)
	if err != nil {
		return 0, fmt.Errorf("updateAlbum %v", err)
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return 0, fmt.Errorf("updateAlbum %v", err)
	}
	return rowAffected, nil
}
