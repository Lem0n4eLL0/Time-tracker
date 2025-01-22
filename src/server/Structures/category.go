package structures

type Category struct {
	Id          int    `db:"category_id" json:"id"`
	Name        string `db:"category_name" json:"name"`
	Description any    `db:"description" json:"description"`
}
