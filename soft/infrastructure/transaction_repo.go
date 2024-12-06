package infrastructure

import (
	"context"
	"main/apperrors"
	"main/db"
	"main/models"
	"time"
)

type TransactionRepository struct {
	Db db.SqlDatabaseProvider
}

func (repo *TransactionRepository) MakeTransaction(ctx context.Context,
	wallet *models.UserWallet, transaction *models.Transaction) error {
	query := `INSERT INTO transaction (amount, time, wallet_id) 
				VALUES (?, ?, ?)`

	_, err := repo.Db.ExecContext(ctx, query,
		transaction.Amount, time.Now().Format("2006-01-02 15:04:05"), wallet.Id)
	if err != nil {
		return &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	return nil
}

func (repo *TransactionRepository) GetAllTransactionsForUser(ctx context.Context,
	user *models.User) ([]*models.Transaction, error) {
	query := `SELECT * FROM transaction as t
		WHERE t.wallet_id = (SELECT w.id FROM user_wallet as w WHERE w.user_id = ?)`

	rows, err := repo.Db.QueryContext(
		ctx, query,
		user.Id)
	if err != nil {
		return nil, &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	transactions := make([]*models.Transaction, 0)
	for rows.Next() {
		var transaction models.Transaction
		var tm []uint8
		err := rows.Scan(
			&transaction.Id,
			&transaction.Amount,
			&tm, &transaction.WalletId)
		if err != nil {
			return nil, &apperrors.ErrInternal{
				Message: err.Error(),
			}
		}
		if len(tm) > 0 {
			t, err := time.Parse("2006-01-02 15:04:05", string(tm))
			if err != nil {
				return nil, &apperrors.ErrInternal{
					Message: err.Error(),
				}
			}
			transaction.Time = t
		}
		transactions = append(transactions, &transaction)
	}
	return transactions, nil
}

func (repo *TransactionRepository) GetAvgTransactions(ctx context.Context) ([]*models.AvgTransaction, error) {
	query := `SELECT t.*, avg(t.amount) over(partition by wallet_id) as avg_transaction
					FROM transaction as t
					WHERE t.amount > 0`
	rows, err := repo.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}

	transactions := make([]*models.AvgTransaction, 0)
	for rows.Next() {
		var transaction models.AvgTransaction
		err := rows.Scan(
			&transaction.Id,
			&transaction.Amount,
			&transaction.Time, &transaction.WalletId,
			&transaction.AvgTransaction)
		if err != nil {
			return nil, &apperrors.ErrInternal{
				Message: err.Error(),
			}
		}
		transactions = append(transactions, &transaction)
	}
	return transactions, nil
}
