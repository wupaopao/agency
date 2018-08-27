package cidl

type AckOrgOrganizationList struct {
	Count uint32          `db:"Count"`
	List  []*Organization `db:"List"`
}

func NewAckOrgOrganizationList() *AckOrgOrganizationList {
	return &AckOrgOrganizationList{
		List: make([]*Organization, 0),
	}
}

type MetaApiOrgOrganizationList struct {
}

var META_ORG_ORGANIZATION_LIST = &MetaApiOrgOrganizationList{}

func (m *MetaApiOrgOrganizationList) GetMethod() string { return "GET" }
func (m *MetaApiOrgOrganizationList) GetURL() string    { return "/agency/org/organization/list" }
func (m *MetaApiOrgOrganizationList) GetName() string   { return "OrgOrganizationList" }
func (m *MetaApiOrgOrganizationList) GetType() string   { return "json" }

// 获取团购组织列表
type ApiOrgOrganizationList struct {
	MetaApiOrgOrganizationList
	Ack   *AckOrgOrganizationList
	Query struct {
		Page     uint32 `form:"page" binding:"required,gt=0" db:"Page"`
		PageSize uint32 `form:"page_size" binding:"required,gt=0,lt=50" db:"PageSize"`
		Search   string `form:"search" db:"Search"`
	}
}

func (m *ApiOrgOrganizationList) GetQuery() interface{}  { return &m.Query }
func (m *ApiOrgOrganizationList) GetParams() interface{} { return nil }
func (m *ApiOrgOrganizationList) GetAsk() interface{}    { return nil }
func (m *ApiOrgOrganizationList) GetAck() interface{}    { return m.Ack }
func MakeApiOrgOrganizationList() ApiOrgOrganizationList {
	return ApiOrgOrganizationList{
		Ack: NewAckOrgOrganizationList(),
	}
}

// 组织
type AckOrgOrganizationInfoByID struct {
	Organization
	ManagerIdCardNumber string `db:"ManagerIdCardNumber"`
	ManagerIdCardFront  string `db:"ManagerIdCardFront"`
	ManagerIdCardBack   string `db:"ManagerIdCardBack"`
	ManagerWxNickname   string `db:"ManagerWxNickname"`
}

func NewAckOrgOrganizationInfoByID() *AckOrgOrganizationInfoByID {
	return &AckOrgOrganizationInfoByID{}
}

type MetaApiOrgOrganizationInfoByID struct {
}

var META_ORG_ORGANIZATION_INFO_BY_ID = &MetaApiOrgOrganizationInfoByID{}

func (m *MetaApiOrgOrganizationInfoByID) GetMethod() string { return "GET" }
func (m *MetaApiOrgOrganizationInfoByID) GetURL() string {
	return "/agency/org/organization/info/:organization_id"
}
func (m *MetaApiOrgOrganizationInfoByID) GetName() string { return "OrgOrganizationInfoByID" }
func (m *MetaApiOrgOrganizationInfoByID) GetType() string { return "json" }

// 获取团购组织详细信息
type ApiOrgOrganizationInfoByID struct {
	MetaApiOrgOrganizationInfoByID
	Ack    *AckOrgOrganizationInfoByID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiOrgOrganizationInfoByID) GetQuery() interface{}  { return nil }
func (m *ApiOrgOrganizationInfoByID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgOrganizationInfoByID) GetAsk() interface{}    { return nil }
func (m *ApiOrgOrganizationInfoByID) GetAck() interface{}    { return m.Ack }
func MakeApiOrgOrganizationInfoByID() ApiOrgOrganizationInfoByID {
	return ApiOrgOrganizationInfoByID{
		Ack: NewAckOrgOrganizationInfoByID(),
	}
}

type MetaApiOrgOrganizationEnableInfoByOrganizationID struct {
}

var META_ORG_ORGANIZATION_ENABLE_INFO_BY_ORGANIZATION_ID = &MetaApiOrgOrganizationEnableInfoByOrganizationID{}

