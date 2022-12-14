package db

import (
	"comedians/src/cfg"
	commentsModel "comedians/src/core/concerts/comments/model"
	usersConcertsModel "comedians/src/core/usersConcerts/model"	
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() error {
	c := cfg.Config()
	dsn := fmt.Sprintf("host=%s port=%s user=%s sslmode=disable password=%s dbname=%s", c.PgHost, c.PgPort, c.PgUser, c.PgPassword, c.PgDbName)
	var err error

	DBS, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = DBS.AutoMigrate(&usersConcertsModel.Permission{}, &usersConcertsModel.Role{}, &usersConcertsModel.User{}, &usersConcertsModel.Concert{}, &commentsModel.Comment{})

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

var DBS *gorm.DB
