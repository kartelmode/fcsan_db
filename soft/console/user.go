package console

import (
	"context"
	"fmt"
	"main/models"
	"time"
)

func (c *Console) Register(ctx context.Context) error {
	user = &models.User{}
	fmt.Println("Type your first name:")
	c.readString(&user.FirstName)
	fmt.Println("Type your last name:")
	c.readString(&user.LastName)
	for {
		birthDate := ""
		fmt.Println("Type your birthday in format yyyy-mm-dd:")
		c.readString(&birthDate)
		var err error
		user.BirthDate, err = time.Parse("2006-01-02", birthDate)
		if err == nil {
			break
		}
		fmt.Println("Type a valid date: ", err.Error())
	}
	fmt.Println("Type your password:")
	c.readString(&user.Password)
	fmt.Println("Type your email:")
	c.readString(&user.Email)
	err := c.UserRepo.RegisterUser(ctx, user)
	if err != nil {
		user = nil
		fmt.Println("Error occured:", err.Error())
		return err
	}
	c.UserRepo.GetUserByEmailAndPassword(ctx, user)
	return nil
}

func (c *Console) Login(ctx context.Context) error {
	user = &models.User{}
	fmt.Println("Type your email:")
	c.readString(&user.Email)
	fmt.Println("Type your password:")
	c.readString(&user.Password)
	_, err := c.UserRepo.GetUserByEmailAndPassword(ctx, user)
	if err != nil {
		fmt.Println("Error occured:", err.Error())
		user = nil
	}
	return err
}

func (c *Console) Logout() error {
	user = nil
	return nil
}

func (c *Console) ShowWallet(ctx context.Context) error {
	wallet, err := c.WalletRepo.GetWalletByUserId(ctx, user)
	if err != nil {
		return err
	}
	fmt.Println("Balance: ", wallet.Balance)
	return nil
}

func (c *Console) TopBalance(ctx context.Context) error {
	fmt.Println("Type amount of transaction:")
	transaction := &models.Transaction{}
	c.readInt(&transaction.Amount)

	wallet, err := c.WalletRepo.GetWalletByUserId(ctx, user)
	if err != nil {
		fmt.Println("Error occured:", err.Error())
		return err
	}
	err = c.TransactionRepo.MakeTransaction(ctx, wallet, transaction)
	return err
}

func (c *Console) ShowTransactions(ctx context.Context) error {
	transactions, err := c.TransactionRepo.GetAllTransactionsForUser(ctx, user)
	if err != nil {
		fmt.Println("Error occured:", err.Error())
		return err
	}
	if len(transactions) == 0 {
		fmt.Println("You have made 0 transactions")
		return nil
	}
	for i := 0; i < len(transactions); i++ {
		fmt.Println("Amount of transaction:", transactions[i].Amount, "created at", transactions[i].Time)
	}
	return nil
}
