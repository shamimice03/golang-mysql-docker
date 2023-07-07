package main

import "fmt"

func deleteAlbum(id int64) (int64, error) {
	result, err := db.Exec("DELETE FROM album WHERE id = ?", id)
	if err != nil {
		return 0, fmt.Errorf("deleteAlbum: %v", err)
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return 0, fmt.Errorf("delete album %v", err)
	}
	return rowAffected, nil
}
