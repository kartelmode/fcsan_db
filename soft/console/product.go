package console

import (
	"context"
	"fmt"
	"main/models"
)

func (c *Console) ShowProducts(ctx context.Context) ([]*models.Product, error) {
	products, err := c.ProductRepo.GetAllProducts(ctx)
	if err != nil {
		fmt.Println("Error occured: ", err.Error())
		return nil, err
	}
	for i := 0; i < len(products); i++ {
		fmt.Println(i+1, ". Product: ", products[i].Name, products[i].Description, products[i].Cost)
	}
	return products, nil
}

func (c *Console) ShowProductsInBasket(ctx context.Context) error {
	products, err := c.ProductRepo.GetAllProductsInUserBasket(ctx, user)
	if err != nil {
		fmt.Println("Error occured: ", err.Error())
		return err
	}
	for i := 0; i < len(products); i++ {
		fmt.Println(i+1, ". Product: ", products[i].Name, "with description:\n", products[i].Description, products[i].Cost)
	}
	return nil
}

func (c *Console) TakeProductToBasket(ctx context.Context) error {
	fmt.Println("Select a product:")
	products, err := c.ShowProducts(ctx)
	if err != nil {
		return err
	}
	if len(products) == 0 {
		fmt.Println("There are no available products to take")
		return nil
	}
	for {
		value, err := c.Readers.GetIntReader().ReadRange(1, len(products))
		if err != nil {
			continue
		}
		c.UserRepo.TakeProductToBasket(ctx, user, products[value-1])
		break
	}
	return nil
}

func (c *Console) ShowOrdersStats(ctx context.Context) error {
	stats, err := c.OrderRepo.GetAllOrdersForUser(ctx, user)
	if err != nil {
		fmt.Println("An error occured: ", err.Error())
		return err
	}
	for i := 0; i < len(stats); i++ {
		fmt.Println("Order with id ", stats[i].Id, " costs ", stats[i].TotalCost, " in status ", stats[i].Status)
	}
	fmt.Println()
	return nil
}

func (c *Console) MakeAnOrder(ctx context.Context) error {
	err := c.UserRepo.MakeAnOrder(ctx, user)
	if err != nil {
		fmt.Println("An error occured: ", err.Error())
	}
	fmt.Println()
	return err
}

func (c *Console) CreateProduct(ctx context.Context) error {
	product := &models.Product{}

	fmt.Println("Select an organization for a product:")
	orgs, err := c.ShowAdminsOrganizations(ctx)
	if err != nil {
		return err
	}
	var id int = 0
	c.readIntRange(&id, 1, len(orgs))
	product.OrganizationId = orgs[id-1].Id

	fmt.Println("Type name of a product:")
	c.readString(&product.Name)
	fmt.Println("Type description of a product:")
	c.readString(&product.Description)
	fmt.Println("Type cost of a product:")
	c.readInt(&product.Cost)

	err = c.ProductRepo.CreateProduct(ctx, product)
	if err != nil {
		fmt.Println("Error occured: ", err.Error())
	}
	return nil
}
