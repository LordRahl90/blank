package models

//Comment struct
type Comment struct {
	ID      int    `db:"id"`
	PostID  int    `db:"post_id"`
	Comment string `db:"comment"`
}

//CreateComment function
func (db *Database) CreateComment(c Comment) (int, error) {
	return 1, nil
}
