package agency

import "time"

type StaffRoleType int

const (
	StaffRoleTypeSuperAdministrator StaffRoleType = 1
	StaffRoleTypeCommon             StaffRoleType = 2
)

func (m StaffRoleType) String() string {
	switch m {

	case StaffRoleTypeSuperAdministrator:
		return "StaffRoleTypeSuperAdministrator<enum StaffRoleType>"
	case StaffRoleTypeCommon:
		return "StaffRoleTypeCommon<enum StaffRoleType>"
	default:
		return "UNKNOWN_Name_<StaffRoleType>"
	}
}

// 团购组织成员角色
type StaffRole struct {
	RoleId            uint32                `db:"rol_id"`
	OrganizationId    uint32                `db:"org_id"`
	RoleName          string                `db:"name"`
	RoleAuthorization *RoleAuthorizationMap `db:"authorization"`
	IsDisable         bool                  `db:"is_disable"`
	Type              StaffRoleType         `db:"type"`
	Version           uint32                `db:"version"`
}

func NewStaffRole() *StaffRole {
	return &StaffRole{
		RoleAuthorization: NewRoleAuthorizationMap(),
	}
}

// 团购组织成员
type Staff struct {
	UserId           string    `db:"uid"`
	OrganizationId   uint32    `db:"org_id"`
	OrganizationName string    `db:"org_name"`
	Name             string    `db:"name"`
	Mobile           string    `db:"mobile"`
	RoleId           uint32    `db:"rol_id"`
	RoleName         string    `db:"rol_name"`
	IsDisable        bool      `db:"is_disable"`
	CreateTime       time.Time `db:"create_time"`
}

func NewStaff() *Staff {
	return &Staff{}
}

// 组织信息完善状态
type OrganizationPerfectionState int

const (
	// 默认
	OrganizationPerfectionStateDefault OrganizationPerfectionState = 0
	// 已完善资料
	OrganizationPerfectionStateComplete OrganizationPerfectionState = 1
	// 需完善资料
	OrganizationPerfectionStateNeedComplete OrganizationPerfectionState = 2
)

func (m OrganizationPerfectionState) String() string {
	switch m {

	case OrganizationPerfectionStateDefault:
		return "OrganizationPerfectionStateDefault<enum OrganizationPerfectionState>"
	case OrganizationPerfectionStateComplete:
		return "OrganizationPerfectionStateComplete<enum OrganizationPerfectionState>"
	case OrganizationPerfectionStateNeedComplete:
		return "OrganizationPerfectionStateNeedComplete<enum OrganizationPerfectionState>"
	default:
		return "UNKNOWN_Name_<OrganizationPerfectionState>"
	}
}

// 团购模式
type OrganizationGroupBuyingMode int

const (
	// 报单模式
	OrganizationGroupBuyingModeReport OrganizationGroupBuyingMode = 1
	// 下单模式
	OrganizationGroupBuyingModeOrder OrganizationGroupBuyingMode = 2
)

func (m OrganizationGroupBuyingMode) String() string {
	switch m {

	case OrganizationGroupBuyingModeReport:
		return "OrganizationGroupBuyingModeReport<enum OrganizationGroupBuyingMode>"
	case OrganizationGroupBuyingModeOrder:
		return "OrganizationGroupBuyingModeOrder<enum OrganizationGroupBuyingMode>"
	default:
		return "UNKNOWN_Name_<OrganizationGroupBuyingMode>"
	}
}

// 团购组织
type Organization struct {
	OrganizationId  uint32                      `db:"org_id"`
	Name            string                      `db:"name"`
	Logo            string                      `db:"logo"`
	Province        string                      `db:"province"`
	City            string                      `db:"city"`
	Address         string                      `db:"address"`
	PostCode        string                      `db:"post_code"`
	BankName        string                      `db:"bank_name"`
	BankAccount     string                      `db:"bank_account"`
	BankAccountName string                      `db:"bank_account_name"`
	CompanyName     string                      `db:"company_name"`
	LicenseNumber   string                      `db:"license_number"`
	LicensePicture  string                      `db:"license_picture"`
	ManagerUserId   string                      `db:"manager_uid"`
	ManagerName     string                      `db:"manager_name"`
	ManagerMobile   string                      `db:"manager_mobile"`
	GroupBuyingMode OrganizationGroupBuyingMode `db:"group_buying_mode"`
	PerfectionState OrganizationPerfectionState `db:"perfection_state"`
	IsDisable       bool                        `db:"is_disable"`
	CreateTime      time.Time                   `db:"create_time"`
}

func NewOrganization() *Organization {
	return &Organization{}
}
