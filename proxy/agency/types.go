package agency

import "business/agency/cidl"

type RoleAuthorizationMap cidl.RoleAuthorizationMap

func NewRoleAuthorizationMap() *RoleAuthorizationMap {
	return &RoleAuthorizationMap{}
}