func (m *MetaApiOrgOrganizationEnableInfoByOrganizationID) GetMethod() string { return "GET" }
func (m *MetaApiOrgOrganizationEnableInfoByOrganizationID) GetURL() string {
	return "/agency/org/organization/enable_info/:organization_id"
}
func (m *MetaApiOrgOrganizationEnableInfoByOrganizationID) GetName() string {
	return "OrgOrganizationEnableInfoByOrganizationID"
}
func (m *MetaApiOrgOrganizationEnableInfoByOrganizationID) GetType() string { return "json" }

// 获取可用组织的纪录
type ApiOrgOrganizationEnableInfoByOrganizationID struct {
	MetaApiOrgOrganizationEnableInfoByOrganizationID
	Ack    *Organization
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiOrgOrganizationEnableInfoByOrganizationID) GetQuery() interface{}  { return nil }
func (m *ApiOrgOrganizationEnableInfoByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgOrganizationEnableInfoByOrganizationID) GetAsk() interface{}    { return nil }
func (m *ApiOrgOrganizationEnableInfoByOrganizationID) GetAck() interface{}    { return m.Ack }
func MakeApiOrgOrganizationEnableInfoByOrganizationID() ApiOrgOrganizationEnableInfoByOrganizationID {
	return ApiOrgOrganizationEnableInfoByOrganizationID{
		Ack: NewOrganization(),
	}
}

type AckOrgOrganizationGetOrganizationID struct {
	OrganizationId uint32 `db:"OrganizationId"`
}

func NewAckOrgOrganizationGetOrganizationID() *AckOrgOrganizationGetOrganizationID {
	return &AckOrgOrganizationGetOrganizationID{}
}

type MetaApiOrgOrganizationGetOrganizationID struct {
}

var META_ORG_ORGANIZATION_GET_ORGANIZATION_ID = &MetaApiOrgOrganizationGetOrganizationID{}

func (m *MetaApiOrgOrganizationGetOrganizationID) GetMethod() string { return "GET" }
func (m *MetaApiOrgOrganizationGetOrganizationID) GetURL() string {
	return "/agency/org/organization/get_organization_id"
}
func (m *MetaApiOrgOrganizationGetOrganizationID) GetName() string {
	return "OrgOrganizationGetOrganizationID"
}
func (m *MetaApiOrgOrganizationGetOrganizationID) GetType() string { return "json" }

// 通过uid获取团购组织id
type ApiOrgOrganizationGetOrganizationID struct {
	MetaApiOrgOrganizationGetOrganizationID
	Ack   *AckOrgOrganizationGetOrganizationID
	Query struct {
		UserID string `form:"uid" binding:"required" db:"UserID"`
	}
}

func (m *ApiOrgOrganizationGetOrganizationID) GetQuery() interface{}  { return &m.Query }
func (m *ApiOrgOrganizationGetOrganizationID) GetParams() interface{} { return nil }
func (m *ApiOrgOrganizationGetOrganizationID) GetAsk() interface{}    { return nil }
func (m *ApiOrgOrganizationGetOrganizationID) GetAck() interface{}    { return m.Ack }
func MakeApiOrgOrganizationGetOrganizationID() ApiOrgOrganizationGetOrganizationID {
	return ApiOrgOrganizationGetOrganizationID{
		Ack: NewAckOrgOrganizationGetOrganizationID(),
	}
}

type AckOrgRoleListByOrganizationID struct {
	Count uint32       `db:"Count"`
	List  []*StaffRole `db:"List"`
}

func NewAckOrgRoleListByOrganizationID() *AckOrgRoleListByOrganizationID {
	return &AckOrgRoleListByOrganizationID{
		List: make([]*StaffRole, 0),
	}
}

type MetaApiOrgRoleListByOrganizationID struct {
}

var META_ORG_ROLE_LIST_BY_ORGANIZATION_ID = &MetaApiOrgRoleListByOrganizationID{}

