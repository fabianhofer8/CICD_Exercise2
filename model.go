package main

import (
	"database/sql"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *product) getProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM products WHERE id=$1",
		p.ID).Scan(&p.Name, &p.Price)
}

func (p *product) updateProduct(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3",
			p.Name, p.Price, p.ID)

	return err
}

func (p *product) deleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)

	return err
}

func (p *product) createProduct(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO products(name, price) VALUES($1, $2) RETURNING id",
		p.Name, p.Price).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	rows, err := db.Query(
		"SELECT id, name,  price FROM products LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func getProductsUnder(db *sql.DB, value int) ([]product, error) {
	rows, err := db.Query(
		"SELECT id, name,  price FROM products where price < $1",
		value)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := []product{}
	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
func getProductsOver(db *sql.DB, value int) ([]product, error) {
	rows, err := db.Query(
		"SELECT id, name,  price FROM products where price > $1",
		value)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := []product{}
	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func getMin(db *sql.DB) (product, error) {
	row := db.QueryRow(
		"SELECT name, price FROM products ORDER BY price ASC limit 1",
	)
	var p product
	err := row.Scan(&p.Name, &p.Price)
	if err != nil {
		return product{}, err
	}
	println(p.Name)

	return p, nil
}
func getMax(db *sql.DB) (product, error) {
	row := db.QueryRow(
		"SELECT name, price " +
			"FROM products ORDER BY price DESC limit 1",
	)
	var p product
	err := row.Scan(&p.Name, &p.Price)
	if err != nil {
		return product{}, err
	}
	println(p.Name)
	return p, nil
}
func getProductsLike(db *sql.DB, value string) ([]product, error) {
	rows, err := db.Query(
		"SELECT id, name, price FROM products WHERE "+
			"name LIKE '%' || $1 || '%'",
		value)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := []product{}
	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
