package productmodel

import (
	"errors"
	"server-side/config"
	"server-side/entities"
)

func GetAll() []entities.Product {
	rows, err := config.DB.Query(`SELECT * FROM products`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Sku, &product.Quantity, &product.Location, &product.Status); err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products
}

func Create(product entities.Product) (int64, error) {
	result, err := config.DB.Exec(`
		INSERT INTO products (name, sku, quantity, location, status)
		VALUE(?, ?, ?, ?, ?)`,
		product.Name, product.Sku, product.Quantity, product.Location, product.Status,
	)

	if err != nil {
		panic(err)
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return lastID, nil
}

func Detail(id int) (entities.Product, error) {
	row := config.DB.QueryRow(`
		SELECT id, name, sku, quantity, location, status FROM products WHERE id = ?`,
		id,
	)

	var product entities.Product
	if err := row.Scan(&product.Id, &product.Name, &product.Sku, &product.Quantity, &product.Location, &product.Status); err != nil {
		return entities.Product{}, err
	}

	return product, nil
}

func Update(id int, product entities.Product) (entities.Product, error) {
	_, err := config.DB.Exec(`
			UPDATE products
			SET name = ?, sku = ?, quantity = ?, location = ?, status = ?
			WHERE id = ?`,
		product.Name, product.Sku, product.Quantity, product.Location, product.Status, id,
	)

	if err != nil {
		return entities.Product{}, err
	}

	row := config.DB.QueryRow(`
		SELECT id, name, sku, quantity, location, status FROM products WHERE id = ?`,
		id,
	)

	var updatedProduct entities.Product
	if err := row.Scan(&updatedProduct.Id, &updatedProduct.Name, &updatedProduct.Sku, &updatedProduct.Quantity, &updatedProduct.Location, &updatedProduct.Status); err != nil {
		return entities.Product{}, err
	}
	return updatedProduct, nil
}

func Delete(id int) error {
	result, err := config.DB.Exec(`DELETE FROM products WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("product ID does not exist")
	}

	return nil
}
