package infrastructure

import (
	"context"
	"main/apperrors"
	"main/db"
	"main/models"
)

type OrderRepository struct {
	Db db.SqlDatabaseProvider
}

func (repo *OrderRepository) GetAllOrdersForUser(
	ctx context.Context, user *models.User) ([]*models.UserOrder, error) {
	query := `SELECT id, status, total_cost
				FROM user_order as o
				WHERE user_id = ?`

	rows, err := repo.Db.QueryContext(ctx, query, user.Id)
	if err != nil {
		return nil, &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}

	orders := make([]*models.UserOrder, 0)
	for rows.Next() {
		var order models.UserOrder
		err := rows.Scan(
			&order.Id,
			&order.Status, &order.TotalCost,
		)
		if err != nil {
			return nil, &apperrors.ErrInternal{
				Message: err.Error(),
			}
		}
		orders = append(orders, &order)
	}
	return orders, nil
}
