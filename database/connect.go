package database

import (
	"fmt"
	"strconv"

	"github.com/horlabyc/go-todo/config"
	"github.com/jinzhu/gorm"
)

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	configData := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"),
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)
	DB, err := gorm.Open("postgres", configData)
	if err != nil {
		fmt.Println(
			err.Error(),
		)
		panic("connection to database failed")
	}
	fmt.Println("Connection Opened to Database")
}