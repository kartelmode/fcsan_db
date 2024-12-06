package console

import (
	"context"
	"fmt"
	"main/db"
	"main/infrastructure"
)

type Console struct {
	OrderRepo          *infrastructure.OrderRepository
	OrgRepo            *infrastructure.OrganizationRepository
	ProductRepo        *infrastructure.ProductRepository
	TransactionRepo    *infrastructure.TransactionRepository
	UserRepo           *infrastructure.UserRepository
	WalletRepo         *infrastructure.WalletRepository
	Readers            Readers
	TransactionManager db.TransactionManager
}

type IntReader interface {
	Read() (int, error)
	ReadRange(left, right int) (int, error)
}

type StringReader interface {
	Read() (string, error)
}

type PathReader interface {
	Read() (string, error)
}

type Readers interface {
	GetIntReader() IntReader
	GetStringReader() StringReader
	GetPathReader() PathReader
}

func (c *Console) readString(text *string) {
	for {
		value, err := c.Readers.GetStringReader().Read()
		if err != nil {
			continue
		}
		*text = value
		break
	}
}

func (c *Console) readInt(value *int) {
	for {
		val, err := c.Readers.GetIntReader().Read()
		if err != nil {
			continue
		}
		*value = val
		break
	}
}

func (c *Console) readIntRange(value *int, left, right int) {
	for {
		val, err := c.Readers.GetIntReader().ReadRange(left, right)
		if err != nil {
			continue
		}
		*value = val
		break
	}
}

func (c *Console) printCommands(ctx context.Context) {
	idCmd := -1
	if user == nil {
		idCmd = 0
	} else {
		if isAdmin, err := c.UserRepo.IsUserAdmin(ctx, user); isAdmin && err == nil {
			idCmd = 2
		} else {
			idCmd = 1
		}
	}
	for _, val := range Cmds[idCmd] {
		fmt.Println(val)
	}
}

func (c *Console) selectCommand(ctx context.Context) int {
	id := 0
	if user == nil {
		c.readIntRange(&id, 1, len(Cmds[0]))
	} else {
		if isAdmin, err := c.UserRepo.IsUserAdmin(ctx, user); isAdmin && err == nil {
			c.readIntRange(&id, 1, len(Cmds[2]))
		} else {
			c.readIntRange(&id, 1, len(Cmds[1]))
		}
	}
	return id
}

func (c *Console) executeCommand(ctx context.Context, cmd int) error {
	if user == nil {
		switch cmd {
		case 1:
			return c.Register(ctx)
		case 2:
			return c.Login(ctx)
		case 3:
			_, err := c.ShowProducts(ctx)
			return err
		}
	} else {
		if isAdmin, err := c.UserRepo.IsUserAdmin(ctx, user); isAdmin && err == nil {
			switch cmd {
			case 1:
				return c.Logout()
			case 2:
				_, err := c.ShowProducts(ctx)
				return err
			case 3:
				_, err := c.ShowOrganizationRequests(ctx)
				return err
			case 4:
				return c.ApproveOrganizationRequest(ctx)
			case 5:
				return c.RejectOrganizationRequest(ctx)
			}
		} else {
			switch cmd {
			case 1:
				return c.Logout()
			case 2:
				_, err := c.ShowProducts(ctx)
				return err
			case 3:
				return c.ShowOrdersStats(ctx)
			case 4:
				return c.TakeProductToBasket(ctx)
			case 5:
				return c.MakeAnOrder(ctx)
			case 6:
				return c.CreateOrganization(ctx)
			case 7:
				return c.CreateProduct(ctx)
			case 8:
				return c.ShowWallet(ctx)
			case 9:
				return c.TopBalance(ctx)
			case 10:
				return c.ShowTransactions(ctx)
			case 11:
				return c.ShowProductsInBasket(ctx)
			}
		}
	}
	return nil
}

func (c *Console) Run(ctx context.Context) {
	for {
		err := c.TransactionManager.Run(ctx, func(ctx context.Context) (rerr error) {
			c.printCommands(ctx)
			cmd := c.selectCommand(ctx)
			rerr = c.executeCommand(ctx, cmd)
			return rerr
		})
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
	}
}
