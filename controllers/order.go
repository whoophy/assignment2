package controllers

import (
	"assignment/config"
	"assignment/structs"
	"fmt"
)

func GetOrders() ([]structs.Order, error) {
	db := config.DbInit()
	defer db.Close()
	var orders []structs.Order
	sqlStatement := `SELECT * from orders`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var order structs.Order
		err = rows.Scan(&order.OrderID, &order.CostumerName, &order.OrderAt)
		items := GetItem(order.OrderID)
		order.Items = items
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}
	return orders, err
}

func DeleteOrder(order_id int) (string, error) {
	db := config.DbInit()
	defer db.Close()
	DeleteItem(order_id)
	sqlStatement := `DELETE from orders WHERE order_id=$1`
	row, err := db.Exec(sqlStatement, order_id)
	if err != nil {
		panic(err)
	}
	count, err := row.RowsAffected()
	if err != nil {
		panic(err)
	}
	if count <= 0 {
		return fmt.Sprintf("tidak dapat menemukan data dengan id %v", order_id), err
	}
	fmt.Printf("jumlah data yang ditemukan %v", count)
	return fmt.Sprintf("Data dengan id %v berhasil terhapus", order_id), err
}

func CreateOrder(order structs.Order) string {
	db := config.DbInit()
	defer db.Close()
	sqlStatement := `
	INSERT INTO orders (costumer_name, ordered_at) 
	VALUES($1,$2) RETURNING order_id`
	var id int64
	err := db.QueryRow(sqlStatement, order.CostumerName, order.OrderAt).Scan(&id)
	for _, v := range order.Items {
		CreateItem(v, id)
	}
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Berhasil menambahkan %s dengan id - %v", order.CostumerName, id)
}

func UpdateOrder(order structs.Order, id int) string {
	db := config.DbInit()
	defer db.Close()
	sqlStatement := `UPDATE orders SET costumer_name=$1, ordered_at=$2 WHERE order_id=$3`
	row, err := db.Exec(sqlStatement, order.CostumerName, order.OrderAt, id)
	if err != nil {
		panic(err)
	}
	count, err := row.RowsAffected()
	if err != nil {
		panic(err)
	}
	if count <= 0 {
		return fmt.Sprintf("Data dengan order id %v Tidak ditemukan", id)
	}
	for _, v := range order.Items {
		fmt.Println(v)
		UpdateItem(v, id)
	}
	return fmt.Sprintf("Data dengan order id %v berhasil terupdate", id)
}
