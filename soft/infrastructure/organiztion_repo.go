package infrastructure

import (
	"context"
	"main/apperrors"
	"main/db"
	"main/models"
)

type OrganizationRepository struct {
	Db db.SqlDatabaseProvider
}

func (repo *OrganizationRepository) CreateOrganization(ctx context.Context, organization *models.Organization) error {
	query := `INSERT INTO organization 
				(name, description, phone_number, user_id) 
				VALUES (?, ?, ?, ?)`

	_, err := repo.Db.ExecContext(ctx, query,
		organization.Name, organization.Description,
		organization.PhoneNumber, organization.UserId)
	if err != nil {
		return &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	err = repo.Db.QueryRowContext(ctx, `SELECT LAST_INSERT_ID()`).Scan(&organization.Id)
	if err != nil {
		return &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	return nil
}

func (repo *OrganizationRepository) CreateOrganizationRequest(ctx context.Context, organizationReq *models.OrganizationRequest) error {
	query := `INSERT INTO organization_request 
				(description, document, status, organization_id) 
				VALUES (?, ?, ?, ?)`

	_, err := repo.Db.ExecContext(ctx, query,
		organizationReq.Description, organizationReq.Document,
		organizationReq.Status, organizationReq.OrganizationId)
	if err != nil {
		return &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	return nil
}

func (repo *OrganizationRepository) GetAllOrganizationsForOrgAdmin(
	ctx context.Context, user *models.User) ([]*models.Organization, error) {
	query := `SELECT o.* FROM organization AS o
				INNER JOIN organization_manager AS om ON om.organization_id = o.id
				INNER JOIN organization_request AS r ON r.organization_id = o.id
				WHERE om.user_id = ? AND r.status = 'Accepted'`

	rows, err := repo.Db.QueryContext(ctx, query, user.Id)
	if err != nil {
		return nil, &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}

	organizations := make([]*models.Organization, 0)
	for rows.Next() {
		var organization models.Organization
		err := rows.Scan(
			&organization.Id,
			&organization.Name, &organization.Description,
			&organization.PhoneNumber,
			&organization.UserId)
		if err != nil {
			return nil, &apperrors.ErrInternal{
				Message: err.Error(),
			}
		}
		organizations = append(organizations, &organization)
	}
	return organizations, nil
}

func (repo *OrganizationRepository) GetAllOrganizations(
	ctx context.Context, user *models.User) ([]*models.Organization, error) {
	query := `SELECT o.* FROM organization AS o
				INNER JOIN organization_request as r
				on r.status = 'Accepted' AND o.id = r.organization_id`

	rows, err := repo.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}

	organizations := make([]*models.Organization, 0)
	for rows.Next() {
		var organization models.Organization
		err := rows.Scan(
			&organization.Id,
			&organization.Name, &organization.Description,
			&organization.PhoneNumber,
			&organization.UserId)
		if err != nil {
			return nil, &apperrors.ErrInternal{
				Message: err.Error(),
			}
		}
		organizations = append(organizations, &organization)
	}
	return organizations, nil
}

func (repo *OrganizationRepository) GetWaitingOrganizationRequests(
	ctx context.Context, user *models.User) ([]*models.OrganizationRequest, error) {
	query := `SELECT * FROM organization_request WHERE status = 'Waiting'`

	rows, err := repo.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}

	organizationReqs := make([]*models.OrganizationRequest, 0)
	for rows.Next() {
		var organizationReq models.OrganizationRequest
		err := rows.Scan(
			&organizationReq.Id,
			&organizationReq.Description,
			&organizationReq.Document, &organizationReq.Status,
			&organizationReq.OrganizationId)
		if err != nil {
			return nil, &apperrors.ErrInternal{
				Message: err.Error(),
			}
		}
		organizationReqs = append(organizationReqs, &organizationReq)
	}
	return organizationReqs, nil
}

func (repo *OrganizationRepository) UpdateOrganizationRequest(
	ctx context.Context, organizationReq *models.OrganizationRequest) error {
	query := `UPDATE organization_request
				SET status = ?`
	_, err := repo.Db.ExecContext(ctx, query, organizationReq.Status)
	if err != nil {
		return &apperrors.ErrInternal{
			Message: err.Error(),
		}
	}
	return nil
}
