package migration

import (
	"fmt"
	"go-smart-dormitory/database"
	"go-smart-dormitory/model/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.Penghuni{})

	if err != nil {
		log.Println("Error running migration")
	}

	fmt.Println("Migration run successfully")
}
