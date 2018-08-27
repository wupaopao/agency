package cidl

type AckAdminOrganizationList struct {
	Count uint32          `db:"Count"`
	List  []*Organization `db:"List"`
}

func NewAckAdminOrganizationList() *AckAdminOrganizationList {
	return &AckAdminOrganizationList{
		List: make([]*Organization, 0),
	}
}

type MetaApiAdminOrganizationList struct {
}

var META_ADMIN_ORGANIZATION_LIST = &MetaApiAdminOrganizationList{}

func (m *MetaApiAdminOrganizationList) GetMethod() string { return "GET" }
func (m *MetaApiAdminOrganizationList) GetURL() string    { return "/agency/admin/organization/list" }
func (m *MetaApiAdminOrganizationList) GetName() string   { return "AdminOrganizationList" }
func (m *MetaApiAdminOrganizationList) GetType() string   { return "json" }

// 获取团购组织列表
type ApiAdminOrganizationList struct {
	MetaApiAdminOrganizationList
	Ack   *AckAdminOrganizationList
	Query struct {
		Page     uint32 `form:"page" binding:"required,gt=0" db:"Page"`
		PageSize uint32 `form:"page_size" binding:"required,gt=0,lt=50" db:"PageSize"`
		Search   string `form:"search" db:"Search"`
	}
}

func (m *ApiAdminOrganizationList) GetQuery() interface{}  { return &m.Query }
func (m *ApiAdminOrganizationList) GetParams() interface{} { return nil }
func (m *ApiAdminOrganizationList) GetAsk() interface{}    { return nil }
func (m *ApiAdminOrganizationList) GetAck() interface{}    { return m.Ack }
func MakeApiAdminOrganizationList() ApiAdminOrganizationList {
	return ApiAdminOrganizationList{
		Ack: NewAckAdminOrganizationList(),
	}
}

// 组织
type AckAdminOrganizationInfoByID struct {
	Organization
	ManagerIdCardNumber string `db:"ManagerIdCardNumber"`
	ManagerIdCardFront  string `db:"ManagerIdCardFront"`
	ManagerIdCardBack   string `db:"ManagerIdCardBack"`
	ManagerWxNickname   string `db:"ManagerWxNickname"`
}

func NewAckAdminOrganizationInfoByID() *AckAdminOrganizationInfoByID {
	return &AckAdminOrganizationInfoByID{}
}

type MetaApiAdminOrganizationInfoByID struct {
}

var META_ADMIN_ORGANIZATION_INFO_BY_ID = &MetaApiAdminOrganizationInfoByID{}

func (m *MetaApiAdminOrganizationInfoByID) GetMethod() string { return "GET" }
func (m *MetaApiAdminOrganizationInfoByID) GetURL() string {
	return "/agency/admin/organization/info/:organization_id"
}
func (m *MetaApiAdminOrganizationInfoByID) GetName() string { return "AdminOrganizationInfoByID" }
func (m *MetaApiAdminOrganizationInfoByID) GetType() string { return "json" }

// 获取团购组织详细信息
type ApiAdminOrganizationInfoByID struct {
	MetaApiAdminOrganizationInfoByID
	Ack    *AckAdminOrganizationInfoByID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiAdminOrganizationInfoByID) GetQuery() interface{}  { return nil }
func (m *ApiAdminOrganizationInfoByID) GetParams() interface{} { return &m.Params }
func (m *ApiAdminOrganizationInfoByID) GetAsk() interface{}    { return nil }
func (m *ApiAdminOrganizationInfoByID) GetAck() interface{}    { return m.Ack }
func MakeApiAdminOrganizationInfoByID() ApiAdminOrganizationInfoByID {
	return ApiAdminOrganizationInfoByID{
		Ack: NewAckAdminOrganizationInfoByID(),
	}
}

// 组织图片
type AckPicToken struct {
	OriginalFileName string `db:"OriginalFileName"`
	Token            string `db:"Token"`
	Key              string `db:"Key"`
	StoreUrl         string `db:"StoreUrl"`
	AccessUrl        string `db:"AccessUrl"`
}

func NewAckPicToken() *AckPicToken {
	return &AckPicToken{}
}

type AskAdminOrganizationAddPicToken struct {
	FileNames []string `db:"FileNames"`
}

func NewAskAdminOrganizationAddPicToken() *AskAdminOrganizationAddPicToken {
	return &AskAdminOrganizationAddPicToken{
		FileNames: make([]string, 0),
	}
}

type AckAdminOrganizationAddPicToken struct {
	Tokens []*AckPicToken `db:"Tokens"`
}

func NewAckAdminOrganizationAddPicToken() *AckAdminOrganizationAddPicToken {
	return &AckAdminOrganizationAddPicToken{
		Tokens: make([]*AckPicToken, 0),
	}
}

type MetaApiAdminOrganizationAddPicToken struct {
}

var META_ADMIN_ORGANIZATION_ADD_PIC_TOKEN = &MetaApiAdminOrganizationAddPicToken{}

