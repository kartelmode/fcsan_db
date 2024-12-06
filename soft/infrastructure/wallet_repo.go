package infrastructure

import (
	"context"
	"main/apperrors"
	"main/db"
	"main/models"
)

type WalletRepository struct {
	Db db.SqlDatabaseProvider
}

func (repo *WalletRepository) GetWalletByUserId(ctx context.Context, user *models.User) (*models.UserWallet, error) {
	query := `SELECT * FROM user_wallet WHERE user_id = ?`

	wallet := &models.UserWallet{}
	row := repo.Db.QueryRowContext(ctx, query, user.Id)
	err := row.Scan(&wallet.Id, &wallet.Balance, &wallet.UserId)
	if err != nil {
		return nil, &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	return wallet, nil
}
