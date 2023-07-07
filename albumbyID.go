package main

import (
	"database/sql"
	"fmt"
)

func albumByID(id int64) (Album, error) {
	var alb Album
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("album by id %d: no such album", id)
		}
		return alb, fmt.Errorf("albumById %d: %v", id, err)
	}

	return alb, nil
}
