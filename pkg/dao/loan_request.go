package dao

import (
	"context"

	"bitbucket.com/finease/backend/pkg/models"
)

type LoanRequest interface {
	Create(ctx context.Context, loanRequest *models.LoanRequest) (*models.LoanRequest, error)
	Update(ctx context.Context, id string, patch *models.LoanRequest) (*models.LoanRequest, error)
	Delete(ctx context.Context, id string) error
	FindByUserId(ctx context.Context, userUuid string) ([]*models.LoanRequest, error)
	FindByProposalId(ctx context.Context, proposalUuid string) ([]*models.LoanRequest, error)
	FindById(ctx context.Context, id string) (*models.LoanRequest, error)
	FindAll(ctx context.Context) ([]*models.LoanRequest, error)
}