func (m *MetaApiOrgRoleListByOrganizationID) GetMethod() string { return "GET" }
func (m *MetaApiOrgRoleListByOrganizationID) GetURL() string {
	return "/agency/org/role/list/:organization_id"
}
func (m *MetaApiOrgRoleListByOrganizationID) GetName() string { return "OrgRoleListByOrganizationID" }
func (m *MetaApiOrgRoleListByOrganizationID) GetType() string { return "json" }

// 团购成员角色权限列表
type ApiOrgRoleListByOrganizationID struct {
	MetaApiOrgRoleListByOrganizationID
	Ack    *AckOrgRoleListByOrganizationID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
	Query struct {
		Page     uint32 `form:"page" binding:"required,gt=0" db:"Page"`
		PageSize uint32 `form:"page_size" binding:"required,gt=0,lt=50" db:"PageSize"`
	}
}

func (m *ApiOrgRoleListByOrganizationID) GetQuery() interface{}  { return &m.Query }
func (m *ApiOrgRoleListByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgRoleListByOrganizationID) GetAsk() interface{}    { return nil }
func (m *ApiOrgRoleListByOrganizationID) GetAck() interface{}    { return m.Ack }
func MakeApiOrgRoleListByOrganizationID() ApiOrgRoleListByOrganizationID {
	return ApiOrgRoleListByOrganizationID{
		Ack: NewAckOrgRoleListByOrganizationID(),
	}
}

type AckOrgStaffListByOrganizationID struct {
	Count uint32   `db:"Count"`
	List  []*Staff `db:"List"`
}

func NewAckOrgStaffListByOrganizationID() *AckOrgStaffListByOrganizationID {
	return &AckOrgStaffListByOrganizationID{
		List: make([]*Staff, 0),
	}
}

type MetaApiOrgStaffListByOrganizationID struct {
}

var META_ORG_STAFF_LIST_BY_ORGANIZATION_ID = &MetaApiOrgStaffListByOrganizationID{}

func (m *MetaApiOrgStaffListByOrganizationID) GetMethod() string { return "GET" }
func (m *MetaApiOrgStaffListByOrganizationID) GetURL() string {
	return "/agency/org/staff/list/:organization_id"
}
func (m *MetaApiOrgStaffListByOrganizationID) GetName() string { return "OrgStaffListByOrganizationID" }
func (m *MetaApiOrgStaffListByOrganizationID) GetType() string { return "json" }

// 获取团购成员列表
type ApiOrgStaffListByOrganizationID struct {
	MetaApiOrgStaffListByOrganizationID
	Ack    *AckOrgStaffListByOrganizationID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
	Query struct {
		Page     uint32 `form:"page" binding:"required,gt=0" db:"Page"`
		PageSize uint32 `form:"page_size" binding:"required,gt=0,lt=50" db:"PageSize"`
	}
}

func (m *ApiOrgStaffListByOrganizationID) GetQuery() interface{}  { return &m.Query }
func (m *ApiOrgStaffListByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgStaffListByOrganizationID) GetAsk() interface{}    { return nil }
func (m *ApiOrgStaffListByOrganizationID) GetAck() interface{}    { return m.Ack }
func MakeApiOrgStaffListByOrganizationID() ApiOrgStaffListByOrganizationID {
	return ApiOrgStaffListByOrganizationID{
		Ack: NewAckOrgStaffListByOrganizationID(),
	}
}

type AskOrgStaffAddByOrganizationID struct {
	Name   string `binding:"required,lte=64" db:"Name"`
	Mobile string `binding:"required,numeric" db:"Mobile"`
	RoleId uint32 `binding:"required,gt=0" db:"RoleId"`
}

func NewAskOrgStaffAddByOrganizationID() *AskOrgStaffAddByOrganizationID {
	return &AskOrgStaffAddByOrganizationID{}
}

type MetaApiOrgStaffAddByOrganizationID struct {
}

var META_ORG_STAFF_ADD_BY_ORGANIZATION_ID = &MetaApiOrgStaffAddByOrganizationID{}

