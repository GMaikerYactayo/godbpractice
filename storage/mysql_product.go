package storage

import (
	"database/sql"
	"fmt"
	"github.com/GMaikerYactayo/godbpractice/pkg/product"
)

const (
	mysqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		create_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`

	mysqlCreateProduct = `INSERT INTO products(name, observations, price, create_at)
		VALUES(?, ? ,? ,?)`

	mysqlGetAllProduct = `SELECT id, name, observations, price, create_at, updated_at
		FROM products`

	mysqlGetByIdProduct = mysqlGetAllProduct + ` WHERE id = ?`

	mysqlUpdateProduct = `UPDATE products 
				SET name = ?, observations = ?, price = ?, updated_at = ?
                WHERE id = ?`

	mysqlDeleteProduct = `DELETE FROM products 
				WHERE id = ?`
)

// mySQLProduct used for work with mysql - product
type mySQLProduct struct {
	db *sql.DB
}

// newMySQLProduct return a new pinter of mySQLProduct
func newMySQLProduct(db *sql.DB) *mySQLProduct {
	return &mySQLProduct{db}
}

// Migrate implement the interface product.Storage
func (p *mySQLProduct) Migrate() error {
	stmt, err := p.db.Prepare(mysqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Product migration successfully executed")
	return nil
}

// Create implement the interface product.Storage
func (p *mySQLProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(mysqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreateAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = uint(id)

	fmt.Printf("The product was created successfully with ID: %d\n", m.ID)
	return nil
}

// GetAll implement the interface product.Storage
func (p *mySQLProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(mysqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(product.Models, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ms, nil
}

// GetByID implement the interface product.Storage
func (p *mySQLProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(mysqlGetByIdProduct)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

// Update implement the interface product.Storage
func (p *mySQLProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(mysqlUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNull(m.UpdateAt),
		m.ID,
	)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("don't exist the product id: %d", m.ID)
	}
	fmt.Println("Product was updated successfully")
	return nil
}

// Delete implement the interface product.Storage
func (p *mySQLProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(mysqlDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	fmt.Println("Product was deleted successfully")
	return nil
}
