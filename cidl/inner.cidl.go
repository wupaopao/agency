package cidl

type MetaApiInnerAgencyOrganizationInfoByOrganizationID struct {
}

var META_INNER_AGENCY_ORGANIZATION_INFO_BY_ORGANIZATION_ID = &MetaApiInnerAgencyOrganizationInfoByOrganizationID{}

func (m *MetaApiInnerAgencyOrganizationInfoByOrganizationID) GetMethod() string { return "GET" }
func (m *MetaApiInnerAgencyOrganizationInfoByOrganizationID) GetURL() string {
	return "/inner/agency/organization/info/:organization_id"
}
func (m *MetaApiInnerAgencyOrganizationInfoByOrganizationID) GetName() string {
	return "InnerAgencyOrganizationInfoByOrganizationID"
}
func (m *MetaApiInnerAgencyOrganizationInfoByOrganizationID) GetType() string { return "json" }

// 获取团购组织
type ApiInnerAgencyOrganizationInfoByOrganizationID struct {
	MetaApiInnerAgencyOrganizationInfoByOrganizationID
	Ack    *Organization
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiInnerAgencyOrganizationInfoByOrganizationID) GetQuery() interface{}  { return nil }
func (m *ApiInnerAgencyOrganizationInfoByOrganizationID) GetParams() interface{} { return &m.Params }
func (m *ApiInnerAgencyOrganizationInfoByOrganizationID) GetAsk() interface{}    { return nil }
func (m *ApiInnerAgencyOrganizationInfoByOrganizationID) GetAck() interface{}    { return m.Ack }
func MakeApiInnerAgencyOrganizationInfoByOrganizationID() ApiInnerAgencyOrganizationInfoByOrganizationID {
	return ApiInnerAgencyOrganizationInfoByOrganizationID{
		Ack: NewOrganization(),
	}
}

type MetaApiInnerAgencyOrganizationEnableInfoByOrganizationID struct {
}

var META_INNER_AGENCY_ORGANIZATION_ENABLE_INFO_BY_ORGANIZATION_ID = &MetaApiInnerAgencyOrganizationEnableInfoByOrganizationID{}

func (m *MetaApiInnerAgencyOrganizationEnableInfoByOrganizationID) GetMethod() string { return "GET" }
func (m *MetaApiInnerAgencyOrganizationEnableInfoByOrganizationID) GetURL() string {
	return "/inner/agency/organization/enable_info/:organization_id"
}
func (m *MetaApiInnerAgencyOrganizationEnableInfoByOrganizationID) GetName() string {
	return "InnerAgencyOrganizationEnableInfoByOrganizationID"
}
func (m *MetaApiInnerAgencyOrganizationEnableInfoByOrganizationID) GetType() string { return "json" }

// 获取可用团购组织
type ApiInnerAgencyOrganizationEnableInfoByOrganizationID struct {
	MetaApiInnerAgencyOrganizationEnableInfoByOrganizationID
	Ack    *Organization
	Params struct {
		OrganizationID uint32 `form:"organization_id" binding:"required,gt=0" db:"OrganizationID"`
	}
}

func (m *ApiInnerAgencyOrganizationEnableInfoByOrganizationID) GetQuery() interface{} { return nil }
func (m *ApiInnerAgencyOrganizationEnableInfoByOrganizationID) GetParams() interface{} {
	return &m.Params
}
func (m *ApiInnerAgencyOrganizationEnableInfoByOrganizationID) GetAsk() interface{} { return nil }
func (m *ApiInnerAgencyOrganizationEnableInfoByOrganizationID) GetAck() interface{} { return m.Ack }
func MakeApiInnerAgencyOrganizationEnableInfoByOrganizationID() ApiInnerAgencyOrganizationEnableInfoByOrganizationID {
	return ApiInnerAgencyOrganizationEnableInfoByOrganizationID{
		Ack: NewOrganization(),
	}
}

type MetaApiInnerAgencyOrganizationInfoByUserIDByUserID struct {
}

var META_INNER_AGENCY_ORGANIZATION_INFO_BY_USER_ID_BY_USER_ID = &MetaApiInnerAgencyOrganizationInfoByUserIDByUserID{}

func (m *MetaApiInnerAgencyOrganizationInfoByUserIDByUserID) GetMethod() string { return "GET" }
func (m *MetaApiInnerAgencyOrganizationInfoByUserIDByUserID) GetURL() string {
	return "/inner/agency/organization/info_by_user_id/:user_id"
}
func (m *MetaApiInnerAgencyOrganizationInfoByUserIDByUserID) GetName() string {
	return "InnerAgencyOrganizationInfoByUserIDByUserID"
}
func (m *MetaApiInnerAgencyOrganizationInfoByUserIDByUserID) GetType() string { return "json" }

// 通过uid获取团购组织
type ApiInnerAgencyOrganizationInfoByUserIDByUserID struct {
	MetaApiInnerAgencyOrganizationInfoByUserIDByUserID
	Ack    *Organization
	Params struct {
		UserID string `form:"user_id" db:"UserID"`
	}
}

