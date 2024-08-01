package repository

import (
	"database/sql"
	"go-docker-api/internal/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connnection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connnection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)

	if err != nil {
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		if errRows := rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		); errRows != nil {
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int

	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		" VALUES ($1, $2) RETURN id")

	if err != nil {
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")

	if err != nil {
		return nil, err
	}

	var product model.Product

	if err = query.QueryRow(id_product).Scan(&product.ID, &product.Name, &product.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &product, nil
}
