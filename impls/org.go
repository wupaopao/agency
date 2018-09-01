package impls

import (
	"business/agency/cidl"
	"business/agency/common/db"
	"business/user/proxy/user"
	"common/api"

	"github.com/mz-eco/mz/http"
)

func init() {
	AddOrgOrganizationListHandler()
	AddOrgOrganizationInfoByIDHandler()

	AddOrgOrganizationEnableInfoByOrganizationIDHandler()
	AddOrgOrganizationGetOrganizationIDHandler()

	AddOrgRoleListByOrganizationIDHandler()
	AddOrgStaffListByOrganizationIDHandler()

	AddOrgStaffAddByOrganizationIDHandler()
	AddOrgStaffEditByOrganizationIDByUserIDHandler()
	AddOrgStaffDisableByOrganizationIDByUserIDHandler()

	AddOrgRoleGetByOrganizationIDByRoleIDHandler()
	AddOrgRoleAddByOrganizationIDHandler()
	AddOrgRoleEditByOrganizationIDByRoleIDHandler()
	AddOrgRoleDisableByOrganizationIDByRoleIDHandler()

	AddOrgAuthorizationListHandler()

}

type OrgOrganizationListImpl struct {
	cidl.ApiOrgOrganizationList
}

func AddOrgOrganizationListHandler() {
	AddHandler(
		cidl.META_ORG_ORGANIZATION_LIST,
		func() http.ApiHandler {
			return &OrgOrganizationListImpl{
				ApiOrgOrganizationList: cidl.MakeApiOrgOrganizationList(),
			}
		},
	)
}

func (m *OrgOrganizationListImpl) Handler(ctx *http.Context) {
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
type OrgOrganizationInfoByIDImpl struct {
	cidl.ApiOrgOrganizationInfoByID
}

func AddOrgOrganizationInfoByIDHandler() {
	AddHandler(
		cidl.META_ORG_ORGANIZATION_INFO_BY_ID,
		func() http.ApiHandler {
			return &OrgOrganizationInfoByIDImpl{
				ApiOrgOrganizationInfoByID: cidl.MakeApiOrgOrganizationInfoByID(),
			}
		},
	)
}

func (m *OrgOrganizationInfoByIDImpl) Handler(ctx *http.Context) {
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
		ctx.Errorf(api.ErrProxyFailed, "get manager user of organization failed. %s", err)
		return
	}

	m.Ack = &cidl.AckOrgOrganizationInfoByID{
		Organization: *organization,
	}
	m.Ack.ManagerWxNickname = manager.Nickname
	m.Ack.ManagerIdCardNumber = manager.IdCardNumber
	m.Ack.ManagerIdCardFront = manager.IdCardFront
	m.Ack.ManagerIdCardBack = manager.IdCardBack

	ctx.Json(m.Ack)
}

// 获取可用组织的纪录
type OrgOrganizationEnableInfoByOrganizationIDImpl struct {
	cidl.ApiOrgOrganizationEnableInfoByOrganizationID
}

func AddOrgOrganizationEnableInfoByOrganizationIDHandler() {
	AddHandler(
		cidl.META_ORG_ORGANIZATION_ENABLE_INFO_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &OrgOrganizationEnableInfoByOrganizationIDImpl{
				ApiOrgOrganizationEnableInfoByOrganizationID: cidl.MakeApiOrgOrganizationEnableInfoByOrganizationID(),
			}
		},
	)
}

func (m *OrgOrganizationEnableInfoByOrganizationIDImpl) Handler(ctx *http.Context) {
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
type OrgOrganizationGetOrganizationIDImpl struct {
	cidl.ApiOrgOrganizationGetOrganizationID
}

func AddOrgOrganizationGetOrganizationIDHandler() {
	AddHandler(
		cidl.META_ORG_ORGANIZATION_GET_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &OrgOrganizationGetOrganizationIDImpl{
				ApiOrgOrganizationGetOrganizationID: cidl.MakeApiOrgOrganizationGetOrganizationID(),
			}
		},
	)
}

