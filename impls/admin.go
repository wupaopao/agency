package impls

import (
	"fmt"
	"time"

	"business/agency/cidl"
	"business/agency/common/db"
	"business/user/proxy/user"
	"common/api"
	"common/file"

	"github.com/mz-eco/mz/http"
	"github.com/mz-eco/mz/utils"
)

func init() {
	AddAdminOrganizationListHandler()
	AddAdminOrganizationInfoByIDHandler()

	AddAdminOrganizationAddPicTokenHandler()
	AddAdminOrganizationAddHandler()
	AddAdminOrganizationEditByOrganizationIDHandler()
	AddAdminOrganizationDisableByOrganizationIDHandler()
	AddAdminOrganizationEnableInfoByOrganizationIDHandler()
	AddAdminOrganizationGetOrganizationIDHandler()

	AddAdminRoleListByOrganizationIDHandler()
	AddAdminStaffListByOrganizationIDHandler()

	AddAdminStaffAddByOrganizationIDHandler()
	AddAdminStaffEditByOrganizationIDByUserIDHandler()
	AddAdminStaffDisableByOrganizationIDByUserIDHandler()

	AddAdminRoleGetByOrganizationIDByRoleIDHandler()
	AddAdminRoleAddByOrganizationIDHandler()
	AddAdminRoleEditByOrganizationIDByRoleIDHandler()
	AddAdminRoleDisableByOrganizationIDByRoleIDHandler()

	AddAdminAuthorizationListHandler()

}

type AdminOrganizationListImpl struct {
	cidl.ApiAdminOrganizationList
}

func AddAdminOrganizationListHandler() {
	AddHandler(
		cidl.META_ADMIN_ORGANIZATION_LIST,
		func() http.ApiHandler {
			return &AdminOrganizationListImpl{
				ApiAdminOrganizationList: cidl.MakeApiAdminOrganizationList(),
			}
		},
	)
}

func (m *AdminOrganizationListImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	ack := m.Ack
	dbAgency := db.NewMallAgency()
	if m.Query.Search == "" {
		ack.Count, err = dbAgency.OrganizationCount()
	} else {
		ack.Count, err = dbAgency.OrganizationSearchCount(m.Query.Search)
	}

	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get organization count failed. %s", err)
		return
	}

	if ack.Count == 0 {
		ctx.Json(m.Ack)
		return
	}

	if m.Query.Search == "" {
		ack.List, err = dbAgency.OrganizationList(m.Query.Page, m.Query.PageSize, false)
	} else {
		ack.List, err = dbAgency.OrganizationSearchList(m.Query.Page, m.Query.PageSize, m.Query.Search, false)
	}

	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get organization list failed. %s", err)
		return
	}

	ctx.Json(ack)
}

// 获取团购组织详细信息
type AdminOrganizationInfoByIDImpl struct {
	cidl.ApiAdminOrganizationInfoByID
}

func AddAdminOrganizationInfoByIDHandler() {
	AddHandler(
		cidl.META_ADMIN_ORGANIZATION_INFO_BY_ID,
		func() http.ApiHandler {
			return &AdminOrganizationInfoByIDImpl{
				ApiAdminOrganizationInfoByID: cidl.MakeApiAdminOrganizationInfoByID(),
			}
		},
	)
}

func (m *AdminOrganizationInfoByIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	organizationId := m.Params.OrganizationID
	dbAgency := db.NewMallAgency()
	organization, err := dbAgency.GetOrganization(organizationId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get organization from db failed. %s", err)
		return
	}

	manager, err := user.NewProxy("user-service").InnerUserInfoByUserID(organization.ManagerUserId)
	if err != nil {
		ctx.ProxyErrorf(err, "get manager user of organization failed. %s", err)
		return
	}

	m.Ack = &cidl.AckAdminOrganizationInfoByID{
		Organization: *organization,
	}
	m.Ack.ManagerWxNickname = manager.Nickname
	m.Ack.ManagerIdCardNumber = manager.IdCardNumber
	m.Ack.ManagerIdCardFront = manager.IdCardFront
	m.Ack.ManagerIdCardBack = manager.IdCardBack

	ctx.Json(m.Ack)
}

