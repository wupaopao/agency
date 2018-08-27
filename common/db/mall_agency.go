package db

import (
	"business/agency/cidl"
	"database/sql"
	"errors"
	"fmt"

	"github.com/mz-eco/mz/conn"
	"github.com/mz-eco/mz/log"
)

type MallAgency struct {
	DB *conn.DB
}

func NewMallAgency() *MallAgency {
	return &MallAgency{
		DB: conn.NewDB("mal_agency"),
	}
}

func (m *MallAgency) GetOrganization(organizationID uint32) (organization *cidl.Organization, err error) {
	organization = &cidl.Organization{}
	strSql := `
		SELECT
			org_id,
			name,
			logo,
			province,
			city,
			address,
			post_code,
			bank_name,
			bank_account,
			bank_account_name,
			company_name,
			license_number,
			license_picture,
			manager_uid,
			manager_name,
			manager_mobile,
			group_buying_mode,
			perfection_state,
			is_disable,
			create_time
		FROM
			agc_organization
		WHERE org_id=?
	`
	err = m.DB.Get(organization, strSql, organizationID)
	if err != nil {
		organization = nil
		if err != conn.ErrNoRows {
			log.Warnf("get organization from db failed. %s", err)
			return
		}
		return
	}
	return
}

func (m *MallAgency) GetOrganizationByUserId(userId string) (organization *cidl.Organization, err error) {
	organization = &cidl.Organization{}
	strSql := `
		SELECT
			org_id,
			name,
			logo,
			province,
			city,
			address,
			post_code,
			bank_name,
			bank_account,
			bank_account_name,
			company_name,
			license_number,
			license_picture,
			manager_uid,
			manager_name,
			manager_mobile,
			group_buying_mode,
			perfection_state,
			is_disable,
			create_time
		FROM
			agc_organization
		WHERE manager_uid=?
	`
	err = m.DB.Get(organization, strSql, userId)
	if err != nil {
		organization = nil
		if err != conn.ErrNoRows {
			log.Warnf("get organization from db failed. %s", err)
			return
		}
		return
	}
	return
}

func (m *MallAgency) GetOrganizationId(userID string) (organizationId uint32, err error) {
	strSql := `SELECT org_id FROM agc_organization WHERE manager_uid=?`
	err = m.DB.Get(&organizationId, strSql, userID)
	return
}

func (m *MallAgency) GetEnableOrganization(organizationID uint32) (organization *cidl.Organization, err error) {
	organization = &cidl.Organization{}
	strSql := `
		SELECT
			org_id,
			name,
			logo,
			province,
			city,
			address,
			post_code,
			bank_name,
			bank_account,
			bank_account_name,
			company_name,
			license_number,
			license_picture,
			manager_uid,
			manager_name,
			manager_mobile,
			group_buying_mode,
			perfection_state,
			is_disable,
			create_time
		FROM
			agc_organization
		WHERE org_id=? AND is_disable=0
	`
	err = m.DB.Get(organization, strSql, organizationID)
	if err != nil {
		organization = nil
		if err != conn.ErrNoRows {
			log.Warnf("get organization from db failed. %s", err)
			return
		}
		return
	}
	return
}

func (m *MallAgency) AddOrganization(organization *cidl.Organization) (result sql.Result, err error) {
	strSql := `
		INSERT INTO
			agc_organization
			(
				name,
				logo,
				province,
				city,
				address,
				post_code,
				bank_name,
				bank_account,
				bank_account_name,
				company_name,
				license_number,
				license_picture,
				manager_uid,
				manager_name,
				manager_mobile,
				group_buying_mode,
				perfection_state,
				is_disable
			)
		VALUES
			(
				:name,
				:logo,
				:province,
				:city,
				:address,
				:post_code,
				:bank_name,
				:bank_account,
				:bank_account_name,
				:company_name,
				:license_number,
				:license_picture,
				:manager_uid,
				:manager_name,
				:manager_mobile,
				:group_buying_mode,
				:perfection_state,
				:is_disable
			)
	`
	result, err = m.DB.NamedExec(strSql, organization)
	if err != nil {
		log.Warnf("insert organization failed. %s", err)
		return
	}

	return
}

