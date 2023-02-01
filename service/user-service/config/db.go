package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConnection() *gorm.DB{
	var dbHosst = "localhost"
	var dbUser = "developer"
	var dbPassword = "devspassword"
	var dbName = "user"
	var dbPort = "5432" 
	
	dbString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHosst, dbUser, dbPassword, dbName, dbPort)
	dbCon, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return dbCon
}