// 获取图片上传TOKEN
type AdminOrganizationAddPicTokenImpl struct {
	cidl.ApiAdminOrganizationAddPicToken
}

func AddAdminOrganizationAddPicTokenHandler() {
	AddHandler(
		cidl.META_ADMIN_ORGANIZATION_ADD_PIC_TOKEN,
		func() http.ApiHandler {
			return &AdminOrganizationAddPicTokenImpl{
				ApiAdminOrganizationAddPicToken: cidl.MakeApiAdminOrganizationAddPicToken(),
			}
		},
	)
}

func (m *AdminOrganizationAddPicTokenImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	today, err := utils.DayStartTime(time.Now())
	if err != nil {
		ctx.Errorf(api.ErrServer, "get day start time failed. %s", err)
		return
	}

	qiniu, err := file.GetQiniuPublicBucket()
	if err != nil {
		ctx.Errorf(api.ErrServer, "get qiniu public bucket failed. %s", err)
		return
	}

	prefix := fmt.Sprintf("bilimall/organization/%d/", today.Unix())
	for _, fileName := range m.Ask.FileNames {
		if fileName == "" {
			ctx.Errorf(api.ErrWrongParams, "empty pic file name. %s", err)
			return
		}

		token, key, err := qiniu.GenerateUploadToken(fileName, prefix)
		if err != nil {
			return
		}

		storeUrl := qiniu.StoreUrl(key)
		m.Ack.Tokens = append(m.Ack.Tokens, &cidl.AckPicToken{
			OriginalFileName: fileName,
			Token:            token,
			Key:              key,
			StoreUrl:         storeUrl,
			AccessUrl:        storeUrl,
		})

	}

	ctx.Json(m.Ack)
}

// 添加组织
type AdminOrganizationAddImpl struct {
	cidl.ApiAdminOrganizationAdd
}

func AddAdminOrganizationAddHandler() {
	AddHandler(
		cidl.META_ADMIN_ORGANIZATION_ADD,
		func() http.ApiHandler {
			return &AdminOrganizationAddImpl{
				ApiAdminOrganizationAdd: cidl.MakeApiAdminOrganizationAdd(),
			}
		},
	)
}

