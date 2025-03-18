package databases

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		log.Fatal("env is undifined")
	}
	hostDev := os.Getenv("DB_DEV_HOST")
	usernameDev := os.Getenv("DB_DEV_USERNAME")
	passDev := os.Getenv("DB_DEV_PASS")
	nameDev := os.Getenv("DB_DEV_NAME")
	portDev := os.Getenv("DB_DEV_PORT")
	hostTest := os.Getenv("DB_TEST_HOST")
	usernameTest := os.Getenv("DB_TEST_USERNAME")
	passTest := os.Getenv("DB_TEST_PASS")
	nameTest := os.Getenv("DB_TEST_NAME")
	portTest := os.Getenv("DB_TEST_PORT")
	hostProd := os.Getenv("DB_PROD_HOST")
	usernameProd := os.Getenv("DB_PROD_USERNAME")
	passProd := os.Getenv("DB_PROD_PASS")
	nameProd := os.Getenv("DB_PROD_NAME")
	portProd := os.Getenv("DB_PROD_PORT")

	var url string
	switch {
	case appEnv == "development":
		url = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", usernameDev, passDev, hostDev, portDev, nameDev)
	case appEnv == "test":
		url = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", usernameTest, passTest, hostTest, portTest, nameTest)
	case appEnv == "production":
		url = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", usernameProd, passProd, hostProd, portProd, nameProd)
	default:
		url = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", usernameDev, passDev, hostDev, portDev, nameDev)
	}

	if url == "" {
		log.Fatal("Error: url is undifined")
	}
	conn, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		panic(err)
	}

	if conn != nil {
		log.Printf("Berhasil terkoneksi ke database %s.......", appEnv)
	}

	DB = conn
}
