package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func init_table(db *sql.DB) (err error) {
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS orders (oid INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT, cid INT UNSIGNED, price FLOAT);")
	if err != nil {
		return
	}

	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS customer (cid INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT, name VARCHAR(255), gender ENUM ('Male', 'Female') NOT NULL)")
	if err != nil {
		return
	}
	return
}

func init_data(db *sql.DB) (err error) {
	sqls := []string{
		"INSERT INTO customer (name, gender) value ('Ben','Male');",
		"INSERT INTO customer (name, gender) value ('Alice','Female');",
		"INSERT INTO customer (name, gender) value ('Peter','Male');",
		"INSERT INTO orders (cid, price) value (1,10.23);",
		"INSERT INTO orders (cid, price) value (2,122);",
		"INSERT INTO orders (cid, price) value (2,72.5);",
	}
	for _, sql := range sqls {
		_, err = db.Exec(sql)
		if err != nil {
			return
		}
	}

	return
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:4000)/test_go?charset=utf8mb4")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := init_table(db); err != nil {
		panic(err)
	}
	if err := init_data(db); err != nil {
		panic(err)
	}

	// update
	_, err = db.Exec("UPDATE orders SET price = price + 1 WHERE oid = 1")
	if err != nil {
		panic(err)
	}
	// delete
	_, err = db.Exec("DELETE FROM orders WHERE oid = 1")
	if err != nil {
		panic(err)
	}
	// read
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var oid, cid int
		var price float64
		err := rows.Scan(&oid, &cid, &price)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d %d %.2f\n", oid, cid, price)
	}
	// join
	rows, err = db.Query("SELECT customer.name, orders.price FROM customer, orders WHERE customer.cid = orders.cid and orders.cid = 1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var price float64
		err := rows.Scan(&name, &price)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %.2f\n", name, price)
	}
}
