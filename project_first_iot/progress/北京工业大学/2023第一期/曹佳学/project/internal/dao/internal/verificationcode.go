// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VerificationcodeDao is the data access object for table verificationcode.
type VerificationcodeDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns VerificationcodeColumns // columns contains all the column names of Table for convenient usage.
}

// VerificationcodeColumns defines and stores column names for table verificationcode.
type VerificationcodeColumns struct {
	Phonenumber      string //
	Verificationcode string //
	CreateAt         string // Created Time
	Id               string //
}

// verificationcodeColumns holds the columns for table verificationcode.
var verificationcodeColumns = VerificationcodeColumns{
	Phonenumber:      "phonenumber",
	Verificationcode: "verificationcode",
	CreateAt:         "create_at",
	Id:               "id",
}

// NewVerificationcodeDao creates and returns a new DAO object for table data access.
func NewVerificationcodeDao() *VerificationcodeDao {
	return &VerificationcodeDao{
		group:   "default",
		table:   "verificationcode",
		columns: verificationcodeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VerificationcodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *VerificationcodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *VerificationcodeDao) Columns() VerificationcodeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *VerificationcodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VerificationcodeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VerificationcodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
