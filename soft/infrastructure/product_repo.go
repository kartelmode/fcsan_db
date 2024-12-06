package infrastructure

import (
	"context"
	"main/apperrors"
	"main/db"
	"main/models"
)

type ProductRepository struct {
	Db db.SqlDatabaseProvider
}

func (repo *ProductRepository) CreateProduct(ctx context.Context, product *models.Product) error {
	query := `INSERT INTO product 
				(name, description, cost, organization_id) 
				VALUES (?, ?, ?, ?)`

	_, err := repo.Db.ExecContext(ctx, query,
		product.Name, product.Description,
		product.Cost, product.OrganizationId)
	if err != nil {
		return &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	return nil
}

func (repo *ProductRepository) GetAllProducts(ctx context.Context) ([]*models.Product, error) {
	query := `SELECT p.*
				FROM product as p
				WHERE (SELECT count(1) 
						FROM basket_product as bp
						INNER JOIN user_order as o
						ON o.basket_id = bp.basket_id AND bp.product_id = p.id AND status = 'Accepted') = 0`
	rows, err := repo.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}

	products := make([]*models.Product, 0)
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.Id,
			&product.Name, &product.Description, &product.Cost,
			&product.OrganizationId)
		if err != nil {
			return nil, &apperrors.ErrInternal{
				Message: err.Error(),
			}
		}
		products = append(products, &product)
	}
	return products, nil
}

func (repo *ProductRepository) GetAllProductsInUserBasket(
	ctx context.Context, user *models.User) ([]*models.Product, error) {
	query := `SELECT p.*
				FROM product as p
				INNER JOIN basket_product as bp
				ON bp.product_id = p.id
				AND bp.basket_id = (SELECT u.basket_id FROM user as u WHERE u.id = ?)`
	rows, err := repo.Db.QueryContext(ctx, query, user.Id)
	if err != nil {
		return nil, &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}

	products := make([]*models.Product, 0)
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.Id,
			&product.Name, &product.Description, &product.Cost,
			&product.OrganizationId)
		if err != nil {
			return nil, &apperrors.ErrInternal{
				Message: err.Error(),
			}
		}
		products = append(products, &product)
	}
	return products, nil
}
