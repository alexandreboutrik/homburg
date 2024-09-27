package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	DSN := "host=localhost user=hat_user password=" + HAT_DB_PWD +
		" dbname=hat_db port=9900 sslmode=disable TimeZone=Europe/London"

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  DSN,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