func (m *ApiInnerAgencyOrganizationInfoByUserIDByUserID) GetQuery() interface{}  { return nil }
func (m *ApiInnerAgencyOrganizationInfoByUserIDByUserID) GetParams() interface{} { return &m.Params }
func (m *ApiInnerAgencyOrganizationInfoByUserIDByUserID) GetAsk() interface{}    { return nil }
func (m *ApiInnerAgencyOrganizationInfoByUserIDByUserID) GetAck() interface{}    { return m.Ack }
func MakeApiInnerAgencyOrganizationInfoByUserIDByUserID() ApiInnerAgencyOrganizationInfoByUserIDByUserID {
	return ApiInnerAgencyOrganizationInfoByUserIDByUserID{
		Ack: NewOrganization(),
	}
}

type AckInnerAgencyStaffInfoByUserID struct {
	Staff        *Staff        `db:"Staff"`
	StaffRole    *StaffRole    `db:"StaffRole"`
	Organization *Organization `db:"Organization"`
}

func NewAckInnerAgencyStaffInfoByUserID() *AckInnerAgencyStaffInfoByUserID {
	return &AckInnerAgencyStaffInfoByUserID{
		Staff:        NewStaff(),
		StaffRole:    NewStaffRole(),
		Organization: NewOrganization(),
	}
}

type MetaApiInnerAgencyStaffInfoByUserID struct {
}

var META_INNER_AGENCY_STAFF_INFO_BY_USER_ID = &MetaApiInnerAgencyStaffInfoByUserID{}

func (m *MetaApiInnerAgencyStaffInfoByUserID) GetMethod() string { return "GET" }
func (m *MetaApiInnerAgencyStaffInfoByUserID) GetURL() string {
	return "/inner/agency/staff/info/:user_id"
}
func (m *MetaApiInnerAgencyStaffInfoByUserID) GetName() string { return "InnerAgencyStaffInfoByUserID" }
func (m *MetaApiInnerAgencyStaffInfoByUserID) GetType() string { return "json" }

// 获取用户的团购组织成员信息
type ApiInnerAgencyStaffInfoByUserID struct {
	MetaApiInnerAgencyStaffInfoByUserID
	Ack    *AckInnerAgencyStaffInfoByUserID
	Params struct {
		UserID string `form:"user_id" db:"UserID"`
	}
}

func (m *ApiInnerAgencyStaffInfoByUserID) GetQuery() interface{}  { return nil }
func (m *ApiInnerAgencyStaffInfoByUserID) GetParams() interface{} { return &m.Params }
func (m *ApiInnerAgencyStaffInfoByUserID) GetAsk() interface{}    { return nil }
func (m *ApiInnerAgencyStaffInfoByUserID) GetAck() interface{}    { return m.Ack }
func MakeApiInnerAgencyStaffInfoByUserID() ApiInnerAgencyStaffInfoByUserID {
	return ApiInnerAgencyStaffInfoByUserID{
		Ack: NewAckInnerAgencyStaffInfoByUserID(),
	}
}

type AckInnerAgencyStaffIsDisableByUserID struct {
	IsDisable bool `db:"IsDisable"`
}

func NewAckInnerAgencyStaffIsDisableByUserID() *AckInnerAgencyStaffIsDisableByUserID {
	return &AckInnerAgencyStaffIsDisableByUserID{}
}

type MetaApiInnerAgencyStaffIsDisableByUserID struct {
}

var META_INNER_AGENCY_STAFF_IS_DISABLE_BY_USER_ID = &MetaApiInnerAgencyStaffIsDisableByUserID{}

func (m *MetaApiInnerAgencyStaffIsDisableByUserID) GetMethod() string { return "GET" }
func (m *MetaApiInnerAgencyStaffIsDisableByUserID) GetURL() string {
	return "/inner/agency/staff/is_disable/:user_id"
}
func (m *MetaApiInnerAgencyStaffIsDisableByUserID) GetName() string {
	return "InnerAgencyStaffIsDisableByUserID"
}
func (m *MetaApiInnerAgencyStaffIsDisableByUserID) GetType() string { return "json" }

// 是否被禁用
type ApiInnerAgencyStaffIsDisableByUserID struct {
	MetaApiInnerAgencyStaffIsDisableByUserID
	Ack    *AckInnerAgencyStaffIsDisableByUserID
	Params struct {
		UserID string `form:"user_id" db:"UserID"`
	}
}

func (m *ApiInnerAgencyStaffIsDisableByUserID) GetQuery() interface{}  { return nil }
func (m *ApiInnerAgencyStaffIsDisableByUserID) GetParams() interface{} { return &m.Params }
func (m *ApiInnerAgencyStaffIsDisableByUserID) GetAsk() interface{}    { return nil }
func (m *ApiInnerAgencyStaffIsDisableByUserID) GetAck() interface{}    { return m.Ack }
func MakeApiInnerAgencyStaffIsDisableByUserID() ApiInnerAgencyStaffIsDisableByUserID {
	return ApiInnerAgencyStaffIsDisableByUserID{
		Ack: NewAckInnerAgencyStaffIsDisableByUserID(),
	}
}
