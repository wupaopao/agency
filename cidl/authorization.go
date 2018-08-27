package cidl

import (
	"encoding/json"

	"github.com/mz-eco/mz/log"
)

type RoleAuthorizationMap map[AuthorizationId]*Authorization

func NewRoleAuthorizationMap() *RoleAuthorizationMap {
	return &RoleAuthorizationMap{}
}

func (m *RoleAuthorizationMap) ToString() (strAuthorization string, err error) {
	bytesAuthorizations, err := json.Marshal(m)
	if err != nil {
		log.Warnf("marshal role authorization map failed.")
		return
	}

	strAuthorization = string(bytesAuthorizations)
	return
}

func (m *RoleAuthorizationMap) FromString(strAuthorization string) (err error) {
	err = json.Unmarshal([]byte(strAuthorization), m)
	if err != nil {
		log.Warnf("unmarshal str authorization failed")
		return
	}
	return
}

// 超级管理员
var (
	SuperAdministratorAuthorization = &Authorization{
		AuthorizationId: SuperAdministratorAuthorizationId,
		Title:           "超级管理员",
	}
)

// 成员管理
var (
	OrganizationManagementAuthorization = &Authorization{
		AuthorizationId: OrganizationManagementAuthorizationId,
		Title:           "成员管理",
	}

	ShowStaffAuthorization = &Authorization{
		AuthorizationId: ShowStaffAuthorizationId,
		Title:           "查看成员",
		ParentId:        OrganizationManagementAuthorization.AuthorizationId,
		ParentTitle:     OrganizationManagementAuthorization.Title,
	}

	EditStaffAuthorization = &Authorization{
		AuthorizationId: EditStaffAuthorizationId,
		Title:           "编辑成员",
		ParentId:        OrganizationManagementAuthorization.AuthorizationId,
		ParentTitle:     OrganizationManagementAuthorization.Title,
	}

	ShowStaffRoleAuthorization = &Authorization{
		AuthorizationId: ShowStaffRoleAuthorizationId,
		Title:           "查看角色",
		ParentId:        OrganizationManagementAuthorization.AuthorizationId,
		ParentTitle:     OrganizationManagementAuthorization.Title,
	}

	EditStaffRoleAuthorization = &Authorization{
		AuthorizationId: EditStaffRoleAuthorizationId,
		Title:           "编辑角色",
		ParentId:        OrganizationManagementAuthorization.AuthorizationId,
		ParentTitle:     OrganizationManagementAuthorization.Title,
	}
)

// 社群管理
var (
	CommunityManagementAuthorization = &Authorization{
		AuthorizationId: CommunityManagementAuthorizationId,
		Title:           "社群管理",
	}

	ShowGroupAuthorization = &Authorization{
		AuthorizationId: ShowGroupAuthorizationId,
		Title:           "查看社群",
		ParentId:        CommunityManagementAuthorization.AuthorizationId,
		ParentTitle:     CommunityManagementAuthorization.Title,
	}

	EditGroupAuthorization = &Authorization{
		AuthorizationId: EditGroupAuthorizationId,
		Title:           "编辑社群",
		ParentId:        CommunityManagementAuthorization.AuthorizationId,
		ParentTitle:     CommunityManagementAuthorization.Title,
	}
)

// 团购任务管理
var (
	GroupBuyingManagementAuthorization = &Authorization{
		AuthorizationId: GroupBuyingManagementAuthorizationId,
		Title:           "团购任务管理",
	}

	ShowTaskAuthorization = &Authorization{
		AuthorizationId: ShowTaskAuthorizationId,
		Title:           "查看团购任务",
		ParentId:        GroupBuyingManagementAuthorization.AuthorizationId,
		ParentTitle:     GroupBuyingManagementAuthorization.Title,
	}

	EditTaskAuthorization = &Authorization{
		AuthorizationId: EditTaskAuthorizationId,
		Title:           "编辑团购任务",
		ParentId:        GroupBuyingManagementAuthorization.AuthorizationId,
		ParentTitle:     GroupBuyingManagementAuthorization.Title,
	}
)

