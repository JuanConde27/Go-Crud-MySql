package commons

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"modcrudmysql.com/src/models"
)

func GetConnection() *gorm.DB {
    dsn := "juan:123456@tcp(localhost:3306)/golang_db?charset=utf8&parseTime=True&loc=Local"
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }
    return database
}

func Migrate() {
	db := GetConnection()
	// Cerrar la conexión al final
	sqlDB, err := db.DB()
	if err != nil {
		// Manejar el error de la base de datos
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// Migrar el modelo de persona
	db.AutoMigrate(&models.Persona{})

	log.Println("Migración completada.")
}

