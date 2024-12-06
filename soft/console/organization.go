package console

import (
	"context"
	"fmt"
	"main/models"
)

func (c *Console) CreateOrganization(ctx context.Context) error {
	organization := &models.Organization{}
	fmt.Println("Type name of organization:")
	c.readString(&organization.Name)
	fmt.Println("Type description:")
	c.readString(&organization.Description)
	fmt.Println("Type phone number:")
	c.readString(&organization.PhoneNumber)
	organization.UserId = user.Id

	err := c.OrgRepo.CreateOrganization(ctx, organization)
	if err != nil {
		fmt.Println("Error occured: ", err.Error())
		return err
	}

	organizationReq := &models.OrganizationRequest{
		Description:    organization.Description,
		Document:       []byte{}, // should read path for a file and read a content from it
		Status:         models.Waiting,
		OrganizationId: organization.Id,
	}
	err = c.OrgRepo.CreateOrganizationRequest(ctx, organizationReq)
	if err != nil {
		fmt.Println("Error occured: ", err.Error())
	}
	return err
}

func (c *Console) ShowAdminsOrganizations(ctx context.Context) ([]*models.Organization, error) {
	orgs, err := c.OrgRepo.GetAllOrganizationsForOrgAdmin(ctx, user)
	if err != nil {
		fmt.Println("Error occured: ", err.Error())
		return nil, err
	}
	for i := 0; i < len(orgs); i++ {
		fmt.Println(i+1, ". Organization: ", orgs[i].Name, " with a phone number ", orgs[i].PhoneNumber)
	}
	return orgs, nil
}

func (c *Console) ShowOrganizationRequests(ctx context.Context) ([]*models.OrganizationRequest, error) {
	orgReqs, err := c.OrgRepo.GetWaitingOrganizationRequests(ctx, user)
	if err != nil {
		fmt.Println("Error occured:", err.Error())
		return nil, err
	}
	for i := 0; i < len(orgReqs); i++ {
		fmt.Println(i+1, ". Waining organization request with description:", orgReqs[i].Description)
	}
	return orgReqs, nil
}

func (c *Console) ApproveOrganizationRequest(ctx context.Context) error {
	orgReqs, err := c.ShowOrganizationRequests(ctx)
	if err != nil {
		return err
	}
	if len(orgReqs) == 0 {
		fmt.Println("There are no waiting requests for creating organizations")
	}
	id := 0
	c.readIntRange(&id, 1, len(orgReqs))
	id--
	orgReqs[id].Status = models.Accepted
	err = c.OrgRepo.UpdateOrganizationRequest(ctx, orgReqs[id])
	if err != nil {
		fmt.Println("Error occured:", err.Error())
	}
	return err
}

func (c *Console) RejectOrganizationRequest(ctx context.Context) error {
	orgReqs, err := c.ShowOrganizationRequests(ctx)
	if err != nil {
		return err
	}
	if len(orgReqs) == 0 {
		fmt.Println("There are no waiting requests for creating organizations")
	}
	id := 0
	c.readIntRange(&id, 1, len(orgReqs))
	id--
	orgReqs[id].Status = models.Rejected
	err = c.OrgRepo.UpdateOrganizationRequest(ctx, orgReqs[id])
	if err != nil {
		fmt.Println("Error occured:", err.Error())
	}
	return err
}
