// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/gogf/gf-demo-user/v2/internal/dao/internal"
)

// internalVerificationcodeDao is internal type for wrapping internal DAO implements.
type internalVerificationcodeDao = *internal.VerificationcodeDao

// verificationcodeDao is the data access object for table verificationcode.
// You can define custom methods on it to extend its functionality as you wish.
type verificationcodeDao struct {
	internalVerificationcodeDao
}

var (
	// Verificationcode is globally public accessible object for table verificationcode operations.
	Verificationcode = verificationcodeDao{
		internal.NewVerificationcodeDao(),
	}
)

// Fill with you ideas below.
