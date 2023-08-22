// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Verificationcode is the golang structure of table verificationcode for DAO operations like Where/Data.
type Verificationcode struct {
	g.Meta           `orm:"table:verificationcode, do:true"`
	Phonenumber      interface{} //
	Verificationcode interface{} //
	CreateAt         *gtime.Time // Created Time
	Id               interface{} //
}
