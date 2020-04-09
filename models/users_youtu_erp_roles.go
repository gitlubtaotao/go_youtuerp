package models

type UsersYoutuErpRoles struct {
	UserId int   `xorm:"index(index_users_youtu_erp_roles_on_user_id_and_role_id) INT(11)"`
	RoleId int64 `xorm:"index index(index_users_youtu_erp_roles_on_user_id_and_role_id) BIGINT(20)"`
}
