package main

import (
	"github.com/hidayatarg/go-crud/initalizers"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDatabase()
}

func main() {
	initalizers.SyncDatabase()
	// moved the db initalization to sync database
}
