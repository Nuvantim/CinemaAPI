package request

import repo "api/internal/app/repository"

type Permission struct {
	Name string `validate:"required" json:"name"`
}

type Role struct {
	Name         string  `validate:"required" json:"name"`
	PermissionID []int64 `validate:"omitempty,dive,gt=0" json:"permission_id"`
}

type GetRole struct {
	ID         int64                       `json:"id"`
	Name       string                      `json:"name"`
	Permission []repo.GetPermissionRoleRow `json:"permission"`
}

type GetClient struct {
	ID    int64                   `json:"id"`
	Name  string                  `json:"name"`
	Email string                  `json:"email"`
	Role  []repo.GetRoleClientRow `json:"permission"`
}