func (m *AdminOrganizationAddImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	askAddOrgManager := &user.AskInnerUserUserOrgManagerAddOrUpdate{
		Name:         m.Ask.ManagerName,
		Mobile:       m.Ask.ManagerMobile,
		Nickname:     m.Ask.ManagerNickname,
		IdCardNumber: m.Ask.ManagerIdCardNumber,
		IdCardFront:  m.Ask.ManagerIdCardFront,
		IdCardBack:   m.Ask.ManagerIdCardBack,
	}
	ackAddOrgManager, err := user.NewProxy("user-service").InnerUserUserOrgManagerAddOrUpdate(askAddOrgManager)
	if err != nil {
		ctx.ProxyErrorf(err, "add or update org manager user failed. %s", err)
		return
	}

	managerUid := ackAddOrgManager.UserId
	organization := &cidl.Organization{
		Name:            m.Ask.Name,
		Logo:            m.Ask.Logo,
		Province:        m.Ask.Province,
		City:            m.Ask.City,
		Address:         m.Ask.Address,
		PostCode:        m.Ask.PostCode,
		BankName:        m.Ask.BankName,
		BankAccount:     m.Ask.BankAccount,
		BankAccountName: m.Ask.BankAccountName,
		CompanyName:     m.Ask.CompanyName,
		LicenseNumber:   m.Ask.LicenseNumber,
		LicensePicture:  m.Ask.LicensePicture,
		ManagerUserId:   managerUid,
		ManagerName:     m.Ask.ManagerName,
		ManagerMobile:   m.Ask.ManagerMobile,
		GroupBuyingMode: m.Ask.GroupBuyingMode,
		PerfectionState: cidl.OrganizationPerfectionStateComplete,
		IsDisable:       false,
	}

	dbAgency := db.NewMallAgency()
	result, err := dbAgency.AddOrganization(organization)
	if err != nil {
		ctx.Errorf(api.ErrDBInsertFailed, "add organization failed. %s", err)
		return
	}

	organizationId, err := result.LastInsertId()
	if err != nil {
		ctx.Errorf(api.ErrDBInsertFailed, "get last insert id failed. %s", err)
		return
	}

	organization.OrganizationId = uint32(organizationId)

	// 添加超级管理员
	_, err = dbAgency.AddSuperAdministrator(organization.OrganizationId, &cidl.Staff{
		UserId:           organization.ManagerUserId,
		OrganizationId:   organization.OrganizationId,
		OrganizationName: organization.Name,
		Name:             organization.ManagerName,
		Mobile:           organization.ManagerMobile,
	})

	if err != nil {
		ctx.Errorf(api.ErrDBInsertFailed, "add super administrator failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 编辑组织
type AdminOrganizationEditByOrganizationIDImpl struct {
	cidl.ApiAdminOrganizationEditByOrganizationID
}

func AddAdminOrganizationEditByOrganizationIDHandler() {
	AddHandler(
		cidl.META_ADMIN_ORGANIZATION_EDIT_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &AdminOrganizationEditByOrganizationIDImpl{
				ApiAdminOrganizationEditByOrganizationID: cidl.MakeApiAdminOrganizationEditByOrganizationID(),
			}
		},
	)
}

func (m *AdminOrganizationEditByOrganizationIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	dbAgency := db.NewMallAgency()
	organizationId := m.Params.OrganizationID
	oldOrganization, err := dbAgency.GetOrganization(organizationId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get organization failed. %s", err)
		return
	}

	var managerUid string

	userProxy := user.NewProxy("user-service")
	if m.Ask.ManagerMobile == oldOrganization.ManagerMobile {
		managerUid = oldOrganization.ManagerUserId
		askUpdateOrgManager := &user.AskInnerUserUserOrgManagerUpdate{
			UserId:       oldOrganization.ManagerUserId,
			Name:         m.Ask.ManagerName,
			Mobile:       m.Ask.ManagerMobile,
			Nickname:     m.Ask.ManagerNickname,
			IdCardNumber: m.Ask.ManagerIdCardNumber,
			IdCardFront:  m.Ask.ManagerIdCardFront,
			IdCardBack:   m.Ask.ManagerIdCardBack,
		}

		_, err = userProxy.InnerUserUserOrgManagerUpdate(askUpdateOrgManager)
		if err != nil {
			ctx.ProxyErrorf(err, "add or update org manager user failed. %s", err)
			return
		}

	} else {
		askAddOrgManager := &user.AskInnerUserUserOrgManagerAddOrUpdate{
			Name:         m.Ask.ManagerName,
			Mobile:       m.Ask.ManagerMobile,
			Nickname:     m.Ask.ManagerNickname,
			IdCardNumber: m.Ask.ManagerIdCardNumber,
			IdCardFront:  m.Ask.ManagerIdCardFront,
			IdCardBack:   m.Ask.ManagerIdCardBack,
		}

		ackAddOrgManager, errAdd := userProxy.InnerUserUserOrgManagerAddOrUpdate(askAddOrgManager)
		err = errAdd
		if err != nil {
			ctx.ProxyErrorf(err, "add or update org manager user failed. %s", err)
			return
		}

		managerUid = ackAddOrgManager.UserId
	}

	// 更新组织
	organization := &cidl.Organization{
		OrganizationId:  organizationId,
		Name:            m.Ask.Name,
		Logo:            m.Ask.Logo,
		Province:        m.Ask.Province,
		City:            m.Ask.City,
		Address:         m.Ask.Address,
		PostCode:        m.Ask.PostCode,
		BankName:        m.Ask.BankName,
		BankAccount:     m.Ask.BankAccount,
		BankAccountName: m.Ask.BankAccountName,
		CompanyName:     m.Ask.CompanyName,
		LicenseNumber:   m.Ask.LicenseNumber,
		LicensePicture:  m.Ask.LicensePicture,
		ManagerUserId:   managerUid,
		ManagerName:     m.Ask.ManagerName,
		ManagerMobile:   m.Ask.ManagerMobile,
		GroupBuyingMode: m.Ask.GroupBuyingMode,
		PerfectionState: cidl.OrganizationPerfectionStateComplete,
	}

	_, err = dbAgency.UpdateOrganization(organization)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update organization failed. %s", err)
		return
	}

	// 更换管理员
	if managerUid != oldOrganization.ManagerUserId {
		_, err = userProxy.InnerUserUserOrgManagerUnbind(&user.AskInnerUserUserOrgManagerUnbind{
			OldManagerUid: oldOrganization.ManagerUserId,
		})
		if err != nil {
			ctx.ProxyErrorf(err, "unbind org manager failed. %s", err)
			return
		}
	}

	ctx.Succeed()
}

// 禁用组织
type AdminOrganizationDisableByOrganizationIDImpl struct {
	cidl.ApiAdminOrganizationDisableByOrganizationID
}

func AddAdminOrganizationDisableByOrganizationIDHandler() {
	AddHandler(
		cidl.META_ADMIN_ORGANIZATION_DISABLE_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &AdminOrganizationDisableByOrganizationIDImpl{
				ApiAdminOrganizationDisableByOrganizationID: cidl.MakeApiAdminOrganizationDisableByOrganizationID(),
			}
		},
	)
}

