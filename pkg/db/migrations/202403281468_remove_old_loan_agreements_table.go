package migrations

// Migrations should NEVER use types from other packages. Types can change
// and then migrations run on a _new_ database will fail or behave unexpectedly.
// Instead of importing types, always re-create the type in the migration, as
// is done here, even though the same type is defined in pkg/api

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func removeOldLoanAgreementsTable() *gormigrate.Migration {
	type LoanAgreement struct{}
	return &gormigrate.Migration{
		ID: "202403281468",
		Migrate: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&LoanAgreement{})
		},
	}
}