func (m *OrgOrganizationGetOrganizationIDImpl) Handler(ctx *http.Context) {
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
type OrgRoleListByOrganizationIDImpl struct {
	cidl.ApiOrgRoleListByOrganizationID
}

func AddOrgRoleListByOrganizationIDHandler() {
	AddHandler(
		cidl.META_ORG_ROLE_LIST_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &OrgRoleListByOrganizationIDImpl{
				ApiOrgRoleListByOrganizationID: cidl.MakeApiOrgRoleListByOrganizationID(),
			}
		},
	)
}

func (m *OrgRoleListByOrganizationIDImpl) Handler(ctx *http.Context) {
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

	m.Ack.List, err = dbAgency.StaffRoleList(organizationId, cidl.StaffRoleTypeCommon, m.Query.Page, m.Query.PageSize, false)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get staff role list failed. %s", err)
		return
	}

	ctx.Json(m.Ack)
}

// 团购组织成员
type OrgStaffListByOrganizationIDImpl struct {
	cidl.ApiOrgStaffListByOrganizationID
}

func AddOrgStaffListByOrganizationIDHandler() {
	AddHandler(
		cidl.META_ORG_STAFF_LIST_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &OrgStaffListByOrganizationIDImpl{
				ApiOrgStaffListByOrganizationID: cidl.MakeApiOrgStaffListByOrganizationID(),
			}
		},
	)
}

func (m *OrgStaffListByOrganizationIDImpl) Handler(ctx *http.Context) {
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
type OrgStaffAddByOrganizationIDImpl struct {
	cidl.ApiOrgStaffAddByOrganizationID
}

func AddOrgStaffAddByOrganizationIDHandler() {
	AddHandler(
		cidl.META_ORG_STAFF_ADD_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &OrgStaffAddByOrganizationIDImpl{
				ApiOrgStaffAddByOrganizationID: cidl.MakeApiOrgStaffAddByOrganizationID(),
			}
		},
	)
}