func (m *MetaApiAdminOrganizationAddPicToken) GetMethod() string { return "POST" }
func (m *MetaApiAdminOrganizationAddPicToken) GetURL() string {
	return "/agency/admin/organization/add/pic_token"
}
func (m *MetaApiAdminOrganizationAddPicToken) GetName() string { return "AdminOrganizationAddPicToken" }
func (m *MetaApiAdminOrganizationAddPicToken) GetType() string { return "json" }

type ApiAdminOrganizationAddPicToken struct {
	MetaApiAdminOrganizationAddPicToken
	Ask *AskAdminOrganizationAddPicToken
	Ack *AckAdminOrganizationAddPicToken
}

func (m *ApiAdminOrganizationAddPicToken) GetQuery() interface{}  { return nil }
func (m *ApiAdminOrganizationAddPicToken) GetParams() interface{} { return nil }
func (m *ApiAdminOrganizationAddPicToken) GetAsk() interface{}    { return m.Ask }
func (m *ApiAdminOrganizationAddPicToken) GetAck() interface{}    { return m.Ack }
func MakeApiAdminOrganizationAddPicToken() ApiAdminOrganizationAddPicToken {
	return ApiAdminOrganizationAddPicToken{
		Ask: NewAskAdminOrganizationAddPicToken(),
		Ack: NewAckAdminOrganizationAddPicToken(),
	}
}

type AskAdminOrganizationAdd struct {
	ManagerMobile       string                      `binding:"required,numeric" db:"ManagerMobile"`
	ManagerNickname     string                      `binding:"required,lte=64" db:"ManagerNickname"`
	Name                string                      `binding:"required,lte=128" db:"Name"`
	Logo                string                      `binding:"required,lte=255" db:"Logo"`
	CompanyName         string                      `binding:"required,lte=128" db:"CompanyName"`
	BankAccountName     string                      `binding:"required,lte=64" db:"BankAccountName"`
	BankName            string                      `binding:"required,lte=128" db:"BankName"`
	BankAccount         string                      `binding:"required,lte=20" db:"BankAccount"`
	Province            string                      `binding:"required,lte=40" db:"Province"`
	City                string                      `binding:"required,lte=40" db:"City"`
	LicenseNumber       string                      `binding:"required,lte=20" db:"LicenseNumber"`
	LicensePicture      string                      `binding:"required,lte=255" db:"LicensePicture"`
	Address             string                      `binding:"required,lte=255" db:"Address"`
	PostCode            string                      `binding:"numeric,len=6" db:"PostCode"`
	ManagerName         string                      `binding:"required,lte=64" db:"ManagerName"`
	ManagerIdCardNumber string                      `binding:"required,lte=18" db:"ManagerIdCardNumber"`
	ManagerIdCardFront  string                      `binding:"required,lte=255" db:"ManagerIdCardFront"`
	ManagerIdCardBack   string                      `binding:"required,lte=255" db:"ManagerIdCardBack"`
	GroupBuyingMode     OrganizationGroupBuyingMode `binding:"required" db:"GroupBuyingMode"`
}

func NewAskAdminOrganizationAdd() *AskAdminOrganizationAdd {
	return &AskAdminOrganizationAdd{}
}

type MetaApiAdminOrganizationAdd struct {
}

var META_ADMIN_ORGANIZATION_ADD = &MetaApiAdminOrganizationAdd{}

func (m *MetaApiAdminOrganizationAdd) GetMethod() string { return "POST" }
func (m *MetaApiAdminOrganizationAdd) GetURL() string    { return "/agency/admin/organization/add" }
func (m *MetaApiAdminOrganizationAdd) GetName() string   { return "AdminOrganizationAdd" }
func (m *MetaApiAdminOrganizationAdd) GetType() string   { return "json" }

// 添加组织
type ApiAdminOrganizationAdd struct {
	MetaApiAdminOrganizationAdd
	Ask *AskAdminOrganizationAdd
}

func (m *ApiAdminOrganizationAdd) GetQuery() interface{}  { return nil }
func (m *ApiAdminOrganizationAdd) GetParams() interface{} { return nil }
func (m *ApiAdminOrganizationAdd) GetAsk() interface{}    { return m.Ask }
func (m *ApiAdminOrganizationAdd) GetAck() interface{}    { return nil }
func MakeApiAdminOrganizationAdd() ApiAdminOrganizationAdd {
	return ApiAdminOrganizationAdd{
		Ask: NewAskAdminOrganizationAdd(),
	}
}

