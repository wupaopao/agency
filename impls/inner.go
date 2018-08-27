package impls

import (
	"business/agency/cidl"
	"business/agency/common/db"
	"common/api"

	"github.com/mz-eco/mz/http"
)

func init() {
	AddInnerAgencyOrganizationInfoByOrganizationIDHandler()
	AddInnerAgencyOrganizationInfoByUserIDByUserIDHandler()
	AddInnerAgencyStaffInfoByUserIDHandler()
	AddInnerAgencyStaffIsDisableByUserIDHandler()

}

type InnerAgencyOrganizationInfoByOrganizationIDImpl struct {
	cidl.ApiInnerAgencyOrganizationInfoByOrganizationID
}

func AddInnerAgencyOrganizationInfoByOrganizationIDHandler() {
	AddHandler(
		cidl.META_INNER_AGENCY_ORGANIZATION_INFO_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &InnerAgencyOrganizationInfoByOrganizationIDImpl{
				ApiInnerAgencyOrganizationInfoByOrganizationID: cidl.MakeApiInnerAgencyOrganizationInfoByOrganizationID(),
			}
		},
	)
}

func (m *InnerAgencyOrganizationInfoByOrganizationIDImpl) Handler(ctx *http.Context) {
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
	m.Ack = organization
	ctx.Json(m.Ack)
}

// 获取可用团购组织
type InnerAgencyOrganizationEnableInfoByOrganizationIDImpl struct {
	cidl.ApiInnerAgencyOrganizationEnableInfoByOrganizationID
}

func AddInnerAgencyOrganizationEnableInfoByOrganizationIDHandler() {
	AddHandler(
		cidl.META_INNER_AGENCY_ORGANIZATION_ENABLE_INFO_BY_ORGANIZATION_ID,
		func() http.ApiHandler {
			return &InnerAgencyOrganizationEnableInfoByOrganizationIDImpl{
				ApiInnerAgencyOrganizationEnableInfoByOrganizationID: cidl.MakeApiInnerAgencyOrganizationEnableInfoByOrganizationID(),
			}
		},
	)
}

func (m *InnerAgencyOrganizationEnableInfoByOrganizationIDImpl) Handler(ctx *http.Context) {
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

// 通过uid获取团购组织
type InnerAgencyOrganizationInfoByUserIDByUserIDImpl struct {
	cidl.ApiInnerAgencyOrganizationInfoByUserIDByUserID
}

func AddInnerAgencyOrganizationInfoByUserIDByUserIDHandler() {
	AddHandler(
		cidl.META_INNER_AGENCY_ORGANIZATION_INFO_BY_USER_ID_BY_USER_ID,
		func() http.ApiHandler {
			return &InnerAgencyOrganizationInfoByUserIDByUserIDImpl{
				ApiInnerAgencyOrganizationInfoByUserIDByUserID: cidl.MakeApiInnerAgencyOrganizationInfoByUserIDByUserID(),
			}
		},
	)
}

func (m *InnerAgencyOrganizationInfoByUserIDByUserIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	userID := m.Params.UserID
	dbAgency := db.NewMallAgency()
	organization, err := dbAgency.GetOrganizationByUserId(userID)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get organization from db failed. %s", err)
		return
	}
	m.Ack = organization
	ctx.Json(m.Ack)
}

type InnerAgencyStaffInfoByUserIDImpl struct {
	cidl.ApiInnerAgencyStaffInfoByUserID
}

func AddInnerAgencyStaffInfoByUserIDHandler() {
	AddHandler(
		cidl.META_INNER_AGENCY_STAFF_INFO_BY_USER_ID,
		func() http.ApiHandler {
			return &InnerAgencyStaffInfoByUserIDImpl{
				ApiInnerAgencyStaffInfoByUserID: cidl.MakeApiInnerAgencyStaffInfoByUserID(),
			}
		},
	)
}

func (m *InnerAgencyStaffInfoByUserIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	userId := m.Params.UserID
	dbAgency := db.NewMallAgency()
	staff, err := dbAgency.GetStaff(userId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get staff failed. %s", err)
		return
	}

	staffRole, err := dbAgency.GetStaffRoleByOrgIdRoleId(staff.OrganizationId, staff.RoleId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get staff role failed. %s", err)
		return
	}

	organization, err := dbAgency.GetOrganization(staff.OrganizationId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get organization failed. %s", err)
		return
	}

	m.Ack.Staff = staff
	m.Ack.StaffRole = staffRole
	m.Ack.Organization = organization

	ctx.Json(m.Ack)
}

// 是否被禁用
type InnerAgencyStaffIsDisableByUserIDImpl struct {
	cidl.ApiInnerAgencyStaffIsDisableByUserID
}

func AddInnerAgencyStaffIsDisableByUserIDHandler() {
	AddHandler(
		cidl.META_INNER_AGENCY_STAFF_IS_DISABLE_BY_USER_ID,
		func() http.ApiHandler {
			return &InnerAgencyStaffIsDisableByUserIDImpl{
				ApiInnerAgencyStaffIsDisableByUserID: cidl.MakeApiInnerAgencyStaffIsDisableByUserID(),
			}
		},
	)
}

func (m *InnerAgencyStaffIsDisableByUserIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	userId := m.Params.UserID
	dbAgency := db.NewMallAgency()
	staff, err := dbAgency.GetStaff(userId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get staff failed. %s", err)
		return
	}
	
	m.Ack.IsDisable = staff.IsDisable

	ctx.Json(m.Ack)
}
