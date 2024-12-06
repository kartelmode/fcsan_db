package models

import "time"

type User struct {
	Id        int       `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	BirthDate time.Time `db:"birth_date"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	BasketId  int       `db:"basket_id"`
}

type UserWallet struct {
	Id      int `db:"id"`
	Balance int `db:"balance"`
	UserId  int `db:"user_id"`
}

type Transaction struct {
	Id       int       `db:"id"`
	Amount   int       `db:"amount"`
	Time     time.Time `db:"time"`
	WalletId int       `db:"wallet_id"`
}

type Admin struct {
	Id     int `db:"id"`
	UserId int `db:"user_id"`
}

type Organization struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	PhoneNumber string `db:"phone_number"`
	UserId      int    `db:"user_id"`
}

type RequestStatus string

const (
	Waiting  RequestStatus = "Waiting"
	Rejected RequestStatus = "Rejected"
	Accepted RequestStatus = "Accepted"
)

type OrganizationRequest struct {
	Id             int           `db:"id"`
	Description    string        `db:"description"`
	Document       []byte        `db:"document"`
	Status         RequestStatus `db:"status"`
	OrganizationId int           `db:"organization_id"`
}

type OrganizationManager struct {
	Id             int `db:"id"`
	OrganizationId int `db:"organization_id"`
	UserId         int `db:"user_id"`
}

type Product struct {
	Id             int    `db:"id"`
	Name           string `db:"name"`
	Description    string `db:"description"`
	Cost           int    `db:"cost"`
	OrganizationId int    `db:"organization_id"`
}

type Basket struct {
	Id        int `db:"id"`
	TotalCost int `db:"total_cost"`
}

type BasketProduct struct {
	Id        int `db:"int"`
	ProductId int `db:"product_id"`
	BasketId  int `db:"basket_id"`
}

type UserOrder struct {
	Id        int           `db:"id"`
	Status    RequestStatus `db:"status"`
	BasketId  int           `db:"basket_id"`
	UserId    int           `db:"user_id"`
	TotalCost int           `db:"total_cost"`
}

type AvgTransaction struct {
	Transaction
	AvgTransaction int `db:"avg_transaction"`
}

type OrderStats struct {
	UserOrder
	Sum int `db:"sum"`
}