func (m *MetaApiOrgStaffAddByOrganizationID) GetMethod() string { return "POST" }
func (m *MetaApiOrgStaffAddByOrganizationID) GetURL() string {
	return "/agency/org/staff/add/:organization_id"
}
func (m *MetaApiOrgStaffAddByOrganizationID) GetName() string { return "OrgStaffAddByOrganizationID" }
func (m *MetaApiOrgStaffAddByOrganizationID) GetType() string { return "json" }

// 添加团购组织成员
type ApiOrgStaffAddByOrganizationID struct {
	MetaApiOrgStaffAddByOrganizationID
	Ask    *AskOrgStaffAddByOrganizationID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiOrgStaffAddByOrganizationID) GetQuery() interface{}  { return nil }
func (m *ApiOrgStaffAddByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgStaffAddByOrganizationID) GetAsk() interface{}    { return m.Ask }
func (m *ApiOrgStaffAddByOrganizationID) GetAck() interface{}    { return nil }
func MakeApiOrgStaffAddByOrganizationID() ApiOrgStaffAddByOrganizationID {
	return ApiOrgStaffAddByOrganizationID{
		Ask: NewAskOrgStaffAddByOrganizationID(),
	}
}

type AskOrgStaffEditByOrganizationIDByUserID struct {
	Name   string `binding:"required,lte=64" db:"Name"`
	Mobile string `binding:"required,numeric" db:"Mobile"`
	RoleId uint32 `binding:"required,gt=0" db:"RoleId"`
}

func NewAskOrgStaffEditByOrganizationIDByUserID() *AskOrgStaffEditByOrganizationIDByUserID {
	return &AskOrgStaffEditByOrganizationIDByUserID{}
}

type MetaApiOrgStaffEditByOrganizationIDByUserID struct {
}

var META_ORG_STAFF_EDIT_BY_ORGANIZATION_ID_BY_USER_ID = &MetaApiOrgStaffEditByOrganizationIDByUserID{}

func (m *MetaApiOrgStaffEditByOrganizationIDByUserID) GetMethod() string { return "POST" }
func (m *MetaApiOrgStaffEditByOrganizationIDByUserID) GetURL() string {
	return "/agency/org/staff/edit/:organization_id/:user_id"
}
func (m *MetaApiOrgStaffEditByOrganizationIDByUserID) GetName() string {
	return "OrgStaffEditByOrganizationIDByUserID"
}
func (m *MetaApiOrgStaffEditByOrganizationIDByUserID) GetType() string { return "json" }

// 编辑团购组织成员
type ApiOrgStaffEditByOrganizationIDByUserID struct {
	MetaApiOrgStaffEditByOrganizationIDByUserID
	Ask    *AskOrgStaffEditByOrganizationIDByUserID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		UserID         string `form:"user_id" binding:"required" db:"UserID"`
	}
}

func (m *ApiOrgStaffEditByOrganizationIDByUserID) GetQuery() interface{}  { return nil }
func (m *ApiOrgStaffEditByOrganizationIDByUserID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgStaffEditByOrganizationIDByUserID) GetAsk() interface{}    { return m.Ask }
func (m *ApiOrgStaffEditByOrganizationIDByUserID) GetAck() interface{}    { return nil }
func MakeApiOrgStaffEditByOrganizationIDByUserID() ApiOrgStaffEditByOrganizationIDByUserID {
	return ApiOrgStaffEditByOrganizationIDByUserID{
		Ask: NewAskOrgStaffEditByOrganizationIDByUserID(),
	}
}

type AskOrgStaffDisableByOrganizationIDByUserID struct {
	IsDisable bool `db:"IsDisable"`
}

func NewAskOrgStaffDisableByOrganizationIDByUserID() *AskOrgStaffDisableByOrganizationIDByUserID {
	return &AskOrgStaffDisableByOrganizationIDByUserID{}
}

type MetaApiOrgStaffDisableByOrganizationIDByUserID struct {
}

var META_ORG_STAFF_DISABLE_BY_ORGANIZATION_ID_BY_USER_ID = &MetaApiOrgStaffDisableByOrganizationIDByUserID{}

