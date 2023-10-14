package configs

import "github.com/celpung/clean-gin-architecture/internal/entity"

func AutoMigrage() {
	ConnectDatabase()

	if migrateErr := DB.AutoMigrate(
		&entity.User{},
	); migrateErr != nil {
		panic(migrateErr)
	}
}
