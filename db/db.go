package db

import (
	"fmt"
	"go-embedded-svelte/config"

	"github.com/rs/zerolog/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	// connect to the db
	if db != nil {
		return db
	}

	config := config.GetConfig()
	path := fmt.Sprintf("%s/data.db", config.DataDir)

	log.Debug().Msg("Loading database from " + path)

	localdb, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = localdb
	db.AutoMigrate(&User{})

	return db
}
