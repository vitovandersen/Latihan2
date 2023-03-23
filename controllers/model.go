package controllers

type User struct {
	ID       int    `jason: "Id"`
	Name     string `jason: "name"`
	Age      int    `jason: "age"`
	Address  string `jason: "address"`
	userType int    `jason: "user_type"`
}
type Product struct {
	ID    int    `jason: "Id"`
	Name  string `jason: "name"`
	Price int    `jason: "Price"`
}

type Transaction struct {
	ID        int `jason: "Id"`
	UserID    int `jason: "UserID"`
	ProductID int `jason: "ProductID"`
	Quantity  int `jason: "Quantity"`
}

type UserResponse struct {
	Status  int    `jason: "status"`
	Message string `jason: "message"`
	Data    User   `jason: "data"`
}

type UsersResponse struct {
	Status  int    `jason: "status"`
	Message string `jason: "message"`
	Data    []User `jason: "data"`
}

type ErrorResponse struct {
	Status  int    `jason: "status"`
	Message string `jason: "message"`
	// Err error 	'jason: "error"'

}

type ProductResponse struct {
	Status  int     `jason: "status"`
	Message string  `jason: "message"`
	Data    Product `jason: "data"`
}

type ProductsResponse struct {
	Status  int       `jason: "status"`
	Message string    `jason: "message"`
	Data    []Product `jason: "data"`
}

type TransactionResponse struct {
	Status  int         `jason: "status"`
	Message string      `jason: "message"`
	Data    Transaction `jason: "data"`
}

type TransactionsResponse struct {
	Status  int           `jason: "status"`
	Message string        `jason: "message"`
	Data    []Transaction `jason: "data"`
}
