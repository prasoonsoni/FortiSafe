package db

import (
	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/models"
)

func Migrate() {
	_ = DB.AutoMigrate(&models.User{})
	_ = DB.AutoMigrate(&models.AccountStatusLogs{})
	_ = DB.AutoMigrate(&models.Permission{})
	_ = DB.AutoMigrate(&models.Role{})
}
