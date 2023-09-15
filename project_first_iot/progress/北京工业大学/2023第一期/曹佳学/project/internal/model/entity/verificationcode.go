// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Verificationcode is the golang structure for table verificationcode.
type Verificationcode struct {
	Phonenumber      string      `json:"phonenumber"      description:""`
	Verificationcode string      `json:"verificationcode" description:""`
	CreateAt         *gtime.Time `json:"createAt"         description:"Created Time"`
	Id               int         `json:"id"               description:""`
}
