package database

import (
	"backend-sia/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	env := config.Renderenv("ENV")
	if env == "" {
		panic("env tidak terdeteksi..............")
	}

	host_dev := config.Renderenv("DB_DEV_HOST")
	name_dev := config.Renderenv("DB_DEV_NAME")
	username_dev := config.Renderenv("DB_DEV_USERNAME")
	pass_dev := config.Renderenv("DB_DEV_PASS")
	port_dev := config.Renderenv("DB_DEV_PORT")

	host_test := config.Renderenv("DB_TEST_HOST")
	name_test := config.Renderenv("DB_TEST_NAME")
	username_test := config.Renderenv("DB_TEST_USERNAME")
	pass_test := config.Renderenv("DB_TEST_PASS")
	port_test := config.Renderenv("DB_TEST_PORT")

	host_prod := config.Renderenv("DB_PROD_HOST")
	name_prod := config.Renderenv("DB_PROD_NAME")
	username_prod := config.Renderenv("DB_PROD_USERNAME")
	pass_prod := config.Renderenv("DB_PROD_PASS")
	port_prod := config.Renderenv("DB_PROD_PORT")

	var url string
	switch {
	case env == "development":
		url = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", username_dev, pass_dev, host_dev, port_dev, name_dev)
	case env == "test":
		url = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", username_test, pass_test, host_test, port_test, name_test)
	case env == "production":
		url = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", username_prod, pass_prod, host_prod, port_prod, name_prod)
	default:
		url = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", username_dev, pass_dev, host_dev, port_dev, name_dev)
	}

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if db != nil {
		log.Printf("Koneksi %s berhasil.......", env)
	}
	DB = db
}
