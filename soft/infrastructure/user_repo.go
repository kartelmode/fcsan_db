package infrastructure

import (
	"context"
	"main/apperrors"
	"main/db"
	"main/models"
	"time"
)

type UserRepository struct {
	Db db.SqlDatabaseProvider
}

func (repo *UserRepository) RegisterUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO user 
				(first_name, last_name, birth_date, password, email) 
          		VALUES (?, ?, ?, ?, ?)`

	_, err := repo.Db.ExecContext(ctx, query,
		user.FirstName, user.LastName,
		user.BirthDate.Format("2006-01-02"), user.Password,
		user.Email)
	if err != nil {
		return &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}

	err = repo.Db.QueryRowContext(ctx, `SELECT LAST_INSERT_ID()`).Scan(&user.Id)
	if err != nil {
		return &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	return nil
}

func (repo *UserRepository) GetUserByEmailAndPassword(ctx context.Context, user *models.User) (bool, error) {
	query := `SELECT * FROM user WHERE email = ? AND password = ?`

	var birthDate []uint8
	row := repo.Db.QueryRowContext(ctx, query, user.Email, user.Password)
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &birthDate, &user.Password, &user.Email, &user.BasketId)
	if err != nil {
		return false, &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	if len(birthDate) > 0 {
		t, err := time.Parse("2006-01-02", string(birthDate))
		if err != nil {
			return false, &apperrors.ErrInternal{
				Message: err.Error(),
			}
		}
		user.BirthDate = t
	}

	return true, nil
}

func (repo *UserRepository) IsUserAdmin(ctx context.Context, user *models.User) (bool, error) {
	query := `SELECT count(*) FROM admin WHERE user_id = ?`

	cnt := 0
	row := repo.Db.QueryRowContext(ctx, query, user.Id)
	err := row.Scan(&cnt)
	if err != nil {
		return false, &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	return cnt > 0, nil
}

func (repo *UserRepository) IsOrganizationManager(ctx context.Context, user *models.User, organizationId int) (bool, error) {
	query := `SELECT count(*) FROM organization_manager 
				WHERE user_id = ? AND organization_id = ?`

	cnt := 0
	row := repo.Db.QueryRowContext(ctx, query, user.Id, organizationId)
	err := row.Scan(&cnt)
	if err != nil {
		return false, &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	return cnt > 0, nil
}

func (repo *UserRepository) TakeProductToBasket(ctx context.Context, user *models.User, product *models.Product) error {
	query := `INSERT INTO basket_product 
				(basket_id, product_id) 
				VALUES (?, ?)`

	_, err := repo.Db.ExecContext(ctx, query, user.BasketId, product.Id)
	if err != nil {
		return &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	return nil
}

func (repo *UserRepository) MakeAnOrder(ctx context.Context, user *models.User) error {
	query := `INSERT INTO user_order 
				(status, basket_id, user_id) 
				VALUES (?, ?, ?)`

	_, err := repo.Db.ExecContext(ctx, query,
		models.Accepted, user.BasketId, user.Id)
	if err != nil {
		return &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	return nil
}
