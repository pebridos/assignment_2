package database

import (
	"assignment_2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "P@ssw0rd"
	dbname   = "order_by"
	db       *gorm.DB
	err      error
)

func StartDB() {
	psqInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(psqInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting to database because: ", err)
	}
	db.Debug().AutoMigrate(models.Orders{}, models.Items{})
}

func GetDB() *gorm.DB {
	return db
}

// func CreateOrder(costumer_name string, order_date time.Time) {
// 	db := GetDB()

// 	newOrder := models.Orders{
// 		CostumerName: costumer_name,
// 		OrderedAt:    order_date,
// 		Items: models.Items{
// 			ItemCode: ,

// 		},
// 	}

// 	err := db.Create(&newOrder).Error

// 	if err != nil {
// 		fmt.Println("Error while insert data because: ", err)
// 		return
// 	}
// 	fmt.Println("New Order with data: ", newOrder)
// }
