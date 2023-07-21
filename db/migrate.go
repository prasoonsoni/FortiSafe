package db

import (
	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/models"
)

func Migrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.AccountStatusLogs{})
	DB.AutoMigrate(&models.Permission{})
	DB.AutoMigrate(&models.Role{})
	DB.AutoMigrate(&models.RolePermission{})
	DB.AutoMigrate(&models.Resource{})
	DB.AutoMigrate(&models.Group{})
}
