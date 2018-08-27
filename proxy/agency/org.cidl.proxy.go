package agency

type AckOrgOrganizationList struct {
	Count uint32          `db:"Count"`
	List  []*Organization `db:"List"`
}

func NewAckOrgOrganizationList() *AckOrgOrganizationList {
	return &AckOrgOrganizationList{
		List: make([]*Organization, 0),
	}
}

// 获取团购组织列表
func (m *Proxy) OrgOrganizationList() (*AckOrgOrganizationList, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckOrgOrganizationList
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/agency/org/organization/list",
		nil,
		ack,
		nil,
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
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

// 获取团购组织详细信息
func (m *Proxy) OrgOrganizationInfoByID(OrganizationID uint32,
) (*AckOrgOrganizationInfoByID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckOrgOrganizationInfoByID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/agency/org/organization/info/:organization_id",
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

// 获取可用组织的纪录
func (m *Proxy) OrgOrganizationEnableInfoByOrganizationID(OrganizationID uint32,
) (*Organization, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *Organization
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/agency/org/organization/enable_info/:organization_id",
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

type AckOrgOrganizationGetOrganizationID struct {
	OrganizationId uint32 `db:"OrganizationId"`
}

func NewAckOrgOrganizationGetOrganizationID() *AckOrgOrganizationGetOrganizationID {
	return &AckOrgOrganizationGetOrganizationID{}
}

// 通过uid获取团购组织id
func (m *Proxy) OrgOrganizationGetOrganizationID() (*AckOrgOrganizationGetOrganizationID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckOrgOrganizationGetOrganizationID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/agency/org/organization/get_organization_id",
		nil,
		ack,
		nil,
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
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

// 团购成员角色权限列表
func (m *Proxy) OrgRoleListByOrganizationID(OrganizationID uint32,
) (*AckOrgRoleListByOrganizationID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckOrgRoleListByOrganizationID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/agency/org/role/list/:organization_id",
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

type AckOrgStaffListByOrganizationID struct {
	Count uint32   `db:"Count"`
	List  []*Staff `db:"List"`
}

func NewAckOrgStaffListByOrganizationID() *AckOrgStaffListByOrganizationID {
	return &AckOrgStaffListByOrganizationID{
		List: make([]*Staff, 0),
	}
}

// 获取团购成员列表
func (m *Proxy) OrgStaffListByOrganizationID(OrganizationID uint32,
) (*AckOrgStaffListByOrganizationID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckOrgStaffListByOrganizationID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/agency/org/staff/list/:organization_id",
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

type AskOrgStaffAddByOrganizationID struct {
	Name   string `binding:"required,lte=64" db:"Name"`
	Mobile string `binding:"required,numeric" db:"Mobile"`
	RoleId uint32 `binding:"required,gt=0" db:"RoleId"`
}

func NewAskOrgStaffAddByOrganizationID() *AskOrgStaffAddByOrganizationID {
	return &AskOrgStaffAddByOrganizationID{}
}

// 添加团购组织成员
func (m *Proxy) OrgStaffAddByOrganizationID(OrganizationID uint32,
	ask *AskOrgStaffAddByOrganizationID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/agency/org/staff/add/:organization_id",
		ask,
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

type AskOrgStaffEditByOrganizationIDByUserID struct {
	Name   string `binding:"required,lte=64" db:"Name"`
	Mobile string `binding:"required,numeric" db:"Mobile"`
	RoleId uint32 `binding:"required,gt=0" db:"RoleId"`
}

func NewAskOrgStaffEditByOrganizationIDByUserID() *AskOrgStaffEditByOrganizationIDByUserID {
	return &AskOrgStaffEditByOrganizationIDByUserID{}
}

// 编辑团购组织成员
func (m *Proxy) OrgStaffEditByOrganizationIDByUserID(OrganizationID uint32,
	UserID string,
	ask *AskOrgStaffEditByOrganizationIDByUserID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/agency/org/staff/edit/:organization_id/:user_id",
		ask,
		ack,
		map[string]interface{}{
			"organization_id": OrganizationID,
			"user_id":         UserID,
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

type AskOrgStaffDisableByOrganizationIDByUserID struct {
	IsDisable bool `db:"IsDisable"`
}

func NewAskOrgStaffDisableByOrganizationIDByUserID() *AskOrgStaffDisableByOrganizationIDByUserID {
	return &AskOrgStaffDisableByOrganizationIDByUserID{}
}

// 禁用团购组织成员
func (m *Proxy) OrgStaffDisableByOrganizationIDByUserID(OrganizationID uint32,
	UserID string,
	ask *AskOrgStaffDisableByOrganizationIDByUserID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/agency/org/staff/disable/:organization_id/:user_id",
		ask,
		ack,
		map[string]interface{}{
			"organization_id": OrganizationID,
			"user_id":         UserID,
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

type AckOrgAuthorizationList struct {
	Modules []*StaffRoleAuthorizationGroup `db:"Modules"`
}

func NewAckOrgAuthorizationList() *AckOrgAuthorizationList {
	return &AckOrgAuthorizationList{
		Modules: make([]*StaffRoleAuthorizationGroup, 0),
	}
}

// 权限列表
func (m *Proxy) OrgAuthorizationList() (*AckOrgAuthorizationList, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckOrgAuthorizationList
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/agency/org/authorization/list",
		nil,
		ack,
		nil,
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
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

// 获取角色
func (m *Proxy) OrgRoleGetByOrganizationIDByRoleID(OrganizationID uint32,
	RoleID uint32,
) (*AckOrgRoleGetByOrganizationIDByRoleID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckOrgRoleGetByOrganizationIDByRoleID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/agency/org/role/get/:organization_id/:role_id",
		nil,
		ack,
		map[string]interface{}{
			"organization_id": OrganizationID,
			"role_id":         RoleID,
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

type AskOrgRoleAddByOrganizationID struct {
	RoleName         string   `binding:"required,lte=64" db:"RoleName"`
	AuthorizationIds []uint32 `binding:"required,gt=0" db:"AuthorizationIds"`
}

func NewAskOrgRoleAddByOrganizationID() *AskOrgRoleAddByOrganizationID {
	return &AskOrgRoleAddByOrganizationID{
		AuthorizationIds: make([]uint32, 0),
	}
}

// 添加角色
func (m *Proxy) OrgRoleAddByOrganizationID(OrganizationID uint32,
	ask *AskOrgRoleAddByOrganizationID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/agency/org/role/add/:organization_id",
		ask,
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

type AskOrgRoleEditByOrganizationIDByRoleID struct {
	RoleName         string   `binding:"required,lte=64" db:"RoleName"`
	AuthorizationIds []uint32 `binding:"required,gt=0" db:"AuthorizationIds"`
}

func NewAskOrgRoleEditByOrganizationIDByRoleID() *AskOrgRoleEditByOrganizationIDByRoleID {
	return &AskOrgRoleEditByOrganizationIDByRoleID{
		AuthorizationIds: make([]uint32, 0),
	}
}

// 编辑角色
func (m *Proxy) OrgRoleEditByOrganizationIDByRoleID(OrganizationID uint32,
	RoleID uint32,
	ask *AskOrgRoleEditByOrganizationIDByRoleID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/agency/org/role/edit/:organization_id/:role_id",
		ask,
		ack,
		map[string]interface{}{
			"organization_id": OrganizationID,
			"role_id":         RoleID,
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

type AskOrgRoleDisableByOrganizationIDByRoleID struct {
	IsDisable bool `db:"IsDisable"`
}

func NewAskOrgRoleDisableByOrganizationIDByRoleID() *AskOrgRoleDisableByOrganizationIDByRoleID {
	return &AskOrgRoleDisableByOrganizationIDByRoleID{}
}

// 禁用角色
func (m *Proxy) OrgRoleDisableByOrganizationIDByRoleID(OrganizationID uint32,
	RoleID uint32,
	ask *AskOrgRoleDisableByOrganizationIDByRoleID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/agency/org/role/disable/:organization_id/:role_id",
		ask,
		ack,
		map[string]interface{}{
			"organization_id": OrganizationID,
			"role_id":         RoleID,
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
