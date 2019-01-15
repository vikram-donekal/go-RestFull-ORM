package Helper

import (
	"../../../bin/gorm"
	_ "../../../bin/pq"
	"../DTO"
	"log"
	"errors"
)

var DatabaseConnection *gorm.DB //database



/*const (
	DB_USER     = "postgresadmin"
	DB_PASSWORD = "admin123"
	DB_NAME     = "postgresdb"
)
*/


func  init ()   {
	log.Print("connecting to Data Base Postgresql")

	/*DatabaseInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)*/

	err := errors.New("")
	DatabaseConnection,err  = gorm.Open("postgres", "host=10.177.200.24 port=30005 user=postgresadmin dbname=postgresdb password=admin123 sslmode=disable")
	if err!= nil {
		panic(err.Error())
	}
	//defer  	DatabaseConnection.Close()
	DataBase:=DatabaseConnection.DB()
	//defer  DataBase.Close()
	err =DataBase.Ping()

	if err!= nil {
		panic(err.Error())
	}
	log.Println("creating tables using GORM ")
	DatabaseConnection.DropTableIfExists(&DTO.Customer{})
	DatabaseConnection.CreateTable(&DTO.Customer{})
}


func GetDB() *gorm.DB {
	return DatabaseConnection
}
