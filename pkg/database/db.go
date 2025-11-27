package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	Offset(offset int) *gorm.DB
	Limit(limit int) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) Database
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Model(model interface{}) *gorm.db
	First(dest interface{}, conds ...interface{}) Database
	Save(value interface{}) *gorm.DB
	Updates(interface{}) *gorm.DB
	Order(value interface{}) *gorm.DB
	Error() error
}

type GormDatabase struct {
	*gorm.DB
}

func (db *GormDatabase) Where(query interface{}, args ...interface{}) Database {
	return &GormDatabase{db.DB.Where(query, args...)}
}
func (db *GormDatabase) First(dest interface{}, conds ...interface{}) Database {
	return &GormDatabase{db.DB.First(dest, conds...)}
}
func (db *GormDatabase) Error() error {
	return db.DB.Error
}

func NewDatabase() *gorm.DB {
	var database *gorm.DB
	var err error

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, pass, name, port)

	for i := 1; i <= 3; i++ {
		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		} else {
			log.Printf("Attempt %d: Failed to initialize database. Retrying...", i)
			time.Sleep(3 * time.Second)
		}
	}
	
	database.AutoMigrate(&models.Todo{})
	return database
}