func (m *MallAgency) UpdateOrganization(organization *cidl.Organization) (result sql.Result, err error) {
	strSql := `
		UPDATE
			agc_organization
		SET
			name=:name,
			logo=:logo,
			province=:province,
			city=:city,
			address=:address,
			post_code=:post_code,
			bank_name=:bank_name,
			bank_account=:bank_account,
			bank_account_name=:bank_account_name,
			company_name=:company_name,
			license_number=:license_number,
			license_picture=:license_picture,
			manager_uid=:manager_uid,
			manager_name=:manager_name,
			manager_mobile=:manager_mobile,
			group_buying_mode=:group_buying_mode,
			perfection_state=:perfection_state
		WHERE
			org_id=:org_id
	`
	result, err = m.DB.NamedExec(strSql, organization)
	if err != nil {
		log.Warnf("update organization failed. %s", err)
		return
	}

	return
}

func (m *MallAgency) OrganizationCount() (count uint32, err error) {
	countSql := `SELECT COUNT(*) FROM agc_organization`
	err = m.DB.Get(&count, countSql)
	return
}

func (m *MallAgency) OrganizationSearchCount(search string) (count uint32, err error) {
	countSql := `
		SELECT
			COUNT(*)
		FROM
			agc_organization
		WHERE
			name LIKE ? OR manager_mobile LIKE ? OR manager_name LIKE ?`
	search = "%" + search + "%"
	err = m.DB.Get(&count, countSql, search, search, search)
	return
}

func (m *MallAgency) OrganizationList(page uint32, pageSize uint32, idAsc bool) (organizations []*cidl.Organization, err error) {
	if page <= 0 || pageSize <= 0 {
		err = errors.New("page or pageSize should be greater than 0")
		return
	}

	offset := (page - 1) * pageSize
	strOrderBy := "ASC"
	if false == idAsc {
		strOrderBy = "DESC"
	}
	listSql := `
		SELECT
			org_id,
			name,
			logo,
			province,
			city,
			address,
			post_code,
			bank_name,
			bank_account,
			bank_account_name,
			company_name,
			license_number,
			license_picture,
			manager_uid,
			manager_name,
			manager_mobile,
			group_buying_mode,
			perfection_state,
			is_disable,
			create_time
		FROM
			agc_organization
		ORDER BY org_id %s
		LIMIT ? OFFSET ?
	`
	listSql = fmt.Sprintf(listSql, strOrderBy)
	rows, err := m.DB.Query(listSql, pageSize, offset)
	if err != nil {
		log.Warnf("query organization list failed. %s", err)
		return
	}

	for rows.Next() {
		var organization cidl.Organization
		err = rows.StructScan(&organization)
		if err != nil {
			log.Warnf("scan organization failed. %s", err)
			return
		}

		organizations = append(organizations, &organization)
	}

	return
}

func (m *MallAgency) OrganizationSearchList(page uint32, pageSize uint32, search string, idAsc bool) (organizations []*cidl.Organization, err error) {
	if page <= 0 || pageSize <= 0 {
		err = errors.New("page or pageSize should be greater than 0")
		return
	}

	offset := (page - 1) * pageSize
	strOrderBy := "ASC"
	if false == idAsc {
		strOrderBy = "DESC"
	}
	listSql := `
		SELECT
			org_id,
			name,
			logo,
			province,
			city,
			address,
			post_code,
			bank_name,
			bank_account,
			bank_account_name,
			company_name,
			license_number,
			license_picture,
			manager_uid,
			manager_name,
			manager_mobile,
			group_buying_mode,
			perfection_state,
			is_disable
		FROM
			agc_organization
		WHERE
			name LIKE ? OR manager_name LIKE ? OR manager_mobile LIKE ?
		ORDER BY org_id %s
		LIMIT ? OFFSET ?
	`
	listSql = fmt.Sprintf(listSql, strOrderBy)
	search = "%" + search + "%"
	rows, err := m.DB.Query(listSql, search, search, search, pageSize, offset)
	if err != nil {
		log.Warnf("query organization list failed. %s", err)
		return
	}

	for rows.Next() {
		var organization cidl.Organization
		err = rows.StructScan(&organization)
		if err != nil {
			log.Warnf("scan organization failed. %s", err)
			return
		}

		organizations = append(organizations, &organization)
	}

	return
}

