package database

type Customer struct {
	ID       int64
	Name     string
	LastName string
	Email    string
}

type Customers []Customer

func CreateCustomer(c Customer) Customer {
	var id int
	db := GetConnection()
	row := db.QueryRow("INSERT INTO customer (name, last_name, email) values ($1, $2, $3) RETURNING id", c.Name, c.LastName, c.Email)
	err := row.Scan(&id)
	if err != nil {
		panic(err.Error())
	}
	return *RetrieveCustomerById(int(id))
}

func UpdateCustomer(c Customer) Customer {
	db := GetConnection()
	_, err := db.Exec("UPDATE customer set name=$1, last_name=$2, email=$3 WHERE id=$4", c.Name, c.LastName, c.Email, c.ID)
	if err != nil {
		panic(err.Error())
	}
	return c
}

func RetrieveCustomers() Customers {
	db := GetConnection()
	rows, err := db.Query("SELECT id, name, last_name, email FROM customer")
	if err != nil {
		panic(err.Error())
	}

	customers := Customers{}

	for rows.Next() {
		customer := Customer{}
		err = rows.Scan(&customer.ID, &customer.Name, &customer.LastName, &customer.Email)
		if err != nil {
			panic(err.Error())
		}
		customers = append(customers, customer)
	}

	return customers
}

func RetrieveCustomerById(id int) *Customer {
	db := GetConnection()
	row := db.QueryRow("SELECT id, name, last_name, email FROM customer WHERE id=$1", id)
	customer := Customer{}
	err := row.Scan(&customer.ID, &customer.Name, &customer.LastName, &customer.Email)
	if err != nil {
		return nil //When id not exists
	}
	return &customer
}

func DeleteCustomer(id int) {
	db := GetConnection()
	_, err := db.Exec("DELETE FROM customer WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}
}
