package customer

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Customer struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}
