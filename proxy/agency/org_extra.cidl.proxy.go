package agency

type AckOrgStaffCheckNewMobileByMobile struct {
	CanBeBound  bool   `db:"CanBeBound"`
	IsUserExist bool   `db:"IsUserExist"`
	UserName    string `db:"UserName"`
}

func NewAckOrgStaffCheckNewMobileByMobile() *AckOrgStaffCheckNewMobileByMobile {
	return &AckOrgStaffCheckNewMobileByMobile{}
}

// 添加或者编辑组织成员时，检查新手机是否已经绑定其他用户
func (m *Proxy) OrgStaffCheckNewMobileByMobile(Mobile string,
) (*AckOrgStaffCheckNewMobileByMobile, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckOrgStaffCheckNewMobileByMobile
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/agency/org/staff/check_new_mobile/:mobile",
		nil,
		ack,
		map[string]interface{}{
			"mobile": Mobile,
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
