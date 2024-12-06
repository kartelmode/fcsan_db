package main

import (
	"context"
	"fmt"
	"main/config"
	"main/console"
	"main/db"
	"main/infrastructure"
	"main/readers"
)

func main() {
	databasecfg := config.Database{}
	err := config.ReadConfig("../config/db.json", &databasecfg)
	if err != nil {
		fmt.Println("unable to read config")
		return
	}
	database, err := db.NewPostgresSqlDatabaseProvider(&databasecfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	transactionManager := db.NewTransactionManager(database.GetDb())

	console := &console.Console{
		OrderRepo: &infrastructure.OrderRepository{
			Db: database,
		},
		OrgRepo: &infrastructure.OrganizationRepository{
			Db: database,
		},
		ProductRepo: &infrastructure.ProductRepository{
			Db: database,
		},
		TransactionRepo: &infrastructure.TransactionRepository{
			Db: database,
		},
		UserRepo: &infrastructure.UserRepository{
			Db: database,
		},
		WalletRepo: &infrastructure.WalletRepository{
			Db: database,
		},
		Readers:            readers.NewReaders(),
		TransactionManager: transactionManager,
	}
	fmt.Println("Database successfully configured")
	console.Run(context.Background())
}
