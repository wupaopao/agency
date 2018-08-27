package db

import (
	"fmt"
	"os"
	"testing"

	"business/agency/cidl"

	"github.com/mz-eco/mz/settings"
)

func TestMain(m *testing.M) {
	settings.LoadFrom("../../", "")
	os.Exit(m.Run())
}

func TestMallAgency_AddOrganization(t *testing.T) {
	dbAgency := NewMallAgency()
	organization := &cidl.Organization{
		Name:            "味罗天下-2",
		Logo:            "http://appweb.morefans.com.cn/share/common/img/app-logo.png",
		Province:        "广东省",
		City:            "深圳市",
		Address:         "深圳市南山区宝深路科陆大厦",
		BankName:        "中国建设银行",
		BankAccount:     "1234 4567 1234 5678 123",
		BankAccountName: "张三-2银行账号",
		CompanyName:     "深圳魔饭科技有限公司",
		LicenseNumber:   "9483732828JQ",
		LicensePicture:  "https://ss0.bdstatic.com/70cFuHSh_Q1YnxGkpoWK1HF6hhy/it/u=2826754939,1597104091&fm=27&gp=0.jpg",
		ManagerUserId:   "6",
		ManagerName:     "张三-2",
		ManagerMobile:   "1867672662",
		PerfectionState: cidl.OrganizationPerfectionStateComplete,
	}
	_, err := dbAgency.AddOrganization(organization)
	if err != nil {
		t.Error(err)
		return
	}

}

func TestMallAgency_GetOrganization(t *testing.T) {
	organization, err := NewMallAgency().GetOrganization(1)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(organization)
}

func TestMallAgency_OrganizationCount(t *testing.T) {
	count, err := NewMallAgency().OrganizationCount()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(count)
}

func TestMallAgency_OrganizationList(t *testing.T) {
	organizations, err := NewMallAgency().OrganizationList(1, 2, false)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(organizations)
}

func TestMallAgency_OrganizationSearchCount(t *testing.T) {
	count, err := NewMallAgency().OrganizationSearchCount("张三-1")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(count)
}

func TestMallAgency_OrganizationSearchList(t *testing.T) {
	organizations, err := NewMallAgency().OrganizationSearchList(1, 2, "张三-1", false)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(organizations)
}

func TestMallAgency_AddStaffRole(t *testing.T) {
	staffRole := &cidl.StaffRole{
		OrganizationId: 1,
		RoleName:       "超级管理员",
		RoleAuthorization: &cidl.RoleAuthorizationMap{
			cidl.SuperAdministratorAuthorizationId: cidl.SuperAdministratorAuthorization,
		},
		Version: cidl.StaffRoleRecordVersion,
		Type:    cidl.StaffRoleTypeSuperAdministrator,
	}
	dbAgency := NewMallAgency()
	_, err := dbAgency.AddStaffRole(staffRole)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestMallAgency_GetStaffRoleByOrgIdRoleId(t *testing.T) {
	staffRole, err := NewMallAgency().GetStaffRoleByOrgIdRoleId(1, 1)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(staffRole)
}

func TestMallAgency_AddStaff(t *testing.T) {
	staff := &cidl.Staff{
		UserId:           "308",
		OrganizationId:   1,
		OrganizationName: "味罗天下",
		Name:             "王五-3",
		Mobile:           "1667672663",
		RoleId:           3,
		RoleName:         "编辑-3",
	}
	_, err := NewMallAgency().AddStaff(staff)
	if err != nil {
		t.Error(err)
		return
	}
}
