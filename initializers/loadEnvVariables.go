package initializers

import "Auth/model"

func SyncDataBase() {
	DB.AutoMigrate(&model.User{})
}
