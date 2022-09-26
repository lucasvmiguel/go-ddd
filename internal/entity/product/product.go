package product

type Product struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}
