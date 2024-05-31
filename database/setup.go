package database

import (
	"github.com/amos-pbi/final-task-pbi-rakamin-fullstack-amosaugusto/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := "root:@tcp(localhost:3306)/final_task_pbi_rakamin_fullstack"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = db
	db.AutoMigrate(&models.User{}, &models.Photo{})
}
