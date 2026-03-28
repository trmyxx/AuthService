package initializers

import "github.com/trmyxx/AuthService/internal/model"

func SyncDataBase() {
	DB.AutoMigrate(&model.User{})
}