type AskAdminOrganizationEditByOrganizationID struct {
	ManagerMobile       string                      `binding:"required,numeric" db:"ManagerMobile"`
	ManagerNickname     string                      `binding:"required,lte=64" db:"ManagerNickname"`
	Name                string                      `binding:"required,lte=128" db:"Name"`
	Logo                string                      `binding:"required,lte=255" db:"Logo"`
	CompanyName         string                      `binding:"required,lte=128" db:"CompanyName"`
	BankAccountName     string                      `binding:"required,lte=64" db:"BankAccountName"`
	BankName            string                      `binding:"required,lte=128" db:"BankName"`
	BankAccount         string                      `binding:"required,lte=20" db:"BankAccount"`
	Province            string                      `binding:"required,lte=40" db:"Province"`
	City                string                      `binding:"required,lte=40" db:"City"`
	LicenseNumber       string                      `binding:"required,lte=20" db:"LicenseNumber"`
	LicensePicture      string                      `binding:"required,lte=255" db:"LicensePicture"`
	Address             string                      `binding:"required,lte=255" db:"Address"`
	PostCode            string                      `binding:"numeric,len=6" db:"PostCode"`
	ManagerName         string                      `binding:"required,lte=64" db:"ManagerName"`
	ManagerIdCardNumber string                      `binding:"required,lte=18" db:"ManagerIdCardNumber"`
	ManagerIdCardFront  string                      `binding:"required,lte=255" db:"ManagerIdCardFront"`
	ManagerIdCardBack   string                      `binding:"required,lte=255" db:"ManagerIdCardBack"`
	GroupBuyingMode     OrganizationGroupBuyingMode `binding:"required" db:"GroupBuyingMode"`
}

func NewAskAdminOrganizationEditByOrganizationID() *AskAdminOrganizationEditByOrganizationID {
	return &AskAdminOrganizationEditByOrganizationID{}
}

type MetaApiAdminOrganizationEditByOrganizationID struct {
}

var META_ADMIN_ORGANIZATION_EDIT_BY_ORGANIZATION_ID = &MetaApiAdminOrganizationEditByOrganizationID{}

func (m *MetaApiAdminOrganizationEditByOrganizationID) GetMethod() string { return "POST" }
func (m *MetaApiAdminOrganizationEditByOrganizationID) GetURL() string {
	return "/agency/admin/organization/edit/:organization_id"
}
func (m *MetaApiAdminOrganizationEditByOrganizationID) GetName() string {
	return "AdminOrganizationEditByOrganizationID"
}
func (m *MetaApiAdminOrganizationEditByOrganizationID) GetType() string { return "json" }

// 编辑组织
type ApiAdminOrganizationEditByOrganizationID struct {
	MetaApiAdminOrganizationEditByOrganizationID
	Ask    *AskAdminOrganizationEditByOrganizationID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiAdminOrganizationEditByOrganizationID) GetQuery() interface{}  { return nil }
func (m *ApiAdminOrganizationEditByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiAdminOrganizationEditByOrganizationID) GetAsk() interface{}    { return m.Ask }
func (m *ApiAdminOrganizationEditByOrganizationID) GetAck() interface{}    { return nil }
func MakeApiAdminOrganizationEditByOrganizationID() ApiAdminOrganizationEditByOrganizationID {
	return ApiAdminOrganizationEditByOrganizationID{
		Ask: NewAskAdminOrganizationEditByOrganizationID(),
	}
}

type AskAdminOrganizationDisableByOrganizationID struct {
	IsDisable bool `db:"IsDisable"`
}

func NewAskAdminOrganizationDisableByOrganizationID() *AskAdminOrganizationDisableByOrganizationID {
	return &AskAdminOrganizationDisableByOrganizationID{}
}

type MetaApiAdminOrganizationDisableByOrganizationID struct {
}

var META_ADMIN_ORGANIZATION_DISABLE_BY_ORGANIZATION_ID = &MetaApiAdminOrganizationDisableByOrganizationID{}

func (m *MetaApiAdminOrganizationDisableByOrganizationID) GetMethod() string { return "POST" }
func (m *MetaApiAdminOrganizationDisableByOrganizationID) GetURL() string {
	return "/agency/admin/organization/disable/:organization_id"
}
func (m *MetaApiAdminOrganizationDisableByOrganizationID) GetName() string {
	return "AdminOrganizationDisableByOrganizationID"
}
func (m *MetaApiAdminOrganizationDisableByOrganizationID) GetType() string { return "json" }

