package models

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	//Price       float64   `json:"price"`
	//CreatedOn   time.Time `json:"created_on"`
	//UpdatedOn   time.Time `json:"updated_on"`
}
