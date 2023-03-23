package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM users"

	name := r.URL.Query()["name"]
	age := r.URL.Query()["age"]

	if name != nil {
		fmt.Println(name[0])
		query += " WHERE name= '" + name[0] + " ', "
	}

	if age != nil {
		if name != nil {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " age= '" + age[0] + "' "
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again.")
		return
	}

	var user User
	var users []User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.userType); err != nil {
			log.Println(err)
			sendErrorResponse(w, "Something went wrong, please try again.")
			return
		} else {
			users = append(users, user)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	var response UsersResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = users
	json.NewEncoder(w).Encode(response)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM Products"
	name := r.URL.Query()["name"]

	if name != nil {
		fmt.Println(name[0])
		query += " WHERE name= '" + name[0] + " ', "
	}
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	var product Product
	var products []Product
	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			log.Println(err)
			sendErrorResponse(w, "Something went wrong, please try again.")
			return
		} else {
			products = append(products, product)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	var response ProductsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = products
	json.NewEncoder(w).Encode(response)
}

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM Transactions"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	var transaction Transaction
	var transactions []Transaction
	for rows.Next() {
		if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.ProductID, &transaction.Quantity); err != nil {
			log.Println(err)
			sendErrorResponse(w, "Something went wrong, please try again.")
			return
		} else {
			transactions = append(transactions, transaction)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	var response TransactionsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = transactions
	json.NewEncoder(w).Encode(response)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")
	user_type, _ := strconv.Atoi(r.Form.Get("user_type"))
	_, errQuery := db.Exec("INSERT INTO users(name, age, address, user_type) values (?,?,?,?)",
		name,
		age,
		address,
		user_type,
	)

	var response UserResponse
	if errQuery == nil {
		response.Status = 400
		response.Message = "Success"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "insert Failed!"

	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	name := r.Form.Get("name")
	Price, _ := strconv.Atoi(r.Form.Get("price"))

	_, errQuery := db.Exec("INSERT INTO users(name, price) values (?,?)",
		name,
		Price,
	)

	var response UserResponse
	if errQuery == nil {
		response.Status = 400
		response.Message = "Success"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "insert Failed!"

	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func InsertTransaction(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	UserID, _ := strconv.Atoi(r.Form.Get("User ID"))
	ProductID, _ := strconv.Atoi(r.Form.Get("ProductID"))
	Quantity, _ := strconv.Atoi(r.Form.Get("Quantity"))
	_, errQuery := db.Exec("INSERT INTO users(UserID, ProductID, Quantity) values (?,?,?)",
		UserID,
		ProductID,
		Quantity,
	)

	var response UserResponse
	if errQuery == nil {
		response.Status = 400
		response.Message = "Success"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "insert Failed!"

	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}

	vars := mux.Vars(r)
	userId := vars["user_id"]

	_, errQuery := db.Exec("DELETE FROM users WHERE id=?",
		userId,
	)

	var response UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Delete Failed"

	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func DeleteProducts(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}

	vars := mux.Vars(r)
	product_iD := vars["Product ID"]

	_, errQuery := db.Exec("DELETE FROM products WHERE id=?",
		product_iD,
	)

	var response UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Delete Failed"
		response.Status = 400
		response.Message = "Delete Failed"

	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func DeleteTransactions(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}

	vars := mux.Vars(r)
	Transaction_ID := vars["Transaction ID"]

	_, errQuery := db.Exec("DELETE FROM transactions WHERE id=?",
		Transaction_ID,
	)

	var response UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Delete Failed"
		response.Status = 400
		response.Message = "Delete Failed"

	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}

	vars := mux.Vars(r)
	userId := vars["user_id"]

	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")
	user_type, _ := strconv.Atoi(r.Form.Get("user_type"))

	_, errQuery := db.Exec("UPDATE users SET name=?, age=?, address=?, user_type=? WHERE id=?",
		name,
		age,
		address,
		user_type,
		userId,
	)
	var response UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success Update"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Delete Failed"
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, errQuery := db.Exec("UPDATE products SET name=$2, price=$3 WHERE id=$1",
		product.ID,
		product.Name,
		product.Price,
	)
	var response UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Delete Failed"
		response.Status = 400
		response.Message = "Delete Failed"

	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateTransactions(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	var transaction Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, errQuery := db.Exec("UPDATE transactions SET UserID=$2, ProductID=$3, Quantity=$4 WHERE id=$1",
		transaction.ID,
		transaction.UserID,
		transaction.ProductID,
		transaction.UserID,
	)
	var response UserResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Delete Failed"
		response.Status = 400
		response.Message = "Delete Failed"

	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func sendErrorResponse(w http.ResponseWriter, message string) {
	var response ErrorResponse
	response.Status = 400
	response.Message = message

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(response)
}

// func GetAllUsersORM(w http.ResponseWriter, r *http.Request) {
// 	db := ConnectORM()

// 	var user []User
// 	db.Find(&user)

// 	w.Header().Set("Content-Type", "application/json")
// 	if len(user) < 5 {
// 		var response UsersResponse
// 		response.Status = 200
// 		response.Message = "Success"
// 		response.Data = user
// 		json.NewEncoder(w).Encode(response)
// 	} else {
// 		var response ErrorResponse
// 		response.Status = 400
// 		response.Message = "Error Array Size Not Correct"
// 		json.NewEncoder(w).Encode(response)
// 	}
// }