// 权限ID映射，不包括超级管理员
var RoleAuthorizationMapIds = RoleAuthorizationMap{
	OrganizationManagementAuthorizationId: OrganizationManagementAuthorization,
	ShowStaffAuthorizationId:              ShowStaffAuthorization,
	EditStaffAuthorizationId:              EditStaffAuthorization,
	ShowStaffRoleAuthorizationId:          ShowStaffRoleAuthorization,
	EditStaffRoleAuthorizationId:          EditStaffRoleAuthorization,

	CommunityManagementAuthorizationId: CommunityManagementAuthorization,
	ShowGroupAuthorizationId:           ShowGroupAuthorization,
	EditGroupAuthorizationId:           EditGroupAuthorization,

	GroupBuyingManagementAuthorizationId: GroupBuyingManagementAuthorization,
	ShowTaskAuthorizationId:              ShowTaskAuthorization,
	EditTaskAuthorizationId:              EditTaskAuthorization,
}

// 成员角色权限
func NewStaffRoleAuthorizationByRole(role *StaffRole, authorization *Authorization) *StaffRoleAuthorization {
	staffAuth := &StaffRoleAuthorization{
		Authorization: *authorization,
	}
	_, ok := (*role.RoleAuthorization)[staffAuth.AuthorizationId]
	if ok {
		staffAuth.IsOwn = true
	}

	return staffAuth
}

// 创建组织成员的权限列表，不包括超级管理员
func NewStaffRoleAuthorizationGroupsByRole(role *StaffRole) (groups []*StaffRoleAuthorizationGroup) {
	organizationManagement := NewStaffRoleAuthorizationByRole(role, OrganizationManagementAuthorization)
	showStaff := NewStaffRoleAuthorizationByRole(role, ShowStaffAuthorization)
	editStaff := NewStaffRoleAuthorizationByRole(role, EditStaffAuthorization)
	showRole := NewStaffRoleAuthorizationByRole(role, ShowStaffRoleAuthorization)
	editRole := NewStaffRoleAuthorizationByRole(role, EditStaffRoleAuthorization)

	// 团购组织管理
	groups = append(groups, &StaffRoleAuthorizationGroup{
		ModuleAuthorization: organizationManagement,
		SubAuthorizations: map[AuthorizationId]*StaffRoleAuthorization{
			showStaff.AuthorizationId: showStaff,
			editStaff.AuthorizationId: editStaff,
			showRole.AuthorizationId:  showRole,
			editRole.AuthorizationId:  editRole,
		},
	})

	// 社群管理
	communityManagement := NewStaffRoleAuthorizationByRole(role, CommunityManagementAuthorization)
	showGroup := NewStaffRoleAuthorizationByRole(role, ShowGroupAuthorization)
	editGroup := NewStaffRoleAuthorizationByRole(role, EditGroupAuthorization)
	groups = append(groups, &StaffRoleAuthorizationGroup{
		ModuleAuthorization: communityManagement,
		SubAuthorizations: map[AuthorizationId]*StaffRoleAuthorization{
			showGroup.AuthorizationId: showGroup,
			editGroup.AuthorizationId: editGroup,
		},
	})

	// 团购任务管理
	groupBuying := NewStaffRoleAuthorizationByRole(role, GroupBuyingManagementAuthorization)
	showTask := NewStaffRoleAuthorizationByRole(role, ShowTaskAuthorization)
	editTask := NewStaffRoleAuthorizationByRole(role, EditTaskAuthorization)
	groups = append(groups, &StaffRoleAuthorizationGroup{
		ModuleAuthorization: groupBuying,
		SubAuthorizations: map[AuthorizationId]*StaffRoleAuthorization{
			showTask.AuthorizationId: showTask,
			editTask.AuthorizationId: editTask,
		},
	})

	return
}
