// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id          uint        `json:"id"          description:"User ID"`
	Passport    string      `json:"passport"    description:"User Passport"`
	Password    string      `json:"password"    description:"User Password"`
	Nickname    string      `json:"nickname"    description:"User Nickname"`
	CreateAt    *gtime.Time `json:"createAt"    description:"Created Time"`
	UpdateAt    *gtime.Time `json:"updateAt"    description:"Updated Time"`
	Phonenumber string      `json:"phonenumber" description:""`
	IsAdmin     int         `json:"isAdmin"     description:""`
	Rank        string      `json:"rank"        description:""`
	TimeStamp 	int64
}