// 禁用组织
type ApiAdminOrganizationDisableByOrganizationID struct {
	MetaApiAdminOrganizationDisableByOrganizationID
	Ask    *AskAdminOrganizationDisableByOrganizationID
	Params struct {
		OrganizationID string `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiAdminOrganizationDisableByOrganizationID) GetQuery() interface{}  { return nil }
func (m *ApiAdminOrganizationDisableByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiAdminOrganizationDisableByOrganizationID) GetAsk() interface{}    { return m.Ask }
func (m *ApiAdminOrganizationDisableByOrganizationID) GetAck() interface{}    { return nil }
func MakeApiAdminOrganizationDisableByOrganizationID() ApiAdminOrganizationDisableByOrganizationID {
	return ApiAdminOrganizationDisableByOrganizationID{
		Ask: NewAskAdminOrganizationDisableByOrganizationID(),
	}
}

type MetaApiAdminOrganizationEnableInfoByOrganizationID struct {
}

var META_ADMIN_ORGANIZATION_ENABLE_INFO_BY_ORGANIZATION_ID = &MetaApiAdminOrganizationEnableInfoByOrganizationID{}

func (m *MetaApiAdminOrganizationEnableInfoByOrganizationID) GetMethod() string { return "GET" }
func (m *MetaApiAdminOrganizationEnableInfoByOrganizationID) GetURL() string {
	return "/agency/admin/organization/enable_info/:organization_id"
}
func (m *MetaApiAdminOrganizationEnableInfoByOrganizationID) GetName() string {
	return "AdminOrganizationEnableInfoByOrganizationID"
}
func (m *MetaApiAdminOrganizationEnableInfoByOrganizationID) GetType() string { return "json" }

// 获取可用组织的纪录
type ApiAdminOrganizationEnableInfoByOrganizationID struct {
	MetaApiAdminOrganizationEnableInfoByOrganizationID
	Ack    *Organization
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiAdminOrganizationEnableInfoByOrganizationID) GetQuery() interface{}  { return nil }
func (m *ApiAdminOrganizationEnableInfoByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiAdminOrganizationEnableInfoByOrganizationID) GetAsk() interface{}    { return nil }
func (m *ApiAdminOrganizationEnableInfoByOrganizationID) GetAck() interface{}    { return m.Ack }
func MakeApiAdminOrganizationEnableInfoByOrganizationID() ApiAdminOrganizationEnableInfoByOrganizationID {
	return ApiAdminOrganizationEnableInfoByOrganizationID{
		Ack: NewOrganization(),
	}
}

type AckAdminOrganizationGetOrganizationID struct {
	OrganizationId uint32 `db:"OrganizationId"`
}

func NewAckAdminOrganizationGetOrganizationID() *AckAdminOrganizationGetOrganizationID {
	return &AckAdminOrganizationGetOrganizationID{}
}

type MetaApiAdminOrganizationGetOrganizationID struct {
}

var META_ADMIN_ORGANIZATION_GET_ORGANIZATION_ID = &MetaApiAdminOrganizationGetOrganizationID{}

func (m *MetaApiAdminOrganizationGetOrganizationID) GetMethod() string { return "GET" }
func (m *MetaApiAdminOrganizationGetOrganizationID) GetURL() string {
	return "/agency/admin/organization/get_organization_id"
}
func (m *MetaApiAdminOrganizationGetOrganizationID) GetName() string {
	return "AdminOrganizationGetOrganizationID"
}
func (m *MetaApiAdminOrganizationGetOrganizationID) GetType() string { return "json" }

// 通过uid获取团购组织id
type ApiAdminOrganizationGetOrganizationID struct {
	MetaApiAdminOrganizationGetOrganizationID
	Ack   *AckAdminOrganizationGetOrganizationID
	Query struct {
		UserID string `form:"uid" binding:"required" db:"UserID"`
	}
}

func (m *ApiAdminOrganizationGetOrganizationID) GetQuery() interface{}  { return &m.Query }
func (m *ApiAdminOrganizationGetOrganizationID) GetParams() interface{} { return nil }
func (m *ApiAdminOrganizationGetOrganizationID) GetAsk() interface{}    { return nil }
func (m *ApiAdminOrganizationGetOrganizationID) GetAck() interface{}    { return m.Ack }
func MakeApiAdminOrganizationGetOrganizationID() ApiAdminOrganizationGetOrganizationID {
	return ApiAdminOrganizationGetOrganizationID{
		Ack: NewAckAdminOrganizationGetOrganizationID(),
	}
}

type AckAdminRoleListByOrganizationID struct {
	Count uint32       `db:"Count"`
	List  []*StaffRole `db:"List"`
}

func NewAckAdminRoleListByOrganizationID() *AckAdminRoleListByOrganizationID {
	return &AckAdminRoleListByOrganizationID{
		List: make([]*StaffRole, 0),
	}
}

type MetaApiAdminRoleListByOrganizationID struct {
}

var META_ADMIN_ROLE_LIST_BY_ORGANIZATION_ID = &MetaApiAdminRoleListByOrganizationID{}

func (m *MetaApiAdminRoleListByOrganizationID) GetMethod() string { return "GET" }
func (m *MetaApiAdminRoleListByOrganizationID) GetURL() string {
	return "/agency/admin/role/list/:organization_id"
}
func (m *MetaApiAdminRoleListByOrganizationID) GetName() string {
	return "AdminRoleListByOrganizationID"
}
func (m *MetaApiAdminRoleListByOrganizationID) GetType() string { return "json" }

// 团购成员角色权限列表
type ApiAdminRoleListByOrganizationID struct {
	MetaApiAdminRoleListByOrganizationID
	Ack    *AckAdminRoleListByOrganizationID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
	Query struct {
		Page     uint32 `form:"page" binding:"required,gt=0" db:"Page"`
		PageSize uint32 `form:"page_size" binding:"required,gt=0,lt=50" db:"PageSize"`
	}
}

func (m *ApiAdminRoleListByOrganizationID) GetQuery() interface{}  { return &m.Query }
func (m *ApiAdminRoleListByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiAdminRoleListByOrganizationID) GetAsk() interface{}    { return nil }
func (m *ApiAdminRoleListByOrganizationID) GetAck() interface{}    { return m.Ack }
func MakeApiAdminRoleListByOrganizationID() ApiAdminRoleListByOrganizationID {
	return ApiAdminRoleListByOrganizationID{
		Ack: NewAckAdminRoleListByOrganizationID(),
	}
}

type AckAdminStaffListByOrganizationID struct {
	Count uint32   `db:"Count"`
	List  []*Staff `db:"List"`
}

func NewAckAdminStaffListByOrganizationID() *AckAdminStaffListByOrganizationID {
	return &AckAdminStaffListByOrganizationID{
		List: make([]*Staff, 0),
	}
}

type MetaApiAdminStaffListByOrganizationID struct {
}

var META_ADMIN_STAFF_LIST_BY_ORGANIZATION_ID = &MetaApiAdminStaffListByOrganizationID{}

func (m *MetaApiAdminStaffListByOrganizationID) GetMethod() string { return "GET" }
func (m *MetaApiAdminStaffListByOrganizationID) GetURL() string {
	return "/agency/admin/staff/list/:organization_id"
}
func (m *MetaApiAdminStaffListByOrganizationID) GetName() string {
	return "AdminStaffListByOrganizationID"
}
func (m *MetaApiAdminStaffListByOrganizationID) GetType() string { return "json" }

// 获取团购成员列表
type ApiAdminStaffListByOrganizationID struct {
	MetaApiAdminStaffListByOrganizationID
	Ack    *AckAdminStaffListByOrganizationID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
	Query struct {
		Page     uint32 `form:"page" binding:"required,gt=0" db:"Page"`
		PageSize uint32 `form:"page_size" binding:"required,gt=0,lt=50" db:"PageSize"`
	}
}

func (m *ApiAdminStaffListByOrganizationID) GetQuery() interface{}  { return &m.Query }
func (m *ApiAdminStaffListByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiAdminStaffListByOrganizationID) GetAsk() interface{}    { return nil }
func (m *ApiAdminStaffListByOrganizationID) GetAck() interface{}    { return m.Ack }
func MakeApiAdminStaffListByOrganizationID() ApiAdminStaffListByOrganizationID {
	return ApiAdminStaffListByOrganizationID{
		Ack: NewAckAdminStaffListByOrganizationID(),
	}
}

type AskAdminStaffAddByOrganizationID struct {
	Name   string `binding:"required,lte=64" db:"Name"`
	Mobile string `binding:"required,numeric" db:"Mobile"`
	RoleId uint32 `binding:"required,gt=0" db:"RoleId"`
}

func NewAskAdminStaffAddByOrganizationID() *AskAdminStaffAddByOrganizationID {
	return &AskAdminStaffAddByOrganizationID{}
}

type MetaApiAdminStaffAddByOrganizationID struct {
}

var META_ADMIN_STAFF_ADD_BY_ORGANIZATION_ID = &MetaApiAdminStaffAddByOrganizationID{}

func (m *MetaApiAdminStaffAddByOrganizationID) GetMethod() string { return "POST" }
func (m *MetaApiAdminStaffAddByOrganizationID) GetURL() string {
	return "/agency/admin/staff/add/:organization_id"
}
func (m *MetaApiAdminStaffAddByOrganizationID) GetName() string {
	return "AdminStaffAddByOrganizationID"
}
func (m *MetaApiAdminStaffAddByOrganizationID) GetType() string { return "json" }

// 添加团购组织成员
type ApiAdminStaffAddByOrganizationID struct {
	MetaApiAdminStaffAddByOrganizationID
	Ask    *AskAdminStaffAddByOrganizationID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiAdminStaffAddByOrganizationID) GetQuery() interface{}  { return nil }
func (m *ApiAdminStaffAddByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiAdminStaffAddByOrganizationID) GetAsk() interface{}    { return m.Ask }
func (m *ApiAdminStaffAddByOrganizationID) GetAck() interface{}    { return nil }
func MakeApiAdminStaffAddByOrganizationID() ApiAdminStaffAddByOrganizationID {
	return ApiAdminStaffAddByOrganizationID{
		Ask: NewAskAdminStaffAddByOrganizationID(),
	}
}

type AskAdminStaffEditByOrganizationIDByUserID struct {
	Name   string `binding:"required,lte=64" db:"Name"`
	Mobile string `binding:"required,numeric" db:"Mobile"`
	RoleId uint32 `binding:"required,gt=0" db:"RoleId"`
}

func NewAskAdminStaffEditByOrganizationIDByUserID() *AskAdminStaffEditByOrganizationIDByUserID {
	return &AskAdminStaffEditByOrganizationIDByUserID{}
}

type MetaApiAdminStaffEditByOrganizationIDByUserID struct {
}

var META_ADMIN_STAFF_EDIT_BY_ORGANIZATION_ID_BY_USER_ID = &MetaApiAdminStaffEditByOrganizationIDByUserID{}

func (m *MetaApiAdminStaffEditByOrganizationIDByUserID) GetMethod() string { return "POST" }
func (m *MetaApiAdminStaffEditByOrganizationIDByUserID) GetURL() string {
	return "/agency/admin/staff/edit/:organization_id/:user_id"
}
func (m *MetaApiAdminStaffEditByOrganizationIDByUserID) GetName() string {
	return "AdminStaffEditByOrganizationIDByUserID"
}
func (m *MetaApiAdminStaffEditByOrganizationIDByUserID) GetType() string { return "json" }

// 编辑团购组织成员
type ApiAdminStaffEditByOrganizationIDByUserID struct {
	MetaApiAdminStaffEditByOrganizationIDByUserID
	Ask    *AskAdminStaffEditByOrganizationIDByUserID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		UserID         string `form:"user_id" binding:"required" db:"UserID"`
	}
}

func (m *ApiAdminStaffEditByOrganizationIDByUserID) GetQuery() interface{}  { return nil }
func (m *ApiAdminStaffEditByOrganizationIDByUserID) GetParams() interface{} { return &m.Params }
func (m *ApiAdminStaffEditByOrganizationIDByUserID) GetAsk() interface{}    { return m.Ask }
func (m *ApiAdminStaffEditByOrganizationIDByUserID) GetAck() interface{}    { return nil }
func MakeApiAdminStaffEditByOrganizationIDByUserID() ApiAdminStaffEditByOrganizationIDByUserID {
	return ApiAdminStaffEditByOrganizationIDByUserID{
		Ask: NewAskAdminStaffEditByOrganizationIDByUserID(),
	}
}

type AskAdminStaffDisableByOrganizationIDByUserID struct {
	IsDisable bool `db:"IsDisable"`
}

func NewAskAdminStaffDisableByOrganizationIDByUserID() *AskAdminStaffDisableByOrganizationIDByUserID {
	return &AskAdminStaffDisableByOrganizationIDByUserID{}
}

type MetaApiAdminStaffDisableByOrganizationIDByUserID struct {
}

var META_ADMIN_STAFF_DISABLE_BY_ORGANIZATION_ID_BY_USER_ID = &MetaApiAdminStaffDisableByOrganizationIDByUserID{}

func (m *MetaApiAdminStaffDisableByOrganizationIDByUserID) GetMethod() string { return "POST" }
func (m *MetaApiAdminStaffDisableByOrganizationIDByUserID) GetURL() string {
	return "/agency/admin/staff/disable/:organization_id/:user_id"
}
func (m *MetaApiAdminStaffDisableByOrganizationIDByUserID) GetName() string {
	return "AdminStaffDisableByOrganizationIDByUserID"
}
func (m *MetaApiAdminStaffDisableByOrganizationIDByUserID) GetType() string { return "json" }

// 禁用团购组织成员
type ApiAdminStaffDisableByOrganizationIDByUserID struct {
	MetaApiAdminStaffDisableByOrganizationIDByUserID
	Ask    *AskAdminStaffDisableByOrganizationIDByUserID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		UserID         string `form:"user_id" binding:"required" db:"UserID"`
	}
}

func (m *ApiAdminStaffDisableByOrganizationIDByUserID) GetQuery() interface{}  { return nil }
func (m *ApiAdminStaffDisableByOrganizationIDByUserID) GetParams() interface{} { return &m.Params }
func (m *ApiAdminStaffDisableByOrganizationIDByUserID) GetAsk() interface{}    { return m.Ask }
func (m *ApiAdminStaffDisableByOrganizationIDByUserID) GetAck() interface{}    { return nil }
func MakeApiAdminStaffDisableByOrganizationIDByUserID() ApiAdminStaffDisableByOrganizationIDByUserID {
	return ApiAdminStaffDisableByOrganizationIDByUserID{
		Ask: NewAskAdminStaffDisableByOrganizationIDByUserID(),
	}
}

type AckAdminAuthorizationList struct {
	Modules []*StaffRoleAuthorizationGroup `db:"Modules"`
}

func NewAckAdminAuthorizationList() *AckAdminAuthorizationList {
	return &AckAdminAuthorizationList{
		Modules: make([]*StaffRoleAuthorizationGroup, 0),
	}
}

type MetaApiAdminAuthorizationList struct {
}

var META_ADMIN_AUTHORIZATION_LIST = &MetaApiAdminAuthorizationList{}

func (m *MetaApiAdminAuthorizationList) GetMethod() string { return "GET" }
func (m *MetaApiAdminAuthorizationList) GetURL() string    { return "/agency/admin/authorization/list" }
func (m *MetaApiAdminAuthorizationList) GetName() string   { return "AdminAuthorizationList" }
func (m *MetaApiAdminAuthorizationList) GetType() string   { return "json" }

// 权限列表
type ApiAdminAuthorizationList struct {
	MetaApiAdminAuthorizationList
	Ack *AckAdminAuthorizationList
}

func (m *ApiAdminAuthorizationList) GetQuery() interface{}  { return nil }
func (m *ApiAdminAuthorizationList) GetParams() interface{} { return nil }
func (m *ApiAdminAuthorizationList) GetAsk() interface{}    { return nil }
func (m *ApiAdminAuthorizationList) GetAck() interface{}    { return m.Ack }
func MakeApiAdminAuthorizationList() ApiAdminAuthorizationList {
	return ApiAdminAuthorizationList{
		Ack: NewAckAdminAuthorizationList(),
	}
}

type AckAdminRoleGetByOrganizationIDByRoleID struct {
	StaffRole *StaffRole                     `db:"StaffRole"`
	Modules   []*StaffRoleAuthorizationGroup `db:"Modules"`
}

func NewAckAdminRoleGetByOrganizationIDByRoleID() *AckAdminRoleGetByOrganizationIDByRoleID {
	return &AckAdminRoleGetByOrganizationIDByRoleID{
		StaffRole: NewStaffRole(),
		Modules:   make([]*StaffRoleAuthorizationGroup, 0),
	}
}

type MetaApiAdminRoleGetByOrganizationIDByRoleID struct {
}

var META_ADMIN_ROLE_GET_BY_ORGANIZATION_ID_BY_ROLE_ID = &MetaApiAdminRoleGetByOrganizationIDByRoleID{}

func (m *MetaApiAdminRoleGetByOrganizationIDByRoleID) GetMethod() string { return "GET" }
func (m *MetaApiAdminRoleGetByOrganizationIDByRoleID) GetURL() string {
	return "/agency/admin/role/get/:organization_id/:role_id"
}
func (m *MetaApiAdminRoleGetByOrganizationIDByRoleID) GetName() string {
	return "AdminRoleGetByOrganizationIDByRoleID"
}
func (m *MetaApiAdminRoleGetByOrganizationIDByRoleID) GetType() string { return "json" }

// 获取角色
type ApiAdminRoleGetByOrganizationIDByRoleID struct {
	MetaApiAdminRoleGetByOrganizationIDByRoleID
	Ack    *AckAdminRoleGetByOrganizationIDByRoleID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		RoleID         uint32 `form:"role_id" binding:"required,gt=0" db:"RoleID"`
	}
}

func (m *ApiAdminRoleGetByOrganizationIDByRoleID) GetQuery() interface{}  { return nil }
func (m *ApiAdminRoleGetByOrganizationIDByRoleID) GetParams() interface{} { return &m.Params }
func (m *ApiAdminRoleGetByOrganizationIDByRoleID) GetAsk() interface{}    { return nil }
func (m *ApiAdminRoleGetByOrganizationIDByRoleID) GetAck() interface{}    { return m.Ack }
func MakeApiAdminRoleGetByOrganizationIDByRoleID() ApiAdminRoleGetByOrganizationIDByRoleID {
	return ApiAdminRoleGetByOrganizationIDByRoleID{
		Ack: NewAckAdminRoleGetByOrganizationIDByRoleID(),
	}
}

type AskAdminRoleAddByOrganizationID struct {
	RoleName         string   `binding:"required,lte=64" db:"RoleName"`
	AuthorizationIds []uint32 `binding:"required,gt=0" db:"AuthorizationIds"`
}

func NewAskAdminRoleAddByOrganizationID() *AskAdminRoleAddByOrganizationID {
	return &AskAdminRoleAddByOrganizationID{
		AuthorizationIds: make([]uint32, 0),
	}
}

type MetaApiAdminRoleAddByOrganizationID struct {
}

var META_ADMIN_ROLE_ADD_BY_ORGANIZATION_ID = &MetaApiAdminRoleAddByOrganizationID{}

func (m *MetaApiAdminRoleAddByOrganizationID) GetMethod() string { return "POST" }
func (m *MetaApiAdminRoleAddByOrganizationID) GetURL() string {
	return "/agency/admin/role/add/:organization_id"
}
func (m *MetaApiAdminRoleAddByOrganizationID) GetName() string { return "AdminRoleAddByOrganizationID" }
func (m *MetaApiAdminRoleAddByOrganizationID) GetType() string { return "json" }

// 添加角色
type ApiAdminRoleAddByOrganizationID struct {
	MetaApiAdminRoleAddByOrganizationID
	Ask    *AskAdminRoleAddByOrganizationID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiAdminRoleAddByOrganizationID) GetQuery() interface{}  { return nil }
func (m *ApiAdminRoleAddByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiAdminRoleAddByOrganizationID) GetAsk() interface{}    { return m.Ask }
func (m *ApiAdminRoleAddByOrganizationID) GetAck() interface{}    { return nil }
func MakeApiAdminRoleAddByOrganizationID() ApiAdminRoleAddByOrganizationID {
	return ApiAdminRoleAddByOrganizationID{
		Ask: NewAskAdminRoleAddByOrganizationID(),
	}
}

type AskAdminRoleEditByOrganizationIDByRoleID struct {
	RoleName         string   `binding:"required,lte=64" db:"RoleName"`
	AuthorizationIds []uint32 `binding:"required,gt=0" db:"AuthorizationIds"`
}

func NewAskAdminRoleEditByOrganizationIDByRoleID() *AskAdminRoleEditByOrganizationIDByRoleID {
	return &AskAdminRoleEditByOrganizationIDByRoleID{
		AuthorizationIds: make([]uint32, 0),
	}
}

type MetaApiAdminRoleEditByOrganizationIDByRoleID struct {
}

var META_ADMIN_ROLE_EDIT_BY_ORGANIZATION_ID_BY_ROLE_ID = &MetaApiAdminRoleEditByOrganizationIDByRoleID{}

func (m *MetaApiAdminRoleEditByOrganizationIDByRoleID) GetMethod() string { return "POST" }
func (m *MetaApiAdminRoleEditByOrganizationIDByRoleID) GetURL() string {
	return "/agency/admin/role/edit/:organization_id/:role_id"
}
func (m *MetaApiAdminRoleEditByOrganizationIDByRoleID) GetName() string {
	return "AdminRoleEditByOrganizationIDByRoleID"
}
func (m *MetaApiAdminRoleEditByOrganizationIDByRoleID) GetType() string { return "json" }

// 编辑角色
type ApiAdminRoleEditByOrganizationIDByRoleID struct {
	MetaApiAdminRoleEditByOrganizationIDByRoleID
	Ask    *AskAdminRoleEditByOrganizationIDByRoleID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		RoleID         uint32 `form:"role_id" binding:"required,gt=0" db:"RoleID"`
	}
}

func (m *ApiAdminRoleEditByOrganizationIDByRoleID) GetQuery() interface{}  { return nil }
func (m *ApiAdminRoleEditByOrganizationIDByRoleID) GetParams() interface{} { return &m.Params }
func (m *ApiAdminRoleEditByOrganizationIDByRoleID) GetAsk() interface{}    { return m.Ask }
func (m *ApiAdminRoleEditByOrganizationIDByRoleID) GetAck() interface{}    { return nil }
func MakeApiAdminRoleEditByOrganizationIDByRoleID() ApiAdminRoleEditByOrganizationIDByRoleID {
	return ApiAdminRoleEditByOrganizationIDByRoleID{
		Ask: NewAskAdminRoleEditByOrganizationIDByRoleID(),
	}
}

type AskAdminRoleDisableByOrganizationIDByRoleID struct {
	IsDisable bool `db:"IsDisable"`
}

func NewAskAdminRoleDisableByOrganizationIDByRoleID() *AskAdminRoleDisableByOrganizationIDByRoleID {
	return &AskAdminRoleDisableByOrganizationIDByRoleID{}
}

type MetaApiAdminRoleDisableByOrganizationIDByRoleID struct {
}

var META_ADMIN_ROLE_DISABLE_BY_ORGANIZATION_ID_BY_ROLE_ID = &MetaApiAdminRoleDisableByOrganizationIDByRoleID{}

func (m *MetaApiAdminRoleDisableByOrganizationIDByRoleID) GetMethod() string { return "POST" }
func (m *MetaApiAdminRoleDisableByOrganizationIDByRoleID) GetURL() string {
	return "/agency/admin/role/disable/:organization_id/:role_id"
}
func (m *MetaApiAdminRoleDisableByOrganizationIDByRoleID) GetName() string {
	return "AdminRoleDisableByOrganizationIDByRoleID"
}
func (m *MetaApiAdminRoleDisableByOrganizationIDByRoleID) GetType() string { return "json" }

// 禁用角色
type ApiAdminRoleDisableByOrganizationIDByRoleID struct {
	MetaApiAdminRoleDisableByOrganizationIDByRoleID
	Ask    *AskAdminRoleDisableByOrganizationIDByRoleID
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
		RoleID         uint32 `form:"role_id" binding:"required,gt=0" db:"RoleID"`
	}
}

func (m *ApiAdminRoleDisableByOrganizationIDByRoleID) GetQuery() interface{}  { return nil }
func (m *ApiAdminRoleDisableByOrganizationIDByRoleID) GetParams() interface{} { return &m.Params }
func (m *ApiAdminRoleDisableByOrganizationIDByRoleID) GetAsk() interface{}    { return m.Ask }
func (m *ApiAdminRoleDisableByOrganizationIDByRoleID) GetAck() interface{}    { return nil }
func MakeApiAdminRoleDisableByOrganizationIDByRoleID() ApiAdminRoleDisableByOrganizationIDByRoleID {
	return ApiAdminRoleDisableByOrganizationIDByRoleID{
		Ask: NewAskAdminRoleDisableByOrganizationIDByRoleID(),
	}
}
