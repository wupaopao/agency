package agency

// 获取团购组织
func (m *Proxy) InnerAgencyOrganizationInfoByOrganizationID(OrganizationID uint32,
) (*Organization, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *Organization
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/inner/agency/organization/info/:organization_id",
		nil,
		ack,
		map[string]interface{}{
			"organization_id": OrganizationID,
		},
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}

// 获取可用团购组织
func (m *Proxy) InnerAgencyOrganizationEnableInfoByOrganizationID(OrganizationID uint32,
) (*Organization, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *Organization
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/inner/agency/organization/enable_info/:organization_id",
		nil,
		ack,
		map[string]interface{}{
			"organization_id": OrganizationID,
		},
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}

// 通过uid获取团购组织
func (m *Proxy) InnerAgencyOrganizationInfoByUserIDByUserID(UserID string,
) (*Organization, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *Organization
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/inner/agency/organization/info_by_user_id/:user_id",
		nil,
		ack,
		map[string]interface{}{
			"user_id": UserID,
		},
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
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

// 获取用户的团购组织成员信息
func (m *Proxy) InnerAgencyStaffInfoByUserID(UserID string,
) (*AckInnerAgencyStaffInfoByUserID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckInnerAgencyStaffInfoByUserID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/inner/agency/staff/info/:user_id",
		nil,
		ack,
		map[string]interface{}{
			"user_id": UserID,
		},
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}

type AckInnerAgencyStaffIsDisableByUserID struct {
	IsDisable bool `db:"IsDisable"`
}

func NewAckInnerAgencyStaffIsDisableByUserID() *AckInnerAgencyStaffIsDisableByUserID {
	return &AckInnerAgencyStaffIsDisableByUserID{}
}

// 是否被禁用
func (m *Proxy) InnerAgencyStaffIsDisableByUserID(UserID string,
) (*AckInnerAgencyStaffIsDisableByUserID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckInnerAgencyStaffIsDisableByUserID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/inner/agency/staff/is_disable/:user_id",
		nil,
		ack,
		map[string]interface{}{
			"user_id": UserID,
		},
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}
