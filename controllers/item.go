package controllers

import (
	"assignment/config"
	"assignment/structs"
	"fmt"
)

func GetItem(order_id int) []structs.Item {
	db := config.DbInit()
	defer db.Close()
	var items []structs.Item
	sqlStatement := fmt.Sprintf(`SELECT * FROM items WHERE order_id=%d`, order_id)
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var item structs.Item
		err = rows.Scan(&item.Item_id, &item.Item_code, &item.Description, &item.Quantity, &item.Order_id)
		if err != nil {
			panic(err)
		}
		items = append(items, item)
	}
	return items
}

func CreateItem(item structs.Item, id int64) {
	db := config.DbInit()
	defer db.Close()
	sqlStatement := fmt.Sprintf(`INSERT INTO items (item_code, description,quantity,order_id)
	VALUES('%s','%s',%d,%d);`, item.Item_code, item.Description, item.Quantity, id)
	fmt.Println(sqlStatement)
	db.QueryRow(sqlStatement)
}

func UpdateItem(item structs.Item, id int) {
	db := config.DbInit()
	defer db.Close()
	sqlStatement := fmt.Sprintf(`UPDATE items SET item_code='%s', description='%s',quantity=%d WHERE order_id=%d;`, item.Item_code, item.Description, item.Quantity, id)
	fmt.Println(sqlStatement)
	row, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	count, err := row.RowsAffected()
	if count <= 0 {
		CreateItem(item, int64(id))
	}
}

func DeleteItem(id int) {
	db := config.DbInit()
	defer db.Close()
	sqlStatement := `DELETE from items WHERE order_id=$1`
	db.Exec(sqlStatement, id)
}
