package data

import(
	"log"

	"gorm.io/driver/sqlite"
    "gorm.io/gorm"

    "github.com/jasimpn/trace/models"
)
var DB *gorm.DB

func connectToSQLite() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}

func CreateSQLite(){
	var err error
	DB, err = connectToSQLite()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the User model
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate user database:", err)
	}

	// Auto-migrate the Repo model
	if err := DB.AutoMigrate(&models.Repo{}); err != nil {
		log.Fatal("Failed to migrate repo database:", err)
	}

	// Auto-migrate the Repo model
	if err := DB.AutoMigrate(&models.Storage{}); err != nil {
		log.Fatal("Failed to migrate storage database:", err)
	}

}