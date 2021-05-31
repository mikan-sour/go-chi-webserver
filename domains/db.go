package domains

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Database *gorm.DB
}

func (d *DB) PostgresConnect() {

	dsn := fmt.Sprintf("host=localhost user=postgres password=%s dbname=postgres port=5433 sslmode=disable TimeZone=Asia/Tokyo", "password")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sql, _ := db.DB()
	sql.Ping()

	d.Database = db

}