func (m *MallAgency) GetStaffRoleByOrgIdRoleId(organizationId uint32, roleId uint32) (staffRole *cidl.StaffRole, err error) {
	staffRole = &cidl.StaffRole{}
	strSql := `
		SELECT
			rol_id,
			org_id,
			name,
			authorization,
			is_disable,
			type,
			version
		FROM
			agc_staff_role
		WHERE org_id=? AND rol_id=?
	`
	var strAuthorization string
	queryRow, err := m.DB.QueryRow(strSql,
		organizationId,
		roleId,
	)
	if err != nil {
		log.Warnf("get query row failed. %s", err)
		return
	}

	err = queryRow.Scan(
		&staffRole.RoleId,
		&staffRole.OrganizationId,
		&staffRole.RoleName,
		&strAuthorization,
		&staffRole.IsDisable,
		&staffRole.Type,
		&staffRole.Version,
	)
	if err != nil {
		staffRole = nil
		return
	}

	staffRole.RoleAuthorization = cidl.NewRoleAuthorizationMap()
	err = staffRole.RoleAuthorization.FromString(strAuthorization)
	if err != nil {
		log.Warnf("init role authorization from string failed. %s", err)
		return
	}

	return
}

func (m *MallAgency) AddStaffRole(staffRole *cidl.StaffRole) (result sql.Result, err error) {
	strAuthorization, err := staffRole.RoleAuthorization.ToString()
	if err != nil {
		log.Warnf("marshal role authorization to string failed. %s", err)
		return
	}

	strSql := `
		INSERT INTO agc_staff_role
			(
				org_id,
				name,
				authorization,
				is_disable,
				type,
				version
			)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err = m.DB.Exec(strSql,
		staffRole.OrganizationId,
		staffRole.RoleName,
		strAuthorization,
		staffRole.IsDisable,
		staffRole.Type,
		staffRole.Version,
	)

	return
}

func (m *MallAgency) UpdateStaffRole(organizationId uint32, roleId uint32, roleName string, roleAuthorizations *cidl.RoleAuthorizationMap) (result sql.Result, err error) {
	strAuthorization, err := roleAuthorizations.ToString()
	if err != nil {
		log.Warnf("marshal role authorizations to string failed. %s", err)
		return
	}

	strSql := `UPDATE agc_staff_role SET name=?, authorization=? WHERE org_id=? AND rol_id=?`
	result, err = m.DB.Exec(strSql, roleName, strAuthorization, organizationId, roleId)
	return
}

func (m *MallAgency) SetStaffRoleDisable(organizationId uint32, roleId uint32, isDisable bool) (result sql.Result, err error) {
	strSql := `UPDATE agc_staff_role SET is_disable=? WHERE org_id=? AND rol_id=?`
	result, err = m.DB.Exec(strSql, isDisable, organizationId, roleId)
	return
}

func (m *MallAgency) StaffRoleCount(organizationId uint32, staffRoleType cidl.StaffRoleType) (count uint32, err error) {
	strSql := `SELECT COUNT(*) FROM agc_staff_role WHERE org_id=? AND type=?`
	err = m.DB.Get(&count, strSql, organizationId, staffRoleType)
	return
}

func (m *MallAgency) StaffRoleList(organizationId uint32, staffRoleType cidl.StaffRoleType, page uint32, pageSize uint32, idAsc bool) (staffRoles []*cidl.StaffRole, err error) {
	if page <= 0 || pageSize <= 0 {
		err = errors.New("page or pageSize should be greater than 0")
		return
	}

	offset := (page - 1) * pageSize
	strOrderBy := "ASC"
	if false == idAsc {
		strOrderBy = "DESC"
	}

	strSql := `
		SELECT
			rol_id,
			org_id,
			name,
			authorization,
			is_disable,
			type,
			version
		FROM agc_staff_role
		WHERE org_id=? AND type=?
		ORDER BY rol_id %s
		LIMIT ? OFFSET ?
	`
	strSql = fmt.Sprintf(strSql, strOrderBy)
	rows, err := m.DB.Query(strSql, organizationId, staffRoleType, pageSize, offset)
	if err != nil {
		log.Warnf("query staff role list failed. %s", err)
		return
	}

	for rows.Next() {
		var staffRole cidl.StaffRole
		var strAuthorization string

		err = rows.Scan(
			&staffRole.RoleId,
			&staffRole.OrganizationId,
			&staffRole.RoleName,
			&strAuthorization,
			&staffRole.IsDisable,
			&staffRole.Type,
			&staffRole.Version,
		)
		if err != nil {
			log.Warnf("scan staff role failed. %s", err)
			return
		}

		staffRole.RoleAuthorization = cidl.NewRoleAuthorizationMap()
		err = staffRole.RoleAuthorization.FromString(strAuthorization)
		if err != nil {
			log.Warnf("init role authorization from string failed. %s", err)
			return
		}

		staffRoles = append(staffRoles, &staffRole)
	}

	return
}

func (m *MallAgency) StaffCount(organizationId uint32) (count uint32, err error) {
	strSql := `SELECT COUNT(*) FROM agc_staff WHERE org_id=?`
	err = m.DB.Get(&count, strSql, organizationId)
	return
}

func (m *MallAgency) StaffList(organizationId uint32, page uint32, pageSize uint32, idAsc bool) (staffs []*cidl.Staff, err error) {
	if page <= 0 || pageSize <= 0 {
		err = errors.New("page or pageSize should be greater than 0")
		return
	}

	offset := (page - 1) * pageSize
	strOrderBy := "ASC"
	if false == idAsc {
		strOrderBy = "DESC"
	}
	strSql := `
		SELECT
			org_id,
			uid,
			org_name,
			name,
			mobile,
			rol_id,
			rol_name,
			is_disable,
			create_time
		FROM
			agc_staff
		WHERE org_id=?
		ORDER BY uid %s
		LIMIT ? OFFSET ?
	`
	strSql = fmt.Sprintf(strSql, strOrderBy)
	rows, err := m.DB.Query(strSql, organizationId, pageSize, offset)
	if err != nil {
		log.Warnf("query staff list failed. %s", err)
		return
	}

	for rows.Next() {
		var staff cidl.Staff
		err = rows.StructScan(&staff)
		if err != nil {
			log.Warnf("scan staff failed. %s", err)
			return
		}

		staffs = append(staffs, &staff)
	}

	return
}

func (m *MallAgency) AddStaff(staff *cidl.Staff) (result sql.Result, err error) {
	strSql := `
		INSERT INTO agc_staff
			(
				uid,
				org_id,
				org_name,
				name,
				mobile,
				rol_id,
				rol_name,
				is_disable
			)
		VALUES
			(
				:uid,
				:org_id,
				:org_name,
				:name,
				:mobile,
				:rol_id,
				:rol_name,
				:is_disable
			)
	`
	result, err = m.DB.NamedExec(strSql, staff)
	return
}

func (m *MallAgency) GetStaff(userId string) (staff *cidl.Staff, err error) {
	staff = &cidl.Staff{}
	strSql := `
		SELECT
			uid,
			org_id,
			org_name,
			name,
			mobile,
			rol_id,
			rol_name,
			is_disable,
			create_time
		FROM
			agc_staff
		WHERE uid=?
	`
	err = m.DB.Get(staff, strSql, userId)
	if err != nil {
		staff = nil
		log.Warnf("query staff failed. %s", err)
		return
	}

	return
}

// 添加超级管理员
func (m *MallAgency) AddSuperAdministrator(organizationId uint32, staff *cidl.Staff) (success bool, err error) {

	dbAgency := NewMallAgency()

	// 添加角色
	roleName := "超级管理员"
	staffRole := &cidl.StaffRole{
		OrganizationId: organizationId,
		RoleName:       roleName,
		RoleAuthorization: &cidl.RoleAuthorizationMap{
			cidl.SuperAdministratorAuthorizationId: cidl.SuperAdministratorAuthorization,
		},
		Type:    cidl.StaffRoleTypeSuperAdministrator,
		Version: cidl.StaffRoleRecordVersion,
	}
	result, err := dbAgency.AddStaffRole(staffRole)
	if err != nil {
		log.Warnf("add staff role failed. %s", err)
		return
	}

	roleId, err := result.LastInsertId()
	if err != nil {
		log.Warnf("get new role id failed. %s", err)
		return
	}

	// 添加成员
	staff.RoleId = uint32(roleId)
	staff.RoleName = roleName
	_, err = dbAgency.AddStaff(staff)
	if err != nil {
		log.Warnf("add staff failed.")
		return
	}

	success = true

	return
}