func (m *OrgStaffAddByOrganizationIDImpl) Handler(ctx *http.Context) {
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
type OrgStaffEditByOrganizationIDByUserIDImpl struct {
	cidl.ApiOrgStaffEditByOrganizationIDByUserID
}

func AddOrgStaffEditByOrganizationIDByUserIDHandler() {
	AddHandler(
		cidl.META_ORG_STAFF_EDIT_BY_ORGANIZATION_ID_BY_USER_ID,
		func() http.ApiHandler {
			return &OrgStaffEditByOrganizationIDByUserIDImpl{
				ApiOrgStaffEditByOrganizationIDByUserID: cidl.MakeApiOrgStaffEditByOrganizationIDByUserID(),
			}
		},
	)
}

func (m *OrgStaffEditByOrganizationIDByUserIDImpl) Handler(ctx *http.Context) {
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
type OrgStaffDisableByOrganizationIDByUserIDImpl struct {
	cidl.ApiOrgStaffDisableByOrganizationIDByUserID
}

func AddOrgStaffDisableByOrganizationIDByUserIDHandler() {
	AddHandler(
		cidl.META_ORG_STAFF_DISABLE_BY_ORGANIZATION_ID_BY_USER_ID,
		func() http.ApiHandler {
			return &OrgStaffDisableByOrganizationIDByUserIDImpl{
				ApiOrgStaffDisableByOrganizationIDByUserID: cidl.MakeApiOrgStaffDisableByOrganizationIDByUserID(),
			}
		},
	)
}

func (m *OrgStaffDisableByOrganizationIDByUserIDImpl) Handler(ctx *http.Context) {
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

	askDisable := &user.AskInnerUserSetIsDisableByUserID{
                UserType:   user.OrgStaff,
                IsDisable: m.Ask.IsDisable,
        }
	_, err = user.NewProxy("user-service").InnerUserSetIsDisableByUserID(m.Params.UserID, askDisable)
	if err != nil {
		ctx.Errorf(api.ErrProxyFailed, "set is_disable by user_id failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 获得角色
type OrgRoleGetByOrganizationIDByRoleIDImpl struct {
	cidl.ApiOrgRoleGetByOrganizationIDByRoleID
}

func AddOrgRoleGetByOrganizationIDByRoleIDHandler() {
	AddHandler(
		cidl.META_ORG_ROLE_GET_BY_ORGANIZATION_ID_BY_ROLE_ID,
		func() http.ApiHandler {
			return &OrgRoleGetByOrganizationIDByRoleIDImpl{
				ApiOrgRoleGetByOrganizationIDByRoleID: cidl.MakeApiOrgRoleGetByOrganizationIDByRoleID(),
			}
		},
	)
}

func (m *OrgRoleGetByOrganizationIDByRoleIDImpl) Handler(ctx *http.Context) {
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
type OrgRoleAddByOrganizationIDImpl struct {
	cidl.ApiOrgRoleAddByOrganizationID
}

func AddOrgRoleAddByOrganizationIDHandler() {
	AddHandler(
		cidl.META_ORG_ROLE_ADD_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &OrgRoleAddByOrganizationIDImpl{
				ApiOrgRoleAddByOrganizationID: cidl.MakeApiOrgRoleAddByOrganizationID(),
			}
		},
	)
}

func (m *OrgRoleAddByOrganizationIDImpl) Handler(ctx *http.Context) {
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
type OrgRoleEditByOrganizationIDByRoleIDImpl struct {
	cidl.ApiOrgRoleEditByOrganizationIDByRoleID
}

func AddOrgRoleEditByOrganizationIDByRoleIDHandler() {
	AddHandler(
		cidl.META_ORG_ROLE_EDIT_BY_ORGANIZATION_ID_BY_ROLE_ID,
		func() http.ApiHandler {
			return &OrgRoleEditByOrganizationIDByRoleIDImpl{
				ApiOrgRoleEditByOrganizationIDByRoleID: cidl.MakeApiOrgRoleEditByOrganizationIDByRoleID(),
			}
		},
	)
}

func (m *OrgRoleEditByOrganizationIDByRoleIDImpl) Handler(ctx *http.Context) {
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

type OrgAuthorizationListImpl struct {
	cidl.ApiOrgAuthorizationList
}

func AddOrgAuthorizationListHandler() {
	AddHandler(
		cidl.META_ORG_AUTHORIZATION_LIST,
		func() http.ApiHandler {
			return &OrgAuthorizationListImpl{
				ApiOrgAuthorizationList: cidl.MakeApiOrgAuthorizationList(),
			}
		},
	)
}

func (m *OrgAuthorizationListImpl) Handler(ctx *http.Context) {
	m.Ack.Modules = cidl.NewStaffRoleAuthorizationGroupsByRole(&cidl.StaffRole{
		RoleAuthorization: cidl.NewRoleAuthorizationMap(),
	})
	ctx.Json(m.Ack)
}

// 禁用角色
type OrgRoleDisableByOrganizationIDByRoleIDImpl struct {
	cidl.ApiOrgRoleDisableByOrganizationIDByRoleID
}

func AddOrgRoleDisableByOrganizationIDByRoleIDHandler() {
	AddHandler(
		cidl.META_ORG_ROLE_DISABLE_BY_ORGANIZATION_ID_BY_ROLE_ID,
		func() http.ApiHandler {
			return &OrgRoleDisableByOrganizationIDByRoleIDImpl{
				ApiOrgRoleDisableByOrganizationIDByRoleID: cidl.MakeApiOrgRoleDisableByOrganizationIDByRoleID(),
			}
		},
	)
}

func (m *OrgRoleDisableByOrganizationIDByRoleIDImpl) Handler(ctx *http.Context) {
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
