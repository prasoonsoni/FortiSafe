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

type AssignRoleBody struct {
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
}

type AddAssociatedRolesBody struct {
	ResourceID string `json:"resource_id"`
	Roles      []string `json:"roles"`
}
