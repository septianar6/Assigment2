package database

import (
	"assigment2/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func Start() (Database, error) {
	var host = "localhost"
	var port = 5432
	var username = "postgres"
	var password = "060606"
	var dbName = "orders_by"

	var conn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbName)

	db, err := gorm.Open(postgres.Open(conn))
	if err != nil {
		fmt.Println("error open connection to db", err)
		return Database{}, err
	}

	err = db.Debug().AutoMigrate(model.Order{})
	if err != nil {
		fmt.Println("error on migration", err)
		return Database{}, err
	}

	return Database{
		db: db,
	}, nil
}

func (d Database) GetOrders() ([]model.Order, error) {
	dbg := d.db.Find(&[]model.Order{})

	rows, err := dbg.Rows()
	if err != nil {
		return nil, err
	}

	orders := make([]model.Order, 0)

	for rows.Next() {
		var orders model.Order

		err = rows.Scan(&orders.Order_ID, &orders.Customer_Name, &orders.Ordered_at)
		if err != nil {
			continue
		}

		//orders = append(orders, order)
	}

	return orders, nil
}

func (d Database) CreateOrder(order model.Order) (model.Order, error) {
	dbg := d.db.Create(&order)

	row := dbg.Row()

	var orderResult model.Order

	err := row.Scan(&orderResult.Order_ID, &orderResult.Customer_Name, &orderResult.Ordered_at)

	return orderResult, err
}