func (m *MetaApiOrgStaffDisableByOrganizationIDByUserID) GetMethod() string { return "POST" }
func (m *MetaApiOrgStaffDisableByOrganizationIDByUserID) GetURL() string {
	return "/agency/org/staff/disable/:organization_id/:user_id"
}
func (m *MetaApiOrgStaffDisableByOrganizationIDByUserID) GetName() string {
	return "OrgStaffDisableByOrganizationIDByUserID"
}
func (m *MetaApiOrgStaffDisableByOrganizationIDByUserID) GetType() string { return "json" }

// 禁用团购组织成员
type ApiOrgStaffDisableByOrganizationIDByUserID struct {
	MetaApiOrgStaffDisableByOrganizationIDByUserID
	Ask    *AskOrgStaffDisableByOrganizationIDByUserID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		UserID         string `form:"user_id" binding:"required" db:"UserID"`
	}
}

func (m *ApiOrgStaffDisableByOrganizationIDByUserID) GetQuery() interface{}  { return nil }
func (m *ApiOrgStaffDisableByOrganizationIDByUserID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgStaffDisableByOrganizationIDByUserID) GetAsk() interface{}    { return m.Ask }
func (m *ApiOrgStaffDisableByOrganizationIDByUserID) GetAck() interface{}    { return nil }
func MakeApiOrgStaffDisableByOrganizationIDByUserID() ApiOrgStaffDisableByOrganizationIDByUserID {
	return ApiOrgStaffDisableByOrganizationIDByUserID{
		Ask: NewAskOrgStaffDisableByOrganizationIDByUserID(),
	}
}

type AckOrgAuthorizationList struct {
	Modules []*StaffRoleAuthorizationGroup `db:"Modules"`
}

func NewAckOrgAuthorizationList() *AckOrgAuthorizationList {
	return &AckOrgAuthorizationList{
		Modules: make([]*StaffRoleAuthorizationGroup, 0),
	}
}

type MetaApiOrgAuthorizationList struct {
}

var META_ORG_AUTHORIZATION_LIST = &MetaApiOrgAuthorizationList{}

func (m *MetaApiOrgAuthorizationList) GetMethod() string { return "GET" }
func (m *MetaApiOrgAuthorizationList) GetURL() string    { return "/agency/org/authorization/list" }
func (m *MetaApiOrgAuthorizationList) GetName() string   { return "OrgAuthorizationList" }
func (m *MetaApiOrgAuthorizationList) GetType() string   { return "json" }

// 权限列表
type ApiOrgAuthorizationList struct {
	MetaApiOrgAuthorizationList
	Ack *AckOrgAuthorizationList
}

func (m *ApiOrgAuthorizationList) GetQuery() interface{}  { return nil }
func (m *ApiOrgAuthorizationList) GetParams() interface{} { return nil }
func (m *ApiOrgAuthorizationList) GetAsk() interface{}    { return nil }
func (m *ApiOrgAuthorizationList) GetAck() interface{}    { return m.Ack }
func MakeApiOrgAuthorizationList() ApiOrgAuthorizationList {
	return ApiOrgAuthorizationList{
		Ack: NewAckOrgAuthorizationList(),
	}
}

type AckOrgRoleGetByOrganizationIDByRoleID struct {
	StaffRole *StaffRole                     `db:"StaffRole"`
	Modules   []*StaffRoleAuthorizationGroup `db:"Modules"`
}

func NewAckOrgRoleGetByOrganizationIDByRoleID() *AckOrgRoleGetByOrganizationIDByRoleID {
	return &AckOrgRoleGetByOrganizationIDByRoleID{
		StaffRole: NewStaffRole(),
		Modules:   make([]*StaffRoleAuthorizationGroup, 0),
	}
}

type MetaApiOrgRoleGetByOrganizationIDByRoleID struct {
}

var META_ORG_ROLE_GET_BY_ORGANIZATION_ID_BY_ROLE_ID = &MetaApiOrgRoleGetByOrganizationIDByRoleID{}

