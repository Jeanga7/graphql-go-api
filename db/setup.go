package db

import (
	"log"

	"github.com/Jeanga7/graphql-go-api.git/graph/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	// data source name (liens de connection a ma DB postgres)
	dsn := "host=localhost user=user password=password dbname=graphql_go_api port=5432 sslmode=disable TimeZone=UTC"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Impossible de se connecter à la base de données :", err)
	}

	DB = database
	log.Println("Connexion à la base de données réussie")
	return DB
}

func MigrateDatabase() {
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Échec de la migration :", err)
	}

	log.Println("Migration réussie")
}
