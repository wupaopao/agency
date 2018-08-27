package cidl

import "fmt"

type autoErrorserrorsTcidl int

const (
	ErrAddSuperAdministratorIsForbidden autoErrorserrorsTcidl = 4000 //不允许添加超级管理员权限角色
)

func (m autoErrorserrorsTcidl) Number() int { return int(m) }
func (m autoErrorserrorsTcidl) Message() string {
	switch m {

	case ErrAddSuperAdministratorIsForbidden:
		return "不允许添加超级管理员权限角色"
	default:
		return "UNKNOWN_MESSAGE_autoErrorserrorsTcidl"
	}
}
func (m autoErrorserrorsTcidl) Name() string {
	switch m {

	case ErrAddSuperAdministratorIsForbidden:
		return "ErrAddSuperAdministratorIsForbidden"
	default:
		return "UNKNOWN_Name_autoErrorserrorsTcidl"
	}
}
func (m autoErrorserrorsTcidl) String() string {
	return fmt.Sprintf("[%d:%s]%s", m, m.Name(), m.Message())

}
