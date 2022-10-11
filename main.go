package main

import (
	"assignment_2/database"
	"assignment_2/routers"
)

func main() {
	//Database
	database.StartDB()

	// format := "2006-02-01 15:04:05"
	// date, _ := time.Parse(format, "2019-11-0921:21:46+00:00")
	// fmt.Println(date)
	// database.CreateOrder("Pebrido", date)

	routers.Start().Run("localhost:8001")

}