func (m *AdminOrganizationDisableByOrganizationIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	dbAgency := db.NewMallAgency()
	strSql := `UPDATE agc_organization SET is_disable=? WHERE org_id=?`
	_, err = dbAgency.DB.Exec(strSql, m.Ask.IsDisable, m.Params.OrganizationID)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update organization is_disable failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 获取可用组织的纪录
type AdminOrganizationEnableInfoByOrganizationIDImpl struct {
	cidl.ApiAdminOrganizationEnableInfoByOrganizationID
}

func AddAdminOrganizationEnableInfoByOrganizationIDHandler() {
	AddHandler(
		cidl.META_ADMIN_ORGANIZATION_ENABLE_INFO_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &AdminOrganizationEnableInfoByOrganizationIDImpl{
				ApiAdminOrganizationEnableInfoByOrganizationID: cidl.MakeApiAdminOrganizationEnableInfoByOrganizationID(),
			}
		},
	)
}

func (m *AdminOrganizationEnableInfoByOrganizationIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	organizationId := m.Params.OrganizationID
	dbAgency := db.NewMallAgency()
	organization, err := dbAgency.GetEnableOrganization(organizationId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get enable organization failed. %s", err)
		return
	}

	m.Ack = organization

	ctx.Json(m.Ack)
}

// 通过用户ID获取团购组织ID
type AdminOrganizationGetOrganizationIDImpl struct {
	cidl.ApiAdminOrganizationGetOrganizationID
}

func AddAdminOrganizationGetOrganizationIDHandler() {
	AddHandler(
		cidl.META_ADMIN_ORGANIZATION_GET_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &AdminOrganizationGetOrganizationIDImpl{
				ApiAdminOrganizationGetOrganizationID: cidl.MakeApiAdminOrganizationGetOrganizationID(),
			}
		},
	)
}

func (m *AdminOrganizationGetOrganizationIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	userId := m.Query.UserID
	m.Ack.OrganizationId, err = db.NewMallAgency().GetOrganizationId(userId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get organization id failed. %s", err)
		return
	}

	ctx.Json(m.Ack)
}

// 团购组织角色权限列表
type AdminRoleListByOrganizationIDImpl struct {
	cidl.ApiAdminRoleListByOrganizationID
}

func AddAdminRoleListByOrganizationIDHandler() {
	AddHandler(
		cidl.META_ADMIN_ROLE_LIST_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &AdminRoleListByOrganizationIDImpl{
				ApiAdminRoleListByOrganizationID: cidl.MakeApiAdminRoleListByOrganizationID(),
			}
		},
	)
}

func (m *AdminRoleListByOrganizationIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	dbAgency := db.NewMallAgency()
	organizationId := m.Params.OrganizationID
	m.Ack.Count, err = dbAgency.StaffRoleCount(organizationId, cidl.StaffRoleTypeCommon)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get staff role count failed. %s", err)
		return
	}

	if m.Ack.Count == 0 {
		ctx.Json(m.Ack)
		return
	}

	m.Ack.List, err = dbAgency.StaffRoleList(organizationId, cidl.StaffRoleTypeCommon, m.Query.Page, m.Query.PageSize, true)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get staff role list failed. %s", err)
		return
	}

	ctx.Json(m.Ack)
}

// 团购组织成员
type AdminStaffListByOrganizationIDImpl struct {
	cidl.ApiAdminStaffListByOrganizationID
}

func AddAdminStaffListByOrganizationIDHandler() {
	AddHandler(
		cidl.META_ADMIN_STAFF_LIST_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &AdminStaffListByOrganizationIDImpl{
				ApiAdminStaffListByOrganizationID: cidl.MakeApiAdminStaffListByOrganizationID(),
			}
		},
	)
}

func (m *AdminStaffListByOrganizationIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	ack := m.Ack
	dbAgency := db.NewMallAgency()
	organizationId := m.Params.OrganizationID
	ack.Count, err = dbAgency.StaffCount(organizationId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get staff count failed. %s", err)
		return
	}

	if ack.Count == 0 {
		ctx.Json(m.Ack)
		return
	}

	ack.List, err = dbAgency.StaffList(organizationId, m.Query.Page, m.Query.PageSize, false)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get staff list failed. %s", err)
		return
	}

	ctx.Json(ack)
}

// 添加团购组织成员
type AdminStaffAddByOrganizationIDImpl struct {
	cidl.ApiAdminStaffAddByOrganizationID
}

func AddAdminStaffAddByOrganizationIDHandler() {
	AddHandler(
		cidl.META_ADMIN_STAFF_ADD_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &AdminStaffAddByOrganizationIDImpl{
				ApiAdminStaffAddByOrganizationID: cidl.MakeApiAdminStaffAddByOrganizationID(),
			}
		},
	)
}

func (m *AdminStaffAddByOrganizationIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	askAddOrgStaff := &user.AskInnerUserUserOrgStaffAddOrUpdate{
		Name:   m.Ask.Name,
		Mobile: m.Ask.Mobile,
	}
	ackAddOrgStaff, err := user.NewProxy("user-service").InnerUserUserOrgStaffAddOrUpdate(askAddOrgStaff)
	if err != nil {
		ctx.Errorf(api.ErrProxyFailed, "add or update org manager user failed. %s", err)
		return
	}

	staffUid := ackAddOrgStaff.UserId
	organizationId := m.Params.OrganizationID

	dbAgency := db.NewMallAgency()

	organization, err := dbAgency.GetOrganization(organizationId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get organization failed. %s", err)
		return
	}

	staffRole, err := dbAgency.GetStaffRoleByOrgIdRoleId(organizationId, m.Ask.RoleId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get staff role id failed. %s", err)
		return
	}

	// 不能添加超级管理员权限角色
	for authorizationId, _ := range *staffRole.RoleAuthorization {
		if authorizationId == cidl.SuperAdministratorAuthorizationId {
			ctx.Errorf(cidl.ErrAddSuperAdministratorIsForbidden, "forbidden to add super administrator.")
			return
		}
	}

	staff := &cidl.Staff{
		UserId:           staffUid,
		OrganizationId:   organization.OrganizationId,
		OrganizationName: organization.Name,
		Name:             m.Ask.Name,
		Mobile:           m.Ask.Mobile,
		RoleId:           staffRole.RoleId,
		RoleName:         staffRole.RoleName,
	}

	_, err = dbAgency.AddStaff(staff)
	if err != nil {
		ctx.Errorf(api.ErrDBInsertFailed, "add staff failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 编辑团购组织成员
type AdminStaffEditByOrganizationIDByUserIDImpl struct {
	cidl.ApiAdminStaffEditByOrganizationIDByUserID
}

func AddAdminStaffEditByOrganizationIDByUserIDHandler() {
	AddHandler(
		cidl.META_ADMIN_STAFF_EDIT_BY_ORGANIZATION_ID_BY_USER_ID,
		func() http.ApiHandler {
			return &AdminStaffEditByOrganizationIDByUserIDImpl{
				ApiAdminStaffEditByOrganizationIDByUserID: cidl.MakeApiAdminStaffEditByOrganizationIDByUserID(),
			}
		},
	)
}

func (m *AdminStaffEditByOrganizationIDByUserIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	userId := m.Params.UserID
	_, err = user.NewProxy("user-service").InnerUserUserOrgStaffUpdateByUserID(userId, &user.AskInnerUserUserOrgStaffUpdateByUserID{
		Name:   m.Ask.Name,
		Mobile: m.Ask.Mobile,
	})

	if err != nil {
		ctx.Errorf(api.ErrProxyFailed, "update user name and mobile by proxy failed. %s", err)
		return
	}

	organizationId := m.Params.OrganizationID
	dbAgency := db.NewMallAgency()
	staffRole, err := dbAgency.GetStaffRoleByOrgIdRoleId(organizationId, m.Ask.RoleId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get staff role id failed. %s", err)
		return
	}

	strSql := `UPDATE agc_staff SET rol_id=?, rol_name=? WHERE uid=?`
	_, err = dbAgency.DB.Exec(strSql, staffRole.RoleId, staffRole.RoleName, userId)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update staff failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 禁用成员
type AdminStaffDisableByOrganizationIDByUserIDImpl struct {
	cidl.ApiAdminStaffDisableByOrganizationIDByUserID
}

func AddAdminStaffDisableByOrganizationIDByUserIDHandler() {
	AddHandler(
		cidl.META_ADMIN_STAFF_DISABLE_BY_ORGANIZATION_ID_BY_USER_ID,
		func() http.ApiHandler {
			return &AdminStaffDisableByOrganizationIDByUserIDImpl{
				ApiAdminStaffDisableByOrganizationIDByUserID: cidl.MakeApiAdminStaffDisableByOrganizationIDByUserID(),
			}
		},
	)
}

func (m *AdminStaffDisableByOrganizationIDByUserIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	dbAgency := db.NewMallAgency()
	strSql := `UPDATE agc_staff SET is_disable=? WHERE uid=? AND org_id=?`
	_, err = dbAgency.DB.Exec(strSql, m.Ask.IsDisable, m.Params.UserID, m.Params.OrganizationID)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update staff is_disable failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 获得角色
type AdminRoleGetByOrganizationIDByRoleIDImpl struct {
	cidl.ApiAdminRoleGetByOrganizationIDByRoleID
}

func AddAdminRoleGetByOrganizationIDByRoleIDHandler() {
	AddHandler(
		cidl.META_ADMIN_ROLE_GET_BY_ORGANIZATION_ID_BY_ROLE_ID,
		func() http.ApiHandler {
			return &AdminRoleGetByOrganizationIDByRoleIDImpl{
				ApiAdminRoleGetByOrganizationIDByRoleID: cidl.MakeApiAdminRoleGetByOrganizationIDByRoleID(),
			}
		},
	)
}

func (m *AdminRoleGetByOrganizationIDByRoleIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	dbAgency := db.NewMallAgency()
	m.Ack.StaffRole, err = dbAgency.GetStaffRoleByOrgIdRoleId(m.Params.OrganizationID, m.Params.RoleID)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get staff role failed. %s", err)
		return
	}

	m.Ack.Modules = cidl.NewStaffRoleAuthorizationGroupsByRole(m.Ack.StaffRole)
	ctx.Json(m.Ack)
}

// 添加角色
type AdminRoleAddByOrganizationIDImpl struct {
	cidl.ApiAdminRoleAddByOrganizationID
}

func AddAdminRoleAddByOrganizationIDHandler() {
	AddHandler(
		cidl.META_ADMIN_ROLE_ADD_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &AdminRoleAddByOrganizationIDImpl{
				ApiAdminRoleAddByOrganizationID: cidl.MakeApiAdminRoleAddByOrganizationID(),
			}
		},
	)
}

func (m *AdminRoleAddByOrganizationIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	var roleAuthorizations = cidl.NewRoleAuthorizationMap()
	for _, id := range m.Ask.AuthorizationIds {
		authorizationId := cidl.AuthorizationId(id)
		roleAuthorization, ok := cidl.RoleAuthorizationMapIds[authorizationId]
		if !ok {
			ctx.Errorf(api.ErrWrongParams, "wrong authorization id.")
			return
		}

		(*roleAuthorizations)[authorizationId] = roleAuthorization
	}

	staffRole := &cidl.StaffRole{
		OrganizationId:    m.Params.OrganizationID,
		RoleName:          m.Ask.RoleName,
		RoleAuthorization: roleAuthorizations,
		Type:              cidl.StaffRoleTypeCommon,
		Version:           cidl.StaffRoleRecordVersion,
	}

	_, err = db.NewMallAgency().AddStaffRole(staffRole)
	if err != nil {
		ctx.Errorf(api.ErrDBInsertFailed, "add staff role failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 编辑角色
type AdminRoleEditByOrganizationIDByRoleIDImpl struct {
	cidl.ApiAdminRoleEditByOrganizationIDByRoleID
}

func AddAdminRoleEditByOrganizationIDByRoleIDHandler() {
	AddHandler(
		cidl.META_ADMIN_ROLE_EDIT_BY_ORGANIZATION_ID_BY_ROLE_ID,
		func() http.ApiHandler {
			return &AdminRoleEditByOrganizationIDByRoleIDImpl{
				ApiAdminRoleEditByOrganizationIDByRoleID: cidl.MakeApiAdminRoleEditByOrganizationIDByRoleID(),
			}
		},
	)
}

func (m *AdminRoleEditByOrganizationIDByRoleIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	var roleAuthorizations = cidl.NewRoleAuthorizationMap()
	for _, id := range m.Ask.AuthorizationIds {
		authorizationId := cidl.AuthorizationId(id)
		roleAuthorization, ok := cidl.RoleAuthorizationMapIds[authorizationId]
		if !ok {
			ctx.Errorf(api.ErrWrongParams, "wrong authorization id.")
			return
		}

		(*roleAuthorizations)[authorizationId] = roleAuthorization
	}

	_, err = db.NewMallAgency().UpdateStaffRole(m.Params.OrganizationID, m.Params.RoleID, m.Ask.RoleName, roleAuthorizations)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update staff role failed. %s", err)
		return
	}

	ctx.Succeed()
}

type AdminAuthorizationListImpl struct {
	cidl.ApiAdminAuthorizationList
}

func AddAdminAuthorizationListHandler() {
	AddHandler(
		cidl.META_ADMIN_AUTHORIZATION_LIST,
		func() http.ApiHandler {
			return &AdminAuthorizationListImpl{
				ApiAdminAuthorizationList: cidl.MakeApiAdminAuthorizationList(),
			}
		},
	)
}

func (m *AdminAuthorizationListImpl) Handler(ctx *http.Context) {
	m.Ack.Modules = cidl.NewStaffRoleAuthorizationGroupsByRole(&cidl.StaffRole{
		RoleAuthorization: cidl.NewRoleAuthorizationMap(),
	})
	ctx.Json(m.Ack)
}

// 禁用角色
type AdminRoleDisableByOrganizationIDByRoleIDImpl struct {
	cidl.ApiAdminRoleDisableByOrganizationIDByRoleID
}

func AddAdminRoleDisableByOrganizationIDByRoleIDHandler() {
	AddHandler(
		cidl.META_ADMIN_ROLE_DISABLE_BY_ORGANIZATION_ID_BY_ROLE_ID,
		func() http.ApiHandler {
			return &AdminRoleDisableByOrganizationIDByRoleIDImpl{
				ApiAdminRoleDisableByOrganizationIDByRoleID: cidl.MakeApiAdminRoleDisableByOrganizationIDByRoleID(),
			}
		},
	)
}

func (m *AdminRoleDisableByOrganizationIDByRoleIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	_, err = db.NewMallAgency().SetStaffRoleDisable(m.Params.OrganizationID, m.Params.RoleID, m.Ask.IsDisable)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update staff role is_disable failed. %s", err)
		return
	}

	ctx.Succeed()
}
