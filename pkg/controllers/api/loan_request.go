package api

import (
	"bitbucket.com/finease/backend/pkg/models"
	"bitbucket.com/finease/backend/pkg/utils"
)

type LoanRequest struct {
	Uuid          string  `json:"uuid,omitempty"`
	UserUUID      string  `json:"user_uuid,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
	MinInterest   float64 `json:"min_interest,omitempty"`
	MaxInterest   float64 `json:"max_interest,omitempty"`
	DurationToPay int64   `json:"duration,omitempty"`
	Status        *string `json:"status,omitempty"`
	Description   string  `json:"description,omitempty"`
	ProposalUuid  *string `json:"proposal_uuid"`
}

func MapLoanRequestModelToApi(loanRequest *models.LoanRequest) *LoanRequest {
	return &LoanRequest{
		Uuid:          loanRequest.Uuid,
		UserUUID:      loanRequest.UserUUID,
		Amount:        loanRequest.Amount,
		MinInterest:   loanRequest.MinInterest,
		MaxInterest:   loanRequest.MaxInterest,
		DurationToPay: loanRequest.DurationToPay,
		Status:        utils.PtrCopy(loanRequest.Status),
		Description:   loanRequest.Description,
		ProposalUuid:  utils.PtrCopy(loanRequest.ProposalUuid),
	}
}

func MapLoanRequestApiToModel(loanRequest *LoanRequest) *models.LoanRequest {
	return &models.LoanRequest{
		UserUUID:      loanRequest.UserUUID,
		Amount:        loanRequest.Amount,
		MinInterest:   loanRequest.MinInterest,
		MaxInterest:   loanRequest.MaxInterest,
		DurationToPay: loanRequest.DurationToPay,
		Status:        utils.PtrCopy(loanRequest.Status),
		Description:   loanRequest.Description,
		ProposalUuid:  utils.PtrCopy(loanRequest.ProposalUuid),
	}
}
