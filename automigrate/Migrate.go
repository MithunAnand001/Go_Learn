package automigrate

import (
	database "learn_service/Database"
	"learn_service/entity"
	"log"
	"log/slog"
)

func MigrateTables() {
	Db, err := database.ConnectDB()
	if err != nil{
		log.Println("Error occured while connecting with database",err.Error())
	}
	defer database.CloseDB(Db)

	err = Db.Debug().AutoMigrate(&entity.User{})
	if err != nil{
		slog.Error("Error occured while migrating Users table")
	}
}