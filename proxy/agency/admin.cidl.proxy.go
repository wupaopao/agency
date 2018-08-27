package agency

type AckAdminOrganizationList struct {
	Count uint32          `db:"Count"`
	List  []*Organization `db:"List"`
}

func NewAckAdminOrganizationList() *AckAdminOrganizationList {
	return &AckAdminOrganizationList{
		List: make([]*Organization, 0),
	}
}

// 获取团购组织列表
func (m *Proxy) AdminOrganizationList() (*AckAdminOrganizationList, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckAdminOrganizationList
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/agency/admin/organization/list",
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

// 获取团购组织详细信息
func (m *Proxy) AdminOrganizationInfoByID(OrganizationID uint32,
) (*AckAdminOrganizationInfoByID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckAdminOrganizationInfoByID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/agency/admin/organization/info/:organization_id",
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
func (m *Proxy) AdminOrganizationAddPicToken(ask *AskAdminOrganizationAddPicToken,
) (*AckAdminOrganizationAddPicToken, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckAdminOrganizationAddPicToken
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/agency/admin/organization/add/pic_token",
		ask,
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

// 添加组织
func (m *Proxy) AdminOrganizationAdd(ask *AskAdminOrganizationAdd,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/agency/admin/organization/add",
		ask,
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

// 编辑组织
func (m *Proxy) AdminOrganizationEditByOrganizationID(OrganizationID uint32,
	ask *AskAdminOrganizationEditByOrganizationID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/agency/admin/organization/edit/:organization_id",
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

type AskAdminOrganizationDisableByOrganizationID struct {
	IsDisable bool `db:"IsDisable"`
}

func NewAskAdminOrganizationDisableByOrganizationID() *AskAdminOrganizationDisableByOrganizationID {
	return &AskAdminOrganizationDisableByOrganizationID{}
}

// 禁用组织
func (m *Proxy) AdminOrganizationDisableByOrganizationID(OrganizationID string,
	ask *AskAdminOrganizationDisableByOrganizationID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/agency/admin/organization/disable/:organization_id",
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

// 获取可用组织的纪录
func (m *Proxy) AdminOrganizationEnableInfoByOrganizationID(OrganizationID uint32,
) (*Organization, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *Organization
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/agency/admin/organization/enable_info/:organization_id",
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

type AckAdminOrganizationGetOrganizationID struct {
	OrganizationId uint32 `db:"OrganizationId"`
}

func NewAckAdminOrganizationGetOrganizationID() *AckAdminOrganizationGetOrganizationID {
	return &AckAdminOrganizationGetOrganizationID{}
}

// 通过uid获取团购组织id
func (m *Proxy) AdminOrganizationGetOrganizationID() (*AckAdminOrganizationGetOrganizationID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckAdminOrganizationGetOrganizationID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/agency/admin/organization/get_organization_id",
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

type AckAdminRoleListByOrganizationID struct {
	Count uint32       `db:"Count"`
	List  []*StaffRole `db:"List"`
}

func NewAckAdminRoleListByOrganizationID() *AckAdminRoleListByOrganizationID {
	return &AckAdminRoleListByOrganizationID{
		List: make([]*StaffRole, 0),
	}
}

// 团购成员角色权限列表
func (m *Proxy) AdminRoleListByOrganizationID(OrganizationID uint32,
) (*AckAdminRoleListByOrganizationID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckAdminRoleListByOrganizationID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/agency/admin/role/list/:organization_id",
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

type AckAdminStaffListByOrganizationID struct {
	Count uint32   `db:"Count"`
	List  []*Staff `db:"List"`
}

func NewAckAdminStaffListByOrganizationID() *AckAdminStaffListByOrganizationID {
	return &AckAdminStaffListByOrganizationID{
		List: make([]*Staff, 0),
	}
}

// 获取团购成员列表
func (m *Proxy) AdminStaffListByOrganizationID(OrganizationID uint32,
) (*AckAdminStaffListByOrganizationID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckAdminStaffListByOrganizationID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/agency/admin/staff/list/:organization_id",
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

type AskAdminStaffAddByOrganizationID struct {
	Name   string `binding:"required,lte=64" db:"Name"`
	Mobile string `binding:"required,numeric" db:"Mobile"`
	RoleId uint32 `binding:"required,gt=0" db:"RoleId"`
}

func NewAskAdminStaffAddByOrganizationID() *AskAdminStaffAddByOrganizationID {
	return &AskAdminStaffAddByOrganizationID{}
}

// 添加团购组织成员
func (m *Proxy) AdminStaffAddByOrganizationID(OrganizationID uint32,
	ask *AskAdminStaffAddByOrganizationID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/agency/admin/staff/add/:organization_id",
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

type AskAdminStaffEditByOrganizationIDByUserID struct {
	Name   string `binding:"required,lte=64" db:"Name"`
	Mobile string `binding:"required,numeric" db:"Mobile"`
	RoleId uint32 `binding:"required,gt=0" db:"RoleId"`
}

func NewAskAdminStaffEditByOrganizationIDByUserID() *AskAdminStaffEditByOrganizationIDByUserID {
	return &AskAdminStaffEditByOrganizationIDByUserID{}
}

// 编辑团购组织成员
func (m *Proxy) AdminStaffEditByOrganizationIDByUserID(OrganizationID uint32,
	UserID string,
	ask *AskAdminStaffEditByOrganizationIDByUserID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/agency/admin/staff/edit/:organization_id/:user_id",
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

type AskAdminStaffDisableByOrganizationIDByUserID struct {
	IsDisable bool `db:"IsDisable"`
}

func NewAskAdminStaffDisableByOrganizationIDByUserID() *AskAdminStaffDisableByOrganizationIDByUserID {
	return &AskAdminStaffDisableByOrganizationIDByUserID{}
}

// 禁用团购组织成员
func (m *Proxy) AdminStaffDisableByOrganizationIDByUserID(OrganizationID uint32,
	UserID string,
	ask *AskAdminStaffDisableByOrganizationIDByUserID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/agency/admin/staff/disable/:organization_id/:user_id",
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

type AckAdminAuthorizationList struct {
	Modules []*StaffRoleAuthorizationGroup `db:"Modules"`
}

func NewAckAdminAuthorizationList() *AckAdminAuthorizationList {
	return &AckAdminAuthorizationList{
		Modules: make([]*StaffRoleAuthorizationGroup, 0),
	}
}

// 权限列表
func (m *Proxy) AdminAuthorizationList() (*AckAdminAuthorizationList, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckAdminAuthorizationList
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/agency/admin/authorization/list",
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

// 获取角色
func (m *Proxy) AdminRoleGetByOrganizationIDByRoleID(OrganizationID uint32,
	RoleID uint32,
) (*AckAdminRoleGetByOrganizationIDByRoleID, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckAdminRoleGetByOrganizationIDByRoleID
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/agency/admin/role/get/:organization_id/:role_id",
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

type AskAdminRoleAddByOrganizationID struct {
	RoleName         string   `binding:"required,lte=64" db:"RoleName"`
	AuthorizationIds []uint32 `binding:"required,gt=0" db:"AuthorizationIds"`
}

func NewAskAdminRoleAddByOrganizationID() *AskAdminRoleAddByOrganizationID {
	return &AskAdminRoleAddByOrganizationID{
		AuthorizationIds: make([]uint32, 0),
	}
}

// 添加角色
func (m *Proxy) AdminRoleAddByOrganizationID(OrganizationID uint32,
	ask *AskAdminRoleAddByOrganizationID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/agency/admin/role/add/:organization_id",
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

type AskAdminRoleEditByOrganizationIDByRoleID struct {
	RoleName         string   `binding:"required,lte=64" db:"RoleName"`
	AuthorizationIds []uint32 `binding:"required,gt=0" db:"AuthorizationIds"`
}

func NewAskAdminRoleEditByOrganizationIDByRoleID() *AskAdminRoleEditByOrganizationIDByRoleID {
	return &AskAdminRoleEditByOrganizationIDByRoleID{
		AuthorizationIds: make([]uint32, 0),
	}
}

// 编辑角色
func (m *Proxy) AdminRoleEditByOrganizationIDByRoleID(OrganizationID uint32,
	RoleID uint32,
	ask *AskAdminRoleEditByOrganizationIDByRoleID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/agency/admin/role/edit/:organization_id/:role_id",
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

type AskAdminRoleDisableByOrganizationIDByRoleID struct {
	IsDisable bool `db:"IsDisable"`
}

func NewAskAdminRoleDisableByOrganizationIDByRoleID() *AskAdminRoleDisableByOrganizationIDByRoleID {
	return &AskAdminRoleDisableByOrganizationIDByRoleID{}
}

// 禁用角色
func (m *Proxy) AdminRoleDisableByOrganizationIDByRoleID(OrganizationID uint32,
	RoleID uint32,
	ask *AskAdminRoleDisableByOrganizationIDByRoleID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/agency/admin/role/disable/:organization_id/:role_id",
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
