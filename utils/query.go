package utils

const (
	SELECT_MOVIE       = `SELECT * FROM movie`
	SELECT_MOVIE_BY_ID = `SELECT * FROM movie WHERE id = ?`
	INSERT_MOVIE       = `INSERT INTO movie(title, duration, image_url, synopsis) VALUES (?, ?, ?, ?)`
	UPDATE_MOVIE       = `UPDATE movie SET title = ?, duration = ?, image_url = ?, synopsis = ? WHERE id = ?`
	DELETE_MOVIE       = `DELETE FROM movie WHERE id = ?`
)
