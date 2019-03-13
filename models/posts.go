package models

//Post struct
type Post struct {
	Title     string `db:"title"`
	Author    string `db:"author"`
	Publisher string `db:"publisher"`
	Content   string `db:"content"`
}

func (db *Database) createPost(post Post) Post {
	return post
}
