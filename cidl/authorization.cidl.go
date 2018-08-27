package cidl

// 权限
type Authorization struct {
	AuthorizationId AuthorizationId `db:"AuthorizationId"`
	Title           string          `db:"Title"`
	ParentId        AuthorizationId `db:"ParentId"`
	ParentTitle     string          `db:"ParentTitle"`
}

func NewAuthorization() *Authorization {
	return &Authorization{}
}

// 成员管理板块权限ID
type AuthorizationId int

const (
	// 超级管理员
	SuperAdministratorAuthorizationId AuthorizationId = 1
	// （组）团购组织管理（成员管理）
	OrganizationManagementAuthorizationId AuthorizationId = 1000
	// 查看成员
	ShowStaffAuthorizationId AuthorizationId = 1001
	// 编辑成员
	EditStaffAuthorizationId AuthorizationId = 1002
	// 查看角色
	ShowStaffRoleAuthorizationId AuthorizationId = 1003
	// 编辑角色
	EditStaffRoleAuthorizationId AuthorizationId = 1004
	// （组）社群管理
	CommunityManagementAuthorizationId AuthorizationId = 2000
	// 查看社群
	ShowGroupAuthorizationId AuthorizationId = 2001
	// 编辑社群
	EditGroupAuthorizationId AuthorizationId = 2002
	// （组）团购任务管理
	GroupBuyingManagementAuthorizationId AuthorizationId = 3000
	// 查看团购任务
	ShowTaskAuthorizationId AuthorizationId = 3001
	// 编辑团购任务
	EditTaskAuthorizationId AuthorizationId = 3002
)

func (m AuthorizationId) String() string {
	switch m {

	case SuperAdministratorAuthorizationId:
		return "SuperAdministratorAuthorizationId<enum AuthorizationId>"
	case OrganizationManagementAuthorizationId:
		return "OrganizationManagementAuthorizationId<enum AuthorizationId>"
	case ShowStaffAuthorizationId:
		return "ShowStaffAuthorizationId<enum AuthorizationId>"
	case EditStaffAuthorizationId:
		return "EditStaffAuthorizationId<enum AuthorizationId>"
	case ShowStaffRoleAuthorizationId:
		return "ShowStaffRoleAuthorizationId<enum AuthorizationId>"
	case EditStaffRoleAuthorizationId:
		return "EditStaffRoleAuthorizationId<enum AuthorizationId>"
	case CommunityManagementAuthorizationId:
		return "CommunityManagementAuthorizationId<enum AuthorizationId>"
	case ShowGroupAuthorizationId:
		return "ShowGroupAuthorizationId<enum AuthorizationId>"
	case EditGroupAuthorizationId:
		return "EditGroupAuthorizationId<enum AuthorizationId>"
	case GroupBuyingManagementAuthorizationId:
		return "GroupBuyingManagementAuthorizationId<enum AuthorizationId>"
	case ShowTaskAuthorizationId:
		return "ShowTaskAuthorizationId<enum AuthorizationId>"
	case EditTaskAuthorizationId:
		return "EditTaskAuthorizationId<enum AuthorizationId>"
	default:
		return "UNKNOWN_Name_<AuthorizationId>"
	}
}

// 组织成员权限
type StaffRoleAuthorization struct {
	Authorization
	IsOwn bool `db:"IsOwn"`
}

func NewStaffRoleAuthorization() *StaffRoleAuthorization {
	return &StaffRoleAuthorization{}
}

// 成员组
type StaffRoleAuthorizationGroup struct {
	ModuleAuthorization *StaffRoleAuthorization                     `db:"ModuleAuthorization"`
	SubAuthorizations   map[AuthorizationId]*StaffRoleAuthorization `db:"SubAuthorizations"`
}

func NewStaffRoleAuthorizationGroup() *StaffRoleAuthorizationGroup {
	return &StaffRoleAuthorizationGroup{
		ModuleAuthorization: NewStaffRoleAuthorization(),
		SubAuthorizations:   make(map[AuthorizationId]*StaffRoleAuthorization),
	}
}
