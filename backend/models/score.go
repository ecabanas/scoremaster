package models

import (
	"database/sql"
)

// InsertScore inserts a new score into the database.
func InsertScore(db *sql.DB, name string, score int) error {
	query := `INSERT INTO score (name, score) VALUES (?, ?)`
	_, err := db.Exec(query, name, score)
	if err != nil {
		return err
	}
	return nil
}

// GetScores retrieves all scores from the database.
func GetScores(db *sql.DB) ([]map[string]interface{}, error) {
	rows, err := db.Query("SELECT id, name, score FROM score")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores []map[string]interface{}
	for rows.Next() {
		var id int
		var name string
		var score int
		if err := rows.Scan(&id, &name, &score); err != nil {
			return nil, err
		}
		scores = append(scores, map[string]interface{}{
			"id":    id,
			"name":  name,
			"score": score,
		})
	}

	return scores, nil
}