func (m *MetaApiOrgRoleGetByOrganizationIDByRoleID) GetMethod() string { return "GET" }
func (m *MetaApiOrgRoleGetByOrganizationIDByRoleID) GetURL() string {
	return "/agency/org/role/get/:organization_id/:role_id"
}
func (m *MetaApiOrgRoleGetByOrganizationIDByRoleID) GetName() string {
	return "OrgRoleGetByOrganizationIDByRoleID"
}
func (m *MetaApiOrgRoleGetByOrganizationIDByRoleID) GetType() string { return "json" }

// 获取角色
type ApiOrgRoleGetByOrganizationIDByRoleID struct {
	MetaApiOrgRoleGetByOrganizationIDByRoleID
	Ack    *AckOrgRoleGetByOrganizationIDByRoleID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		RoleID         uint32 `form:"role_id" binding:"required,gt=0" db:"RoleID"`
	}
}

func (m *ApiOrgRoleGetByOrganizationIDByRoleID) GetQuery() interface{}  { return nil }
func (m *ApiOrgRoleGetByOrganizationIDByRoleID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgRoleGetByOrganizationIDByRoleID) GetAsk() interface{}    { return nil }
func (m *ApiOrgRoleGetByOrganizationIDByRoleID) GetAck() interface{}    { return m.Ack }
func MakeApiOrgRoleGetByOrganizationIDByRoleID() ApiOrgRoleGetByOrganizationIDByRoleID {
	return ApiOrgRoleGetByOrganizationIDByRoleID{
		Ack: NewAckOrgRoleGetByOrganizationIDByRoleID(),
	}
}

type AskOrgRoleAddByOrganizationID struct {
	RoleName         string   `binding:"required,lte=64" db:"RoleName"`
	AuthorizationIds []uint32 `binding:"required,gt=0" db:"AuthorizationIds"`
}

func NewAskOrgRoleAddByOrganizationID() *AskOrgRoleAddByOrganizationID {
	return &AskOrgRoleAddByOrganizationID{
		AuthorizationIds: make([]uint32, 0),
	}
}

type MetaApiOrgRoleAddByOrganizationID struct {
}

var META_ORG_ROLE_ADD_BY_ORGANIZATION_ID = &MetaApiOrgRoleAddByOrganizationID{}

func (m *MetaApiOrgRoleAddByOrganizationID) GetMethod() string { return "POST" }
func (m *MetaApiOrgRoleAddByOrganizationID) GetURL() string {
	return "/agency/org/role/add/:organization_id"
}
func (m *MetaApiOrgRoleAddByOrganizationID) GetName() string { return "OrgRoleAddByOrganizationID" }
func (m *MetaApiOrgRoleAddByOrganizationID) GetType() string { return "json" }

// 添加角色
type ApiOrgRoleAddByOrganizationID struct {
	MetaApiOrgRoleAddByOrganizationID
	Ask    *AskOrgRoleAddByOrganizationID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiOrgRoleAddByOrganizationID) GetQuery() interface{}  { return nil }
func (m *ApiOrgRoleAddByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgRoleAddByOrganizationID) GetAsk() interface{}    { return m.Ask }
func (m *ApiOrgRoleAddByOrganizationID) GetAck() interface{}    { return nil }
func MakeApiOrgRoleAddByOrganizationID() ApiOrgRoleAddByOrganizationID {
	return ApiOrgRoleAddByOrganizationID{
		Ask: NewAskOrgRoleAddByOrganizationID(),
	}
}

type AskOrgRoleEditByOrganizationIDByRoleID struct {
	RoleName         string   `binding:"required,lte=64" db:"RoleName"`
	AuthorizationIds []uint32 `binding:"required,gt=0" db:"AuthorizationIds"`
}

func NewAskOrgRoleEditByOrganizationIDByRoleID() *AskOrgRoleEditByOrganizationIDByRoleID {
	return &AskOrgRoleEditByOrganizationIDByRoleID{
		AuthorizationIds: make([]uint32, 0),
	}
}

type MetaApiOrgRoleEditByOrganizationIDByRoleID struct {
}

var META_ORG_ROLE_EDIT_BY_ORGANIZATION_ID_BY_ROLE_ID = &MetaApiOrgRoleEditByOrganizationIDByRoleID{}

