package cidl

type AckOrgStaffCheckNewMobileByMobile struct {
	CanBeBound  bool   `db:"CanBeBound"`
	IsUserExist bool   `db:"IsUserExist"`
	UserName    string `db:"UserName"`
}

func NewAckOrgStaffCheckNewMobileByMobile() *AckOrgStaffCheckNewMobileByMobile {
	return &AckOrgStaffCheckNewMobileByMobile{}
}

type MetaApiOrgStaffCheckNewMobileByMobile struct {
}

var META_ORG_STAFF_CHECK_NEW_MOBILE_BY_MOBILE = &MetaApiOrgStaffCheckNewMobileByMobile{}

func (m *MetaApiOrgStaffCheckNewMobileByMobile) GetMethod() string { return "GET" }
func (m *MetaApiOrgStaffCheckNewMobileByMobile) GetURL() string {
	return "/agency/org/staff/check_new_mobile/:mobile"
}
func (m *MetaApiOrgStaffCheckNewMobileByMobile) GetName() string {
	return "OrgStaffCheckNewMobileByMobile"
}
func (m *MetaApiOrgStaffCheckNewMobileByMobile) GetType() string { return "json" }

// 添加或者编辑组织成员时，检查新手机是否已经绑定其他用户
type ApiOrgStaffCheckNewMobileByMobile struct {
	MetaApiOrgStaffCheckNewMobileByMobile
	Ack    *AckOrgStaffCheckNewMobileByMobile
	Params struct {
		Mobile string `form:"mobile" binding:"required,numeric" db:"Mobile"`
	}
}

func (m *ApiOrgStaffCheckNewMobileByMobile) GetQuery() interface{}  { return nil }
func (m *ApiOrgStaffCheckNewMobileByMobile) GetParams() interface{} { return &m.Params }
func (m *ApiOrgStaffCheckNewMobileByMobile) GetAsk() interface{}    { return nil }
func (m *ApiOrgStaffCheckNewMobileByMobile) GetAck() interface{}    { return m.Ack }
func MakeApiOrgStaffCheckNewMobileByMobile() ApiOrgStaffCheckNewMobileByMobile {
	return ApiOrgStaffCheckNewMobileByMobile{
		Ack: NewAckOrgStaffCheckNewMobileByMobile(),
	}
}
