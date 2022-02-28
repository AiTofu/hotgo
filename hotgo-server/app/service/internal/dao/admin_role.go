// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/bufanyun/hotgo/app/service/internal/dao/internal"
)

// adminRoleDao is the data access object for table hg_admin_role.
// You can define custom methods on it to extend its functionality as you wish.
type adminRoleDao struct {
	*internal.AdminRoleDao
}

var (
	// AdminRole is globally public accessible object for table hg_admin_role operations.
	AdminRole = adminRoleDao{
		internal.NewAdminRoleDao(),
	}
)

// Fill with you ideas below.