func (m *MetaApiOrgRoleEditByOrganizationIDByRoleID) GetMethod() string { return "POST" }
func (m *MetaApiOrgRoleEditByOrganizationIDByRoleID) GetURL() string {
	return "/agency/org/role/edit/:organization_id/:role_id"
}
func (m *MetaApiOrgRoleEditByOrganizationIDByRoleID) GetName() string {
	return "OrgRoleEditByOrganizationIDByRoleID"
}
func (m *MetaApiOrgRoleEditByOrganizationIDByRoleID) GetType() string { return "json" }

// 编辑角色
type ApiOrgRoleEditByOrganizationIDByRoleID struct {
	MetaApiOrgRoleEditByOrganizationIDByRoleID
	Ask    *AskOrgRoleEditByOrganizationIDByRoleID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		RoleID         uint32 `form:"role_id" binding:"required,gt=0" db:"RoleID"`
	}
}

func (m *ApiOrgRoleEditByOrganizationIDByRoleID) GetQuery() interface{}  { return nil }
func (m *ApiOrgRoleEditByOrganizationIDByRoleID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgRoleEditByOrganizationIDByRoleID) GetAsk() interface{}    { return m.Ask }
func (m *ApiOrgRoleEditByOrganizationIDByRoleID) GetAck() interface{}    { return nil }
func MakeApiOrgRoleEditByOrganizationIDByRoleID() ApiOrgRoleEditByOrganizationIDByRoleID {
	return ApiOrgRoleEditByOrganizationIDByRoleID{
		Ask: NewAskOrgRoleEditByOrganizationIDByRoleID(),
	}
}

type AskOrgRoleDisableByOrganizationIDByRoleID struct {
	IsDisable bool `db:"IsDisable"`
}

func NewAskOrgRoleDisableByOrganizationIDByRoleID() *AskOrgRoleDisableByOrganizationIDByRoleID {
	return &AskOrgRoleDisableByOrganizationIDByRoleID{}
}

type MetaApiOrgRoleDisableByOrganizationIDByRoleID struct {
}

var META_ORG_ROLE_DISABLE_BY_ORGANIZATION_ID_BY_ROLE_ID = &MetaApiOrgRoleDisableByOrganizationIDByRoleID{}

func (m *MetaApiOrgRoleDisableByOrganizationIDByRoleID) GetMethod() string { return "POST" }
func (m *MetaApiOrgRoleDisableByOrganizationIDByRoleID) GetURL() string {
	return "/agency/org/role/disable/:organization_id/:role_id"
}
func (m *MetaApiOrgRoleDisableByOrganizationIDByRoleID) GetName() string {
	return "OrgRoleDisableByOrganizationIDByRoleID"
}
func (m *MetaApiOrgRoleDisableByOrganizationIDByRoleID) GetType() string { return "json" }

// 禁用角色
type ApiOrgRoleDisableByOrganizationIDByRoleID struct {
	MetaApiOrgRoleDisableByOrganizationIDByRoleID
	Ask    *AskOrgRoleDisableByOrganizationIDByRoleID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		RoleID         uint32 `form:"role_id" binding:"required,gt=0" db:"RoleID"`
	}
}

func (m *ApiOrgRoleDisableByOrganizationIDByRoleID) GetQuery() interface{}  { return nil }
func (m *ApiOrgRoleDisableByOrganizationIDByRoleID) GetParams() interface{} { return &m.Params }
func (m *ApiOrgRoleDisableByOrganizationIDByRoleID) GetAsk() interface{}    { return m.Ask }
func (m *ApiOrgRoleDisableByOrganizationIDByRoleID) GetAck() interface{}    { return nil }
func MakeApiOrgRoleDisableByOrganizationIDByRoleID() ApiOrgRoleDisableByOrganizationIDByRoleID {
	return ApiOrgRoleDisableByOrganizationIDByRoleID{
		Ask: NewAskOrgRoleDisableByOrganizationIDByRoleID(),
	}
}
