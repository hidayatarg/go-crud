package initalizers

import "github.com/hidayatarg/go-crud/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.User{})

	// add an initial user
	AddInitialUser()
}

func AddInitialUser() {
	DB.Create(&models.User{Email: "admin@admin.com", Password: "1234"})
}
