package models

type CreateRoleBody struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

type AddPermissionBody struct {
	RoleID      string   `json:"role_id"`
	Permissions []string `json:"permissions"`
}

type DeletePermissionBody struct {
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
}
