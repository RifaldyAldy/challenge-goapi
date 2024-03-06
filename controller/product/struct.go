package product

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Unit  string `json:"unit"`
